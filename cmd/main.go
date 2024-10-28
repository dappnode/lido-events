package main

import (
	"context"
	"lido-events/internal/adapters/api"
	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"
	csmodule "lido-events/internal/adapters/csModule"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/vebo"
	"math/big"

	"lido-events/internal/aplication/services"
	"lido-events/internal/config"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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
	// TODO: what happens when operator id changes
	// TODO: what happens if new validators -> new validatorIndexes. Affects vebo adapter. We should
	// TODO: where to get validatorIndexes and refSlot from
	veboAdapter, err := vebo.NewVeboAdapter(networkConfig.WsURL, networkConfig.VEBOAddress, []*big.Int{networkConfig.CSMStakingModuleID}, []*big.Int{appConfig.OperatorID}, []*big.Int{}, []*big.Int{})
	if err != nil {
		log.Fatalf("Failed to initialize Vebo adapter: %v", err)
	}
	// TODO: where to get oldAddress, newAddress, oldProposedAddress, newProposedAddress from
	csModuleAdapter, err := csmodule.NewCsModuleAdapter(networkConfig.WsURL, networkConfig.CSModuleAddress, []*big.Int{appConfig.OperatorID}, []common.Address{}, []common.Address{}, []common.Address{}, []common.Address{})
	if err != nil {
		log.Fatalf("Failed to initialize CsModule adapter: %v", err)
	}
	csFeeDistributorAdapter, err := csfeedistributor.NewCsFeeDistributorAdapter(networkConfig.WsURL, networkConfig.CSFeeDistributorAddress)
	if err != nil {
		log.Fatalf("Failed to initialize CsFeeDistributor adapter: %v", err)
	}

	// Initialize services
	notifierService := services.NewNotifierService(notifierAdapter)
	storageService := services.NewStorageService(storageAdapter)
	veboService := services.NewVeboService(storageAdapter, notifierAdapter, veboAdapter)
	csModuleService := services.NewCsModuleService(storageAdapter, notifierAdapter, csModuleAdapter)
	csFeeDistributorService := services.NewCsFeeDistributorService(storageAdapter, notifierAdapter, csFeeDistributorAdapter)

	// Start subscribing to each SC events. Done by sc services.
	if err := veboService.WatchVeboEvents(context.Background()); err != nil {
		log.Fatalf("Failed to subscribe to VEBO events: %v", err)
	}
	if err := csModuleService.WatchCsModuleEvents(context.Background()); err != nil {
		log.Fatalf("Failed to subscribe to CSModule events: %v", err)
	}
	if err := csFeeDistributorService.WatchCsFeeDistributorEvents(context.Background()); err != nil {
		log.Fatalf("Failed to subscribe to CSFeeDistributor events: %v", err)
	}

	// Initialize the API handler (internally sets up routes)
	apiAdapter := api.NewAPIAdapter(storageService, notifierService)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiAdapter.Router))
}
