package usecase

import (
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
)

type Withdraw interface {
	Execute(input WithdrawInput) error
}

type WithdrawInput struct {
	AccountID string `json:"account_id"`
	AssetID   string `json:"asset_id"`
	Amount    float64
}

type withdraw struct {
	repository repository.WalletRepository
}

func NewWithdraw(repository repository.WalletRepository) *withdraw {
	return &withdraw{repository}
}

func (u *withdraw) Execute(input WithdrawInput) error {
	wallet, err := u.repository.GetByAccountID(input.AccountID)
	if err != nil {
		return err
	}
	if wallet == nil {
		return pkgerror.NewNotFound("Wallet")
	}
	if err = wallet.Withdraw(input.AssetID, input.Amount); err != nil {
		return err
	}
	if err = u.repository.Update(wallet); err != nil {
		return err
	}
	return nil
}
