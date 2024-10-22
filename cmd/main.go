package main

import (
	"lido-events/internal/adapters/ethereum"
	"lido-events/internal/infrastructure/cache"
	"lido-events/internal/infrastructure/config"
	"log"
	"net/http"

	"lido-events/internal/adapters/api"
	telegram "lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Load the partial config (operatorId, telegram) from the JSON file
	appConfig, err := config.LoadAppConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load partial config: %v", err)
	}

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

	// Initialize the file storage adapter
	fileStorage := &storage.FileStorage{
		PerformanceFile: "validators-performance.json",
		ExitRequestFile: "exit-requests.json",
		ConfigFile:      "config.json",
	}

	// Initialize the Telegram notifier using the loaded Telegram token
	notifier, err := telegram.NewTelegramNotifier(appConfig.Telegram.Token, appConfig.Telegram.ChatID)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	// Initialize the IndexerService with storage and the Telegram bot as the notifier
	service := &services.Service{
		Storage:  fileStorage,
		Notifier: notifier,
	}

	// Initialize the Ethereum adapter using cached ABIs
	ethereumAdapter, err := ethereum.NewEthereumAdapter(networkConfig.EtherscanURL, networkConfig.ContractABIs, abiCache)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum adapter: %v", err)
	}

	// Start Ethereum event subscription, passing the event handler from the service
	go ethereumAdapter.SubscribeToEvents(service.HandleEthereumEvent)

	handler := &api.APIHandler{
		IndexerService: service,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v0/events_indexer/telegramToken", handler.UpdateTelegramToken).Methods("POST")
	r.HandleFunc("/api/v0/events_indexer/operatorId", handler.UpdateOperatorID).Methods("POST")
	r.HandleFunc("/api/v0/event_indexer/lido_report", handler.GetLidoReport).Methods("GET")
	r.HandleFunc("/api/v0/event_indexer/exit_requests", handler.GetExitRequests).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	log.Println("App running...")
	select {} // Keep the app running
}
