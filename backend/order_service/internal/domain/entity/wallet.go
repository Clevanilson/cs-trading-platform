package entity

import (
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type Wallet interface {
	AccountID() string
	Deposit(assetID string, amount float64) error
	Withdraw(assetID string, amount float64) error
	GetBalanceByAssetID(assetID string) (Balance, error)
	Balances() []Balance
	LockAmount(order Order) error
	UnlockAmount(order Order) error
}

type wallet struct {
	accountID string
	balances  map[string]Balance
}

type WalletBuilder struct {
	AccountID string
}

func NewWallet(builder WalletBuilder) (*wallet, error) {
	return &wallet{
		accountID: builder.AccountID,
		balances:  make(map[string]Balance),
	}, nil
}

func (a *wallet) AccountID() string {
	return a.accountID
}

func (a *wallet) Deposit(assetID string, amount float64) error {
	currentBalance, err := a.GetBalanceByAssetID(assetID)
	if err == nil {
		if err = currentBalance.Deposit(amount); err != nil {
			return err
		}
		return nil
	}
	balance, err := NewBalance(BalanceBuilder{
		AssetID: assetID,
		Amount:  amount,
	})
	if err != nil {
		return err
	}
	a.balances[assetID] = balance
	return nil
}

func (a *wallet) GetBalanceByAssetID(assetID string) (Balance, error) {
	balance, ok := a.balances[assetID]
	if !ok {
		return nil, pkgerror.NewNotFound("Balance")
	}
	return balance, nil
}

func (a *wallet) Withdraw(assetID string, amount float64) error {
	balance, err := a.GetBalanceByAssetID(assetID)
	if err != nil {
		return err
	}
	if err = balance.Withdraw(amount); err != nil {
		return err
	}
	return nil
}

func (a *wallet) Balances() []Balance {
	balances := make([]Balance, 0)
	for _, balance := range a.balances {
		balances = append(balances, balance)
	}
	return balances
}

func (a *wallet) LockAmount(order Order) error {
	if !a.hasFunds(order) {
		return pkgerror.NewDomain("Insufficient funds")
	}
	if order.Side() == "buy" {
		asset, err := a.GetBalanceByAssetID(order.PaymentAsset())
		if err != nil {
			return err
		}
		if err := asset.LockAmount(order.Price() * order.Amount()); err != nil {
			return nil
		}
	} else {
		asset, err := a.GetBalanceByAssetID(order.MainAsset())
		if err != nil {
			return err
		}
		if err := asset.LockAmount(order.Amount()); err != nil {
			return err
		}
	}
	return nil
}

func (a *wallet) UnlockAmount(order Order) error {
	if order.Side() == "buy" {
		asset, err := a.GetBalanceByAssetID(order.PaymentAsset())
		if err != nil {
			return err
		}
		if err := asset.UnlockAmount(order.Price() * order.Amount()); err != nil {
			return err
		}
	} else {
		asset, err := a.GetBalanceByAssetID(order.MainAsset())
		if err != nil {
			return err
		}
		if err := asset.UnlockAmount(order.Amount()); err != nil {
			return err
		}
	}
	return nil
}

func (a *wallet) hasFunds(order Order) bool {
	var assetId string
	var orderValue float64
	if order.Side() == "buy" {
		assetId = order.PaymentAsset()
		orderValue = order.Price() * order.Amount()
	} else {
		assetId = order.MainAsset()
		orderValue = order.Amount()
	}
	balance, err := a.GetBalanceByAssetID(assetId)
	if err != nil {
		return false
	}
	if balance.Amount() < orderValue {
		return false
	}
	return true
}
