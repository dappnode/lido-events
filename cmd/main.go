package main

import (
	"context"
	"lido-events/internal/adapters/api"
	"lido-events/internal/adapters/beaconchain"
	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"
	csfeedistributorimpl "lido-events/internal/adapters/csFeeDistributorImpl"
	csmodule "lido-events/internal/adapters/csModule"
	"lido-events/internal/adapters/execution"
	exitvalidator "lido-events/internal/adapters/exitValidator"
	"lido-events/internal/adapters/ipfs"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/vebo"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"lido-events/internal/application/services"
	"lido-events/internal/config"
	"log"
	"net/http"
)

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
	// Set up context with cancellation and a WaitGroup for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// Set up signal channel to handle OS interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Load configurations
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		log.Fatalf("Failed to load network configuration: %v", err)
	}
	log.Printf("Network config: %+v", networkConfig)

	// Initialize adapters
	storageAdapter := storage.NewStorageAdapter()
	apiAdapter := api.NewAPIAdapter(storageAdapter)

	// Start HTTP server
	server := &http.Server{Addr: ":8080", Handler: apiAdapter.Router}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Wait for initial configuration to be set
	waitForInitialConfig(storageAdapter)

	ipfsAdapter := ipfs.NewIPFSAdapter(networkConfig.IpfsUrl)
	beaconchainAdapter := beaconchain.NewBeaconchainAdapter(networkConfig.BeaconchainURL)
	executionAdapter := execution.NewExecutionAdapter(networkConfig.RpcUrl)
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
	eventsWatcherService := services.NewEventsWatcherService(veboAdapter, csModuleAdapter, csFeeDistributorAdapter, notifierAdapter)
	distributionLogUpdatedScannerService := services.NewDistributionLogUpdatedEventScanner(storageAdapter, notifierAdapter, executionAdapter, csFeeDistributorImplAdapter, networkConfig.CsFeeDistributorBlockDeployment)
	validatorExitRequestScannerService := services.NewValidatorExitRequestEventScanner(storageAdapter, notifierAdapter, veboAdapter, executionAdapter, beaconchainAdapter, networkConfig.VeboBlockDeployment)
	validatorEjectorService := services.NewValidatorEjectorService(storageAdapter, notifierAdapter, exitValidatorAdapter, beaconchainAdapter)
	pendingHashesLoaderService := services.NewPendingHashesLoader(storageAdapter, ipfsAdapter)

	// Start background services
	wg.Add(5)
	go func() {
		defer wg.Done()
		distributionLogUpdatedScannerService.ScanDistributionLogUpdatedEventsCron(ctx, 1*time.Minute)
	}()
	go func() {
		defer wg.Done()
		validatorExitRequestScannerService.ScanValidatorExitRequestEventsCron(ctx, 1*time.Minute)
	}()
	go func() {
		defer wg.Done()
		validatorEjectorService.ValidatorEjectorCron(ctx, 10*time.Minute)
	}()
	go func() {
		defer wg.Done()
		pendingHashesLoaderService.LoadPendingHashesCron(ctx, 1*time.Minute)
	}()
	go func() {
		defer wg.Done()
		if err := eventsWatcherService.WatchCsModuleEvents(ctx); err != nil {
			log.Fatalf("Failed to subscribe to CSModule events: %v", err)
		}
	}()

	// Handle shutdown signals
	go func() {
		<-signalChan
		log.Println("Received shutdown signal. Initiating graceful shutdown...")
		cancel() // Cancel context to signal all services to stop

		// Give the HTTP server time to finish ongoing requests
		serverCtx, serverCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer serverCancel()
		if err := server.Shutdown(serverCtx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All services stopped. Shutting down application.")
}
