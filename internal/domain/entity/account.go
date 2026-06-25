package entity

import (
	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type Account interface {
	Name() string
	ID() string
	Deposit(assetID string, amount float64) error
	Withdraw(assetID string, amount float64) error
	GetBalanceByAssetID(assetID string) (Balance, error)
	Balances() []Balance
	LockAmount(order Order) error
	UnlockAmount(order Order) error
}

type account struct {
	name     valueobject.Name
	id       valueobject.ID
	balances map[string]Balance
}

type AccountBuilder struct {
	Name string
	ID   *string
}

func NewAccount(builder AccountBuilder) (*account, error) {
	name, err := valueobject.NewName(builder.Name)
	if err != nil {
		return nil, err
	}
	return &account{
		name:     name,
		id:       *valueobject.NewID(builder.ID),
		balances: make(map[string]Balance),
	}, nil
}

func (a *account) Name() string {
	return a.name.Value()
}

func (a *account) ID() string {
	return a.id.Value()
}

func (a *account) Deposit(assetID string, amount float64) error {
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

func (a *account) GetBalanceByAssetID(assetID string) (Balance, error) {
	balance, ok := a.balances[assetID]
	if !ok {
		return nil, errorc.NewNotFound("Balance")
	}
	return balance, nil
}

func (a *account) Withdraw(assetID string, amount float64) error {
	balance, err := a.GetBalanceByAssetID(assetID)
	if err != nil {
		return err
	}
	if err = balance.Withdraw(amount); err != nil {
		return err
	}
	return nil
}

func (a *account) Balances() []Balance {
	balances := make([]Balance, 0)
	for _, balance := range a.balances {
		balances = append(balances, balance)
	}
	return balances
}

func (a *account) LockAmount(order Order) error {
	if !a.hasFunds(order) {
		return errorc.NewDomain("Insufficient funds")
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

func (a *account) UnlockAmount(order Order) error {
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

func (a *account) hasFunds(order Order) bool {
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
