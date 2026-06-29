package repository

import "github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"

type WalletRepository interface {
	Update(wallet entity.Wallet) error
	GetByAccountID(accountID string) (entity.Wallet, error)
}
