package main

import (
	"context"
	"lido-events/internal/adapters/api"
	"lido-events/internal/adapters/ethereum"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/aplication/services"
	"lido-events/internal/config"
	"log"
	"net/http"
)

// Hexagonal architecrtures principles:
// - domain: cannot import ✅
// - aplication:
//   - ports: can only import domain ✅
//   - services: can only import domain ✅
// - adapters: can import domain and services ✅
// - config: can only import domain ✅
//     - The config layer is responsible for loading certain configurations from the environment and files that adapters might need.

// Key points:
// - All the communications to the services are done through the ports.
// - It is okey that 1 service uses another service in ordert to be initiated.

func main() {
	// Load configurations
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}
	abiCache, err := config.NewABICache(networkConfig.ContractABIs)
	if err != nil {
		log.Fatalf("Failed to load ABIs: %v", err)
	}
	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load aplication configuration: %v", err)
	}

	// Initialize adapters
	storageAdapter := storage.NewStorageAdapter()
	notifierAdapter, err := notifier.NewNotifierAdapter(appConfig.Telegram.Token, appConfig.Telegram.ChatID)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram notifier: %v", err)
	}
	subscriberAdapter, err := ethereum.NewEthereumSubscriber(networkConfig.WsURL, abiCache.GetAllABIs())
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum subscriber: %v", err)
	}

	// Initialize services
	notifierService := services.NewNotifierService(notifierAdapter)

	storageService := services.NewStorageService(storageAdapter)
	eventService := services.NewEventService(storageAdapter, notifierAdapter, subscriberAdapter)

	// Start subscribing to events. Done by eventService.
	if err := eventService.SubscribeToEvents(context.Background()); err != nil {
		log.Fatalf("Failed to subscribe to events: %v", err)
	}

	// Initialize the API handler (internally sets up routes)
	apiAdapter := api.NewAPIAdapter(storageService, notifierService)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiAdapter.Router))
}
