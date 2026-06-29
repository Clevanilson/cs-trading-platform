package usecase

import (
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

type Deposit interface {
	Execute(input DepositInput) error
}

type DepositInput struct {
	AccountID string `json:"account_id"`
	AssetID   string `json:"asset_id"`
	Amount    float64
}

type deposit struct {
	repository repository.WalletRepository
}

func NewDeposit(repository repository.WalletRepository) *deposit {
	return &deposit{repository}
}

func (u *deposit) Execute(input DepositInput) error {
	wallet, err := u.repository.GetByAccountID(input.AccountID)
	if err != nil {
		return err
	}
	if wallet == nil {
		wallet, err = entity.NewWallet(entity.WalletBuilder{AccountID: input.AccountID})
		if err != nil {
			return err
		}
	}
	if err = wallet.Deposit(input.AssetID, input.Amount); err != nil {
		return err
	}
	if err = u.repository.Update(wallet); err != nil {
		return err
	}
	return nil
}
