package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"lido-events/internal/adapters/beaconchain"
	csfeedistributorimpl "lido-events/internal/adapters/csFeeDistributorImpl"
	"lido-events/internal/adapters/execution"
	exitvalidator "lido-events/internal/adapters/exitValidator"
	"lido-events/internal/adapters/ipfs"
	"lido-events/internal/adapters/notifier"
	relaysallowed "lido-events/internal/adapters/relaysAllowed"
	relaysused "lido-events/internal/adapters/relaysUsed"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/adapters/vebo"
	"lido-events/internal/application/services"
	"lido-events/internal/config"
	"lido-events/internal/logger"

	"github.com/ethereum/go-ethereum/ethclient"
)

var logPrefix = "MAIN"

func main() {
	// Set up context with cancellation and WaitGroup for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// Load configurations
	networkConfig, err := config.LoadNetworkConfig()
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to load network configuration: %v", err)
	}
	logger.DebugWithPrefix(logPrefix, "Network config: %+v", networkConfig)

	// Initiate RPC and Ws Ethereum clients
	rpcClient, err := ethclient.Dial(networkConfig.RpcUrl)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize RPC client: %v", err)
	}
	wsClient, err := ethclient.Dial(networkConfig.WsURL)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize WS client: %v", err)
	}

	storageAdapter, err := storage.NewStorageAdapter(networkConfig.DBDirectory)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize storage adapter: %v", err)
	}
	notifier := notifier.NewNotifier(networkConfig.Network, networkConfig.LidoDnpName)

	relaysUsedAdapter := relaysused.NewRelaysUsedAdapter(networkConfig.DappmanagerUrl, networkConfig.MevBoostDnpName)
	relaysAllowedAdapter, err := relaysallowed.NewRelaysAllowedAdapter(rpcClient, networkConfig.MEVBoostRelaysAllowListAddres, networkConfig.DappmanagerUrl, networkConfig.MevBoostDnpName)
	if err != nil {
		logger.Fatal("Failed to initialize relaysAllowedAdapter: %v", err)
	}
	ipfsAdapter := ipfs.NewIPFSAdapter(networkConfig.IpfsUrl)
	beaconchainAdapter := beaconchain.NewBeaconchainAdapter(networkConfig.BeaconchainURL)
	executionAdapter := execution.NewExecutionAdapter(networkConfig.RpcUrl)
	exitValidatorAdapter := exitvalidator.NewExitValidatorAdapter(beaconchainAdapter, networkConfig.SignerUrl)
	csFeeDistributorImplAdapter, err := csfeedistributorimpl.NewCsFeeDistributorImplAdapter(rpcClient, networkConfig.CSFeeDistributorAddress, networkConfig.BlockChunkSize)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize CsFeeDistributorImplAdapter: %v", err)
	}
	veboAdapter, err := vebo.NewVeboAdapter(wsClient, rpcClient, networkConfig.VEBOAddress, storageAdapter, networkConfig.BlockChunkSize, networkConfig.CSMStakingModuleID)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize VeboAdapter: %v", err)
	}

	// Initialize domain services
	distributionLogUpdatedScannerService := services.NewDistributionLogUpdatedEventScanner(storageAdapter, notifier, executionAdapter, csFeeDistributorImplAdapter, networkConfig.CsFeeDistributorBlockDeployment, networkConfig.CSModuleTxReceipt)
	validatorExitRequestScannerService := services.NewValidatorExitRequestEventScanner(storageAdapter, notifier, veboAdapter, executionAdapter, beaconchainAdapter, networkConfig.VeboBlockDeployment, networkConfig.CSModuleTxReceipt)
	validatorEjectorService := services.NewValidatorEjectorService(networkConfig.BeaconchaUrl, storageAdapter, notifier, exitValidatorAdapter, beaconchainAdapter)
	pendingHashesLoaderService := services.NewPendingHashesLoader(storageAdapter, notifier, ipfsAdapter, networkConfig.MinGenesisTime)
	// Initialize API services
	apiService := services.NewAPIServerService(ctx, networkConfig.ApiPort, storageAdapter, notifier, relaysUsedAdapter, relaysAllowedAdapter, distributionLogUpdatedScannerService, validatorExitRequestScannerService, pendingHashesLoaderService, networkConfig.CORS)
	proxyService := services.NewProxyAPIServerService(networkConfig.ProxyApiPort, networkConfig.LidoKeysApiUrl, networkConfig.CORS)

	// Start API services
	apiService.Start(&wg)
	proxyService.Start(&wg)

	// Wait for and validate initial configuration
	if err := waitForConfig(ctx, storageAdapter); err != nil {
		logger.FatalWithPrefix(logPrefix, "Application shutting down due to configuration validation failure: %v", err)
	}

	// Start domain services
	// go relaysCheckerService.StartRelayMonitoringCron(ctx, 5*time.Minute, &wg)

	distributionLogUpdatedExecutionComplete := make(chan struct{})
	go distributionLogUpdatedScannerService.ScanDistributionLogUpdatedEventsCron(ctx, 384*time.Second, &wg, distributionLogUpdatedExecutionComplete)
	go pendingHashesLoaderService.LoadPendingHashesCron(ctx, 3*time.Hour, &wg, distributionLogUpdatedExecutionComplete)

	exitRequestExecutionComplete := make(chan struct{})
	go validatorExitRequestScannerService.ScanValidatorExitRequestEventsCron(ctx, 384*time.Second, &wg, exitRequestExecutionComplete)
	go validatorEjectorService.ValidatorEjectorCron(ctx, 64*time.Minute, &wg, exitRequestExecutionComplete)

	// Handle OS signals for shutdown
	handleShutdown(cancel, apiService, proxyService)

	// Wait for all goroutines to finish
	wg.Wait()
	logger.InfoWithPrefix(logPrefix, "All services stopped. Shutting down application.")
}

// Helper function to check if operator IDs and Telegram config are available
func waitForConfig(ctx context.Context, storageAdapter *storage.Storage) error {
	for {
		select {
		case <-ctx.Done(): // Exit if the context is canceled
			logger.InfoWithPrefix(logPrefix, "Context canceled before configuration was ready.")
			return ctx.Err()
		default:
			// Check for operator IDs
			operatorIds, err := storageAdapter.GetOperatorIds()
			if err != nil || len(operatorIds) == 0 {
				logger.InfoWithPrefix(logPrefix, "Waiting for operator IDs to be set...")
			} else {
				// Operator IDs are set
				logger.InfoWithPrefix(logPrefix, "Operator IDs are set. Proceeding with initialization.")
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
		logger.InfoWithPrefix(logPrefix, "Received shutdown signal. Initiating graceful shutdown...")
		cancel()

		// Shutdown API services with a timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		apiService.Shutdown(shutdownCtx)
		proxyService.Shutdown(shutdownCtx)
	}()
}
