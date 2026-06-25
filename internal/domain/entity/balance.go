package entity

import "github.com/clevanilson/cs-trading-platform/pkg/errorc"

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
		return nil, errorc.NewDomain("Invalid amount")
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
		return errorc.NewDomain("Invalid amount")
	}
	b.amount += amount
	return nil
}

func (b *balance) Withdraw(amount float64) error {
	if amount <= 0 {
		return errorc.NewDomain("Invalid amount")
	}
	if b.amount < amount {
		return errorc.NewDomain("Invalid amount")
	}
	b.amount -= amount
	return nil
}

func (b *balance) LockAmount(value float64) error {
	if value < 0 {
		return errorc.NewDomain("Invalid amount")
	}
	if value > b.amount {
		return errorc.NewDomain("Insufficient funds")
	}
	b.lockedAmount += value
	return nil
}

func (b *balance) UnlockAmount(value float64) error {
	if value < 0 {
		return errorc.NewDomain("Invalid amount")
	}
	if value > b.lockedAmount {
		return errorc.NewDomain("Invalid amount")
	}
	b.lockedAmount -= value
	return nil
}
