package config

import (
	"lido-events/internal/application/domain"
	"math/big"
	"sync"
)

type ConfigManager struct {
	mu                  sync.RWMutex
	operatorIds         []*big.Int
	operatorIdsUpdateCh chan []*big.Int

	telegramConfig         domain.TelegramConfig
	telegramConfigUpdateCh chan domain.TelegramConfig
}

func NewConfigManager(operatorIds []*big.Int, telegramConfig domain.TelegramConfig) *ConfigManager {
	return &ConfigManager{
		operatorIds:            operatorIds,
		operatorIdsUpdateCh:    make(chan []*big.Int, 1),
		telegramConfig:         telegramConfig,
		telegramConfigUpdateCh: make(chan domain.TelegramConfig, 1),
	}
}

// Methods for operatorIds
func (cm *ConfigManager) GetOperatorIds() []*big.Int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.operatorIds
}

func (cm *ConfigManager) UpdateOperatorIds(newOperatorIds []*big.Int) {
	cm.mu.Lock()
	cm.operatorIds = newOperatorIds
	cm.mu.Unlock()
	cm.operatorIdsUpdateCh <- newOperatorIds
}

func (cm *ConfigManager) GetOperatorIdsUpdateChannel() <-chan []*big.Int {
	return cm.operatorIdsUpdateCh
}

// Methods for telegramConfig
func (cm *ConfigManager) GetTelegramConfig() domain.TelegramConfig {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.telegramConfig
}

func (cm *ConfigManager) UpdateTelegramConfig(newConfig domain.TelegramConfig) {
	cm.mu.Lock()
	cm.telegramConfig = newConfig
	cm.mu.Unlock()
	cm.telegramConfigUpdateCh <- newConfig
}

func (cm *ConfigManager) GetTelegramConfigUpdateChannel() <-chan domain.TelegramConfig {
	return cm.telegramConfigUpdateCh
}
