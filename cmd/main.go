package main

import (
	"context"
	"lido-events/internal/adapters/api"
	"lido-events/internal/adapters/beaconchain"
	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"
	csfeedistributorimpl "lido-events/internal/adapters/csFeeDistributorImpl"
	csmodule "lido-events/internal/adapters/csModule"
	exitvalidator "lido-events/internal/adapters/exitValidator"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/vebo"
	"time"

	"lido-events/internal/application/services"
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

// Helper function to check if operator IDs and Telegram config are available
func waitForInitialConfig(storageAdapter *storage.Storage) {
	for {
		// Check for operator IDs
		operatorIds, err := storageAdapter.GetOperatorIds()
		if err != nil || len(operatorIds) == 0 {
			log.Println("Waiting for operator IDs to be set...")
		} else {
			// Check for Telegram config
			telegramConfig, err := storageAdapter.GetTelegramConfig()
			if err != nil || telegramConfig.Token == "" || telegramConfig.UserID == 0 {
				log.Println("Waiting for Telegram config to be set...")
			} else {
				// Both operator IDs and Telegram config are set
				log.Println("Operator IDs and Telegram config are set. Proceeding with initialization.")
				return
			}
		}
		time.Sleep(2 * time.Second) // Poll every 2 seconds
	}
}

func main() {
	// Set up context to control the lifetime of the services
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensures cancel is called when main exits

	// Load configurations
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}
	// print all the networkConfig
	log.Printf("Network config: %+v", networkConfig)

	// Initialize adapters
	storageAdapter := storage.NewStorageAdapter()
	apiAdapter := api.NewAPIAdapter(storageAdapter)

	// Start HTTP server
	go func() {
		log.Println("Server started on :8080")
		log.Fatal(http.ListenAndServe(":8080", apiAdapter.Router))
	}()

	// Wait for initial configuration to be set
	waitForInitialConfig(storageAdapter)

	beaconchainAdapter := beaconchain.NewBeaconchainAdapter(networkConfig.BeaconchainURL)
	exitValidatorAdapter := exitvalidator.NewExitValidatorAdapter(beaconchainAdapter, networkConfig.SignerUrl)
	notifierAdapter, err := notifier.NewNotifierAdapter(ctx, storageAdapter)
	if err != nil {
		log.Fatalf("Failed to initialize Telegram notifier: %v", err)
	}

	csFeeDistributorImplAdapter, err := csfeedistributorimpl.NewCsFeeDistributorImplAdapter(networkConfig.WsURL, networkConfig.CSFeeDistributorAddress)
	if err != nil {
		log.Fatalf("Failed to initialize CsFeeDistributor adapter: %v", err)
	}

	veboAdapter, err := vebo.NewVeboAdapter(networkConfig.WsURL, networkConfig.VEBOAddress, storageAdapter)
	if err != nil {
		log.Fatalf("Failed to initialize Vebo adapter: %v", err)
	}

	csModuleAdapter, err := csmodule.NewCsModuleAdapter(networkConfig.WsURL, networkConfig.CSModuleAddress, storageAdapter)
	if err != nil {
		log.Fatalf("Failed to initialize CsModule adapter: %v", err)
	}

	csFeeDistributorAdapter, err := csfeedistributor.NewCsFeeDistributorAdapter(networkConfig.WsURL, networkConfig.CSFeeDistributorAddress)
	if err != nil {
		log.Fatalf("Failed to initialize CsFeeDistributor adapter: %v", err)
	}

	// Initialize services
	veboService := services.NewVeboEventsProcessorService(storageAdapter, notifierAdapter, veboAdapter, csFeeDistributorImplAdapter, exitValidatorAdapter, beaconchainAdapter, networkConfig.VeboBlockDeployment)
	csModuleService := services.NewCsmEventsProcessorService(storageAdapter, notifierAdapter, csModuleAdapter)
	csFeeDistributorService := services.NewCsFeeDistributorEventsProcessorService(notifierAdapter, csFeeDistributorAdapter)

	// Start periodic scan for ValidatorExitRequest events
	go veboService.ScanEventsCron(ctx, 1*time.Minute) // TODO: determine interval

	// Start subscribing to each SC event
	if err := csModuleService.WatchCsModuleEvents(ctx); err != nil {
		log.Fatalf("Failed to subscribe to CSModule events: %v", err)
	}
	if err := csFeeDistributorService.WatchCsFeeDistributorEvents(ctx); err != nil {
		log.Fatalf("Failed to subscribe to CSFeeDistributor events: %v", err)
	}

	// Block until context is canceled
	<-ctx.Done()
	log.Println("Shutting down services gracefully.")
}
