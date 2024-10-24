package main

import (
	"lido-events/internal/adapters/ethereum"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/aplication/services"
	"lido-events/internal/config"
	"log"
	"net/http"

	"lido-events/internal/adapters/api"
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
	// (INFRA) Load the network-specific configuration and contract ABIs
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}

	// (INFRA) Load the cached ABIs from the network configuration
	abiCache, err := config.NewABICache(networkConfig.ContractABIs)
	if err != nil {
		log.Fatalf("Failed to load ABIs: %v", err)
	}

	// (INFRA) Load the application configuration
	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	// (ADAPTER) Initialize the file storage adapter
	storageAdapter := storage.NewStorageAdapter()

	// (ADAPTER) Initialize Notifier Adapter (Telegram)
	notifierAdapter, err := notifier.NewNotifierAdapter(appConfig.Telegram.Token, appConfig.Telegram.ChatID)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram notifier: %v", err)
	}

	// (SERVICE) Initialize Notifier Service
	notifierService := services.NewNotifierService(notifierAdapter)

	// (SERVICE) Initialize Operator Service
	storageService := services.NewStorageService(storageAdapter)

	// (SERVICE) Initialize Event Handler with all adapters it needs
	eventService := services.NewEventService(storageAdapter, notifierAdapter)

	// (ADAPTER) Initialize Ethereum Subscriber and Start Subscribing to Events
	// This is the only adapter that needs a running application service to work, so it is initialized last.
	// It is still an adapter, since it mediates between the Ethereum network and an application service.
	ethSubscriber, err := ethereum.NewEthereumSubscriber(networkConfig.WsURL, abiCache.GetAllABIs(), eventService)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum subscriber: %v", err)
	}
	go ethSubscriber.SubscribeToEvents()

	// (INFRA) Initialize the API handler (internally sets up routes)
	apiAdapter := api.NewAPIAdapter(storageService, notifierService)

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiAdapter.Router))

}
