package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"lido-events/internal/adapters/beaconchain"
	"lido-events/internal/adapters/csfeedistributor"
	"lido-events/internal/adapters/csparameters"
	"lido-events/internal/adapters/execution"
	"lido-events/internal/adapters/exitvalidator"
	"lido-events/internal/adapters/ipfs"
	"lido-events/internal/adapters/notifier"
	"lido-events/internal/adapters/relays"
	"lido-events/internal/adapters/storage/exits"
	"lido-events/internal/adapters/storage/performance"
	"lido-events/internal/adapters/vebo"
	"lido-events/internal/application/domain"
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
	config, err := config.LoadNetworkConfig()
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to load network configuration: %v", err)
	}

	// Initialize domain-level notification identifiers based on network
	domain.InitNotifications(config.Network)

	// Initiate RPC Ethereum clients
	rpcClient, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize RPC client: %v", err)
	}

	exitsStorage, err := exits.NewJson(config.DBDirectory)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize storage adapter: %v", err)
	}
	performanceStorage, err := performance.NewPerformance(config.DBDirectory)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize performance storage adapter: %v", err)
	}
	notifier := notifier.NewNotifier(ctx, config.Network, config.DappmanagerUrl, config.LidoDnpName, config.BrainUrl, config.StakersUiUrl, config.BeaconchaUrl)

	relays, err := relays.NewARelays(rpcClient, config.MEVBoostRelaysAllowListAddres, config.DappmanagerUrl, config.MevBoostDnpName)
	if err != nil {
		logger.Fatal("Failed to initialize relays: %v", err)
	}
	ipfs := ipfs.NewIPFS(config.IpfsUrl)
	beaconchain := beaconchain.NewBeaconchain(config.BeaconchainURL)
	execution := execution.NewExecution(config.RpcUrl)
	exitValidator := exitvalidator.NewExitValidator(beaconchain, config.SignerUrl)
	csFeeDistributor, err := csfeedistributor.NewCsFeeDistributor(rpcClient, config.CSFeeDistributorProxyAddress)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize CsFeeDistributorAdapter: %v", err)
	}
	csParameters, err := csparameters.NewCsParameters(rpcClient, config.CSParametersRegistryAddress)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize CsParametersAdapter: %v", err)
	}
	vebo, err := vebo.NewVeboAdapter(rpcClient, config.VEBOAddress, exitsStorage, config.BlockChunkSize, config.CSMStakingModuleID)
	if err != nil {
		logger.FatalWithPrefix(logPrefix, "Failed to initialize VeboAdapter: %v", err)
	}

	// Initialize services
	validatorExitRequestScannerService := services.NewExitRequestEventScanner(exitsStorage, notifier, vebo, execution, beaconchain, csParameters, config.SecondsPerSlot, config.DefaultAllowedExitDelay, config.ExitDelayMultiplier)
	validatorEjectorService := services.NewValidatorEjectorService(exitsStorage, notifier, exitValidator, beaconchain)
	pendingHashesLoaderService := services.NewAllHashesLoader(performanceStorage, notifier, csFeeDistributor, ipfs)
	apiService := services.NewAPIServerService(ctx, config.ApiPort, exitsStorage, performanceStorage, relays, validatorExitRequestScannerService, config.CORS)
	proxyService := services.NewProxyAPIServerService(config.ProxyApiPort, config.LidoKeysApiUrl, config.CORS)
	relaysCheckerService := services.NewRelayCronService(config.StakersUiUrl, relays, notifier)

	// Start servers
	apiService.Start(&wg)
	proxyService.Start(&wg)

	// Start services
	go relaysCheckerService.StartRelayMonitoringCron(ctx, 1*time.Hour, &wg)
	go pendingHashesLoaderService.LoadHashesCron(ctx, 24*time.Hour, &wg)
	go validatorExitRequestScannerService.ScanExitRequestEventsCron(ctx, 384*time.Second, &wg)
	go validatorEjectorService.ValidatorEjectorCron(ctx, 64*time.Minute, &wg)

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
