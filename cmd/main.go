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
// - aplication:
//   - ports: can only import domain ✅
//   - services: can only import domain ✅
// - adapters: can import domain and services ✅
// - infrastructure: can only import domain and ports ✅
//     - The infrastructure layer is responsible for handling technical details and external dependencies, such as database connections, configurations, and integrations with third-party services, to support the core application's functionality

// Key points:
// - All the communications to the services are done through the ports.
// - It is okey that 1 service uses another service in ordert to be initiated.

func main() {
	// (INFRA) Load the network-specific configuration and contract ABIs
	networkConfig, err := infrastructure.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}

	// (INFRA) Load the cached ABIs from the network configuration
	abiCache, err := infrastructure.NewABICache(networkConfig.ContractABIs)
	if err != nil {
		log.Fatalf("Failed to load ABIs: %v", err)
	}

	// (INFRA) Load the application configuration
	appConfig, err := infrastructure.LoadAppConfig("config.json")
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

	// (INFRA) Initialize Ethereum Subscriber and Start Subscribing to Events
	ethSubscriber, err := infrastructure.NewEthereumSubscriber(networkConfig.WsURL, abiCache.GetAllABIs(), eventService)
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
