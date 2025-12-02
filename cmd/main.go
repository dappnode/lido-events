package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"lido-events/internal/adapters/beaconchain"
	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"
	"lido-events/internal/adapters/execution"
	exitvalidator "lido-events/internal/adapters/exitValidator"
	"lido-events/internal/adapters/ipfs"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/relays"
	"lido-events/internal/adapters/storage/exits"
	"lido-events/internal/adapters/storage/performance"
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

	// Initiate RPC Ethereum clients
	rpcClient, err := ethclient.Dial(networkConfig.RpcUrl)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize RPC client: %v", err)
	}

	exitsStorage, err := exits.NewAdapter(networkConfig.DBDirectory)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize storage adapter: %v", err)
	}
	performanceStorage, err := performance.NewAdapter(networkConfig.PerformanceDBPath)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize performance storage adapter: %v", err)
	}
	notifier := notifier.NewNotifier(networkConfig.Network, networkConfig.LidoDnpName)

	relays, err := relays.NewAdapter(rpcClient, networkConfig.MEVBoostRelaysAllowListAddres, networkConfig.DappmanagerUrl, networkConfig.MevBoostDnpName)
	if err != nil {
		logger.Fatal("Failed to initialize relays: %v", err)
	}
	ipfsAdapter := ipfs.NewIPFSAdapter(networkConfig.IpfsUrl)
	beaconchainAdapter := beaconchain.NewBeaconchainAdapter(networkConfig.BeaconchainURL)
	executionAdapter := execution.NewExecutionAdapter(networkConfig.RpcUrl)
	exitValidatorAdapter := exitvalidator.NewExitValidatorAdapter(beaconchainAdapter, networkConfig.SignerUrl)
	csFeeDistributorAdapter, err := csfeedistributor.NewCsFeeDistributorAdapter(rpcClient, networkConfig.CSFeeDistributorProxyAddress)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize CsFeeDistributorAdapter: %v", err)
	}
	veboAdapter, err := vebo.NewVeboAdapter(rpcClient, networkConfig.VEBOAddress, exitsStorage, networkConfig.BlockChunkSize, networkConfig.CSMStakingModuleID)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize VeboAdapter: %v", err)
	}

	// Initialize services
	validatorExitRequestScannerService := services.NewValidatorExitRequestEventScanner(exitsStorage, notifier, veboAdapter, executionAdapter, beaconchainAdapter, networkConfig.VeboBlockDeployment, networkConfig.CSModuleTxReceipt)
	validatorEjectorService := services.NewValidatorEjectorService(networkConfig.BeaconchaUrl, exitsStorage, notifier, exitValidatorAdapter, beaconchainAdapter)
	pendingHashesLoaderService := services.NewAllHashesLoader(performanceStorage, notifier, csFeeDistributorAdapter, ipfsAdapter)
	apiService := services.NewAPIServerService(ctx, networkConfig.ApiPort, exitsStorage, performanceStorage, notifier, relays, validatorExitRequestScannerService, networkConfig.CORS)
	proxyService := services.NewProxyAPIServerService(networkConfig.ProxyApiPort, networkConfig.LidoKeysApiUrl, networkConfig.CORS)
	relaysCheckerService := services.NewRelayCronService(networkConfig.StakersUiUrl, relays, notifier)

	// Start servers
	apiService.Start(&wg)
	proxyService.Start(&wg)

	// Start services
	go relaysCheckerService.StartRelayMonitoringCron(ctx, 5*time.Minute, &wg)
	go pendingHashesLoaderService.LoadHashesCron(ctx, 3*time.Hour, &wg)
	exitRequestExecutionComplete := make(chan struct{})
	go validatorExitRequestScannerService.ScanValidatorExitRequestEventsCron(ctx, 384*time.Second, &wg, exitRequestExecutionComplete)
	go validatorEjectorService.ValidatorEjectorCron(ctx, 64*time.Minute, &wg, exitRequestExecutionComplete)

	// Handle OS signals for shutdown
	handleShutdown(cancel, apiService, proxyService)

	// Wait for all goroutines to finish
	wg.Wait()
	logger.InfoWithPrefix(logPrefix, "All services stopped. Shutting down application.")
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
