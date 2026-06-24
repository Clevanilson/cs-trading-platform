package entity

import (
	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type Account interface {
	Name() string
	ID() string
	Deposit(assetID string, amount int) error
	Withdraw(assetID string, amount int) error
	GetBalanceByAssetID(assetID string) (Balance, error)
	Balances() []Balance
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

func (a *account) Deposit(assetID string, amount int) error {
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

func (a *account) Withdraw(assetID string, amount int) error {
	balance, err := a.GetBalanceByAssetID(assetID)
	if err != nil {
		return err
	}
	err = balance.Withdraw(amount)
	if err != nil {
		return err
	}
	if balance.Amount() == 0 {
		delete(a.balances, assetID)
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
