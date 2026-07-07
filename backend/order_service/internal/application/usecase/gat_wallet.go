package usecase

import (
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

type GetWallet interface {
	Execute(input GetWalletInput) (*GetWalletOutput, error)
}

type GetWalletInput struct {
	AccountID string `json:"account_id"`
}

type GetWalletOutput struct {
	Balances []getWalletBalance `json:"balances"`
}

type getWalletBalance struct {
	AssetID string `json:"asset_id"`
	Amount float64 `json:"amount"`
}

type getWallet struct {
	repository repository.WalletRepository
}

func NewGetWallet(repository repository.WalletRepository) *getWallet {
	return &getWallet{repository}
}

func (u *getWallet) Execute(input GetWalletInput) (*GetWalletOutput, error) {
	wallet, err := u.repository.GetByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		wallet, err = entity.NewWallet(entity.WalletBuilder{AccountID: input.AccountID})
		if err := u.repository.Update(wallet); err != nil {
			return nil, err
		}
	}
	balances := make([]getWalletBalance, len(wallet.Balances()))
	for i, balance := range wallet.Balances() {
		balances[i] = getWalletBalance{
			AssetID: balance.AssetID(),
			Amount: balance.Amount(),
		}
	}
	return &GetWalletOutput{Balances: balances}, nil
}