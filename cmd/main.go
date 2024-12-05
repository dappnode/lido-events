package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"lido-events/internal/adapters/api"
	"lido-events/internal/adapters/beaconchain"
	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"
	csfeedistributorimpl "lido-events/internal/adapters/csFeeDistributorImpl"
	csmodule "lido-events/internal/adapters/csModule"
	"lido-events/internal/adapters/execution"
	exitvalidator "lido-events/internal/adapters/exitValidator"
	"lido-events/internal/adapters/ipfs"
	"lido-events/internal/adapters/notifier"
	proxyapi "lido-events/internal/adapters/proxyApi"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/vebo"
	"lido-events/internal/application/services"
	"lido-events/internal/config"
	"lido-events/internal/logger"
)

func main() {
	// Set up context with cancellation and WaitGroup for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// Load configurations
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		logger.Fatal("Failed to load network configuration: %v", err)
	}
	logger.Debug("Network config: %+v", networkConfig)

	// Initialize adapters
	storageAdapter := storage.NewStorageAdapter()
	notifierAdapter, err := notifier.NewNotifierAdapter(ctx, storageAdapter)
	if err != nil {
		logger.Warn("Telegram notifier not initialized: %v", err)
	}

	apiAdapter := api.NewAPIAdapter(storageAdapter, networkConfig.CORS)
	proxyApiAdapter := proxyapi.NewProxyAPIAdapter(networkConfig.CORS, networkConfig.LidoKeysApiUrl)

	// Initialize API services
	apiService := services.NewAPIServerService(apiAdapter, networkConfig.ApiPort)
	proxyService := services.NewProxyAPIServerService(proxyApiAdapter, networkConfig.ProxyApiPort)

	// Start API services
	apiService.Start(&wg)
	proxyService.Start(&wg)

	// Wait for and validate initial configuration
	if err := waitForConfig(ctx, storageAdapter); err != nil {
		logger.Warn("Application shutting down due to configuration validation failure: %v", err)
		return
	}

	// Initialize domain adapters
	ipfsAdapter := ipfs.NewIPFSAdapter(networkConfig.IpfsUrl)
	beaconchainAdapter := beaconchain.NewBeaconchainAdapter(networkConfig.BeaconchainURL)
	executionAdapter := execution.NewExecutionAdapter(networkConfig.RpcUrl)
	exitValidatorAdapter := exitvalidator.NewExitValidatorAdapter(beaconchainAdapter, networkConfig.SignerUrl)
	csFeeDistributorImplAdapter, _ := csfeedistributorimpl.NewCsFeeDistributorImplAdapter(networkConfig.WsURL, networkConfig.CSFeeDistributorAddress)
	veboAdapter, _ := vebo.NewVeboAdapter(networkConfig.WsURL, networkConfig.VEBOAddress, storageAdapter)
	csModuleAdapter, _ := csmodule.NewCsModuleAdapter(networkConfig.WsURL, networkConfig.CSModuleAddress, storageAdapter)
	csFeeDistributorAdapter, _ := csfeedistributor.NewCsFeeDistributorAdapter(networkConfig.WsURL, networkConfig.CSFeeDistributorAddress)

	// Initialize domain services
	eventsWatcherService := services.NewEventsWatcherService(veboAdapter, csModuleAdapter, csFeeDistributorAdapter, notifierAdapter)
	distributionLogUpdatedScannerService := services.NewDistributionLogUpdatedEventScanner(storageAdapter, notifierAdapter, executionAdapter, csFeeDistributorImplAdapter, networkConfig.CsFeeDistributorBlockDeployment)
	validatorExitRequestScannerService := services.NewValidatorExitRequestEventScanner(storageAdapter, notifierAdapter, veboAdapter, executionAdapter, beaconchainAdapter, networkConfig.VeboBlockDeployment)
	validatorEjectorService := services.NewValidatorEjectorService(storageAdapter, notifierAdapter, exitValidatorAdapter, beaconchainAdapter)
	pendingHashesLoaderService := services.NewPendingHashesLoader(storageAdapter, ipfsAdapter)

	// Start domain services
	distributionLogUpdatedExecutionComplete := make(chan struct{})
	go distributionLogUpdatedScannerService.ScanDistributionLogUpdatedEventsCron(ctx, 384*time.Second, &wg, distributionLogUpdatedExecutionComplete)
	go pendingHashesLoaderService.LoadPendingHashesCron(ctx, 3*time.Hour, &wg, distributionLogUpdatedExecutionComplete)

	exitRequestExecutionComplete := make(chan struct{})
	go validatorExitRequestScannerService.ScanValidatorExitRequestEventsCron(ctx, 384*time.Second, &wg, exitRequestExecutionComplete)
	go validatorEjectorService.ValidatorEjectorCron(ctx, 64*time.Minute, &wg, exitRequestExecutionComplete)

	go eventsWatcherService.WatchAllEvents(ctx, &wg)

	// Handle OS signals for shutdown
	handleShutdown(cancel, apiService, proxyService)

	// Wait for all goroutines to finish
	wg.Wait()
	logger.Info("All services stopped. Shutting down application.")
}

// Helper function to check if operator IDs and Telegram config are available
func waitForConfig(ctx context.Context, storageAdapter *storage.Storage) error {
	for {
		select {
		case <-ctx.Done(): // Exit if the context is canceled
			logger.Info("Context canceled before configuration was ready.")
			return ctx.Err()
		default:
			// Check for operator IDs
			operatorIds, err := storageAdapter.GetOperatorIds()
			if err != nil || len(operatorIds) == 0 {
				logger.Info("Waiting for operator IDs to be set...")
			} else {
				// Operator IDs are set
				logger.Info("Operator IDs are set. Proceeding with initialization.")
				return nil
			}
			time.Sleep(2 * time.Second) // Poll every 2 seconds
		}
	}
}

// handleShutdown manages graceful shutdown for services
func handleShutdown(cancel context.CancelFunc, apiService *services.APIServerService, proxyService *services.ProxyAPIServerService) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		logger.Info("Received shutdown signal. Initiating graceful shutdown...")
		cancel()

		// Shutdown API services with a timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		apiService.Shutdown(shutdownCtx)
		proxyService.Shutdown(shutdownCtx)
	}()
}
