package main

import (
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/aplication/services"
	"lido-events/internal/infrastructure"
	"log"
	"net/http"

	"lido-events/internal/adapters/api"
)

// Hexagonal architecrtures principles:
// - domain: cannot import ✅
// - services: can only import domain ✅
// - adapters: can import domain and services ✅
// - infrastructure: can only import domain ✅

// Key points:
// - All the communications to the services are done through the ports.

func main() {
	// Load the network-specific configuration and contract ABIs
	networkConfig, err := infrastructure.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}

	// Load the cached ABIs from the network configuration
	abiCache, err := infrastructure.NewABICache(networkConfig.ContractABIs)
	if err != nil {
		log.Fatalf("Failed to load ABIs: %v", err)
	}

	appConfig, err := infrastructure.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	// Initialize the file storage adapter
	storageAdapter := storage.NewStorageAdapter()

	// Initialize Notifier Adapter (Telegram)
	notifierAdapter, err := notifier.NewNotifierAdapter(appConfig.Telegram.Token, appConfig.Telegram.ChatID)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram notifier: %v", err)
	}

	// Initialize Notifier Service
	notifierService := services.NewNotifierService(notifierAdapter)

	// Initialize Operator Service
	storageService := services.NewStorageService(storageAdapter)

	// Initialize Event Handler with services
	eventService := services.NewEventService(storageService, notifierService)

	// Initialize Ethereum Subscriber and Start Subscribing to Events
	ethSubscriber, err := infrastructure.NewEthereumSubscriber(networkConfig.WsURL, abiCache.GetAllABIs(), eventService)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum subscriber: %v", err)
	}
	go ethSubscriber.SubscribeToEvents()

	// Initialize the API handler (internally sets up routes)
	apiHandler := api.NewAPIHandler(storageService, notifierService)

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiHandler.Router))

}
