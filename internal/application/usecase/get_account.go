package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type GetAccount struct {
	repository repository.AccountRepository
}

func NewGetAccount(repository repository.AccountRepository) *GetAccount {
	return &GetAccount{repository}
}

func (u *GetAccount) Execute(input GetAccountInput) (*GetAccountOutput, error) {
	account, err := u.repository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errorc.NewNotFound("Account")
	}
	balance := make([]getAccountOutpuBalance, 0)
	for _, _asset := range account.Balances() {
		balance = append(balance, getAccountOutpuBalance{
			AssetID: _asset.AssetID(),
			Amount:  _asset.Amount(),
		})
	}
	return &GetAccountOutput{
		ID:      account.ID(),
		Name:    account.Name(),
		Balance: balance,
	}, nil
}

type GetAccountInput struct {
	ID string `json:"id"`
}

type GetAccountOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance []getAccountOutpuBalance
}

type getAccountOutpuBalance struct {
	AssetID string  `json:"asset_id"`
	Amount  float64 `json:"amount"`
}
