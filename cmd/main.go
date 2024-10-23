package main

import (
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/eventhandler"
	"lido-events/internal/infrastructure/cache"
	"lido-events/internal/infrastructure/config"
	"lido-events/internal/service"
	"lido-events/internal/subscriber"
	"log"
	"net/http"
	"os"

	"lido-events/internal/adapters/api"

	"github.com/gorilla/mux"
)

func main() {
	// Load the network-specific configuration and contract ABIs
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}

	// Load the cached ABIs from the network configuration
	abiCache, err := cache.NewABICache(networkConfig.ContractABIs)
	if err != nil {
		log.Fatalf("Failed to load ABIs: %v", err)
	}

	appConfig, err := config.LoadAppConfig("config.json")
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
	notifierService := service.NewNotifierService(notifierAdapter)

	// Initialize Operator Service
	operatorService := service.NewOperatorService(storageAdapter)

	// Initialize Event Handler with services
	eventHandler := eventhandler.NewEventHandler(operatorService, notifierService)

	// Initialize Ethereum Subscriber and Start Subscribing to Events
	wsURL := os.Getenv("WS_URL")
	if wsURL == "" {
		wsURL = "wss://mainnet.infura.io/ws/v3/YOUR_PROJECT_ID" // Replace with your default WebSocket URL
	}

	ethSubscriber, err := subscriber.NewEthereumSubscriber(wsURL, abiCache.GetAllABIs(), eventHandler)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum subscriber: %v", err)
	}

	go ethSubscriber.SubscribeToEvents()

	// Initialize API Handler
	apiHandler := &api.APIHandler{
		OperatorService: operatorService,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v0/events_indexer/telegramConfig", apiHandler.UpdateTelegramConfig).Methods("POST")
	r.HandleFunc("/api/v0/events_indexer/operatorId", apiHandler.UpdateOperatorID).Methods("POST")
	r.HandleFunc("/api/v0/event_indexer/lido_report", apiHandler.GetLidoReport).Methods("GET")
	r.HandleFunc("/api/v0/event_indexer/exit_requests", apiHandler.GetExitRequests).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	log.Println("App running...")
}
