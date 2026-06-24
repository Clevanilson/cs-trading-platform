package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
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
	repository repository.AccountRepository
}

func NewDeposit(repository repository.AccountRepository) *deposit {
	return &deposit{repository}
}

func (u *deposit) Execute(input DepositInput) error {
	account, err := u.repository.GetByID(input.AccountID)
	if err != nil {
		return err
	}
	if account == nil {
		return errorc.NewNotFound("account")
	}
	if err = account.Deposit(input.AssetID, input.Amount); err != nil {
		return err
	}
	if err = u.repository.Update(account); err != nil {
		return err
	}
	return nil
}
