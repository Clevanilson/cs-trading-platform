package entity

import (
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type Balance interface {
	AssetID() string
	Amount() float64
	Deposit(amount float64) error
	Withdraw(amount float64) error
	LockAmount(amount float64) error
	UnlockAmount(amount float64) error
}

type balance struct {
	assetID      string
	amount       float64
	lockedAmount float64
}

type BalanceBuilder struct {
	AssetID string
	Amount  float64
}

func NewBalance(builder BalanceBuilder) (*balance, error) {
	if builder.Amount <= 0 {
		return nil, pkgerror.NewDomain("Invalid amount")
	}
	if builder.AssetID == "" {
		return nil, pkgerror.NewDomain("Invalid asset ID")
	}
	return &balance{
		assetID:      builder.AssetID,
		amount:       builder.Amount,
		lockedAmount: 0,
	}, nil
}

func (b *balance) AssetID() string {
	return b.assetID
}

func (b *balance) Amount() float64 {
	return b.amount - b.lockedAmount
}

func (b *balance) Deposit(amount float64) error {
	if amount <= 0 {
		return pkgerror.NewDomain("Invalid amount")
	}
	b.amount += amount
	return nil
}

func (b *balance) Withdraw(amount float64) error {
	if amount <= 0 {
		return pkgerror.NewDomain("Invalid amount")
	}
	if b.amount < amount {
		return pkgerror.NewDomain("Invalid amount")
	}
	b.amount -= amount
	return nil
}

func (b *balance) LockAmount(value float64) error {
	if value < 0 {
		return pkgerror.NewDomain("Invalid amount")
	}
	if value > b.amount {
		return pkgerror.NewDomain("Insufficient funds")
	}
	b.lockedAmount += value
	return nil
}

func (b *balance) UnlockAmount(value float64) error {
	if value < 0 {
		return pkgerror.NewDomain("Invalid amount")
	}
	if value > b.lockedAmount {
		return pkgerror.NewDomain("Invalid amount")
	}
	b.lockedAmount -= value
	return nil
}
