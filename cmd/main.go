package main

import (
	"lido-events/internal/adapters/ethereum"
	"lido-events/internal/domain"
	"log"
	"net/http"

	"lido-events/internal/adapters/api"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/telegram"
	"lido-events/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the Ethereum adapter with ABIs and contract addresses
	abiPaths := []string{
		"abi/CSAccounting.json",
		"abi/CSFeeDistributor.json",
		"abi/CSFeeOracle.json",
		"abi/CSModule.json",
		"abi/VEBO.json",
	}
	contractAddresses := []string{
		"0xContractAddress1",
		"0xContractAddress2",
		"0xContractAddress3",
		"0xContractAddress4",
		"0xContractAddress5",
	}

	ethereumAdapter, err := ethereum.NewEthereumAdapter("https://mainnet.infura.io/v3/YOUR_PROJECT_ID", abiPaths, contractAddresses)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum adapter: %v", err)
	}

	// Start listening to events
	go ethereumAdapter.SubscribeToEvents()

	// Initialize the file storage adapter
	fileStorage := &storage.FileStorage{
		PerformanceFile: "validators-performance.json",
		ExitRequestFile: "exit-requests.json",
		ConfigFile:      "config.json",
	}

	// Load the configuration from the config.json file
	config, err := fileStorage.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the Telegram bot as the notifier using the loaded Telegram token
	var notifier domain.Notifier
	if config.TelegramToken != "" {
		bot, err := telegram.NewTelegramNotifier(config.TelegramToken, 0) // TODO: Add chat ID
		if err != nil {
			log.Fatalf("Failed to initialize Telegram bot: %v", err)
		}
		notifier = bot
	} else {
		log.Fatal("No Telegram token provided in the config")
	}

	// Initialize the IndexerService with storage and the Telegram bot as the notifier
	indexerService := &services.IndexerService{
		Storage:  fileStorage,
		Notifier: notifier,
		Config:   config,
	}

	handler := &api.APIHandler{
		IndexerService: indexerService,
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
