package infrarepository

import (
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

type walletMemoryRepository struct {
	wallets map[string]entity.Wallet
}

func NewWalletMemoryRepository() *walletMemoryRepository {
	return &walletMemoryRepository{
		wallets: make(map[string]entity.Wallet),
	}
}

func (r *walletMemoryRepository) Update(wallet entity.Wallet) error {
	r.wallets[wallet.AccountID()] = wallet
	return nil
}

func (r *walletMemoryRepository) GetByAccountID(accountID string) (entity.Wallet, error) {
	wallet, ok := r.wallets[accountID]
	if !ok {
		return nil, nil
	}
	return wallet, nil
}
