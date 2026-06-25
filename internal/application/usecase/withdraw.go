package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
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
	repository repository.AccountRepository
}

func NewWithdraw(repository repository.AccountRepository) *withdraw {
	return &withdraw{repository}
}

func (u *withdraw) Execute(input WithdrawInput) error {
	account, err := u.repository.GetByID(input.AccountID)
	if err != nil {
		return err
	}
	if account == nil {
		return errorc.NewNotFound("Account")
	}
	if account == nil {
		return errorc.NewNotFound("Account")
	}
	if err = account.Withdraw(input.AssetID, input.Amount); err != nil {
		return err
	}
	if err = u.repository.Update(account); err != nil {
		return err
	}
	return nil
}
