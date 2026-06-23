package entity

import "github.com/clevanilson/cs-trading-platform/pkg/errorc"

type Balance interface {
	AssetID() string
	Amount() uint64
	Deposit(amount int) error
	Withdraw(amount int) error
}

type balance struct {
	assetID string
	amount  uint64
}

type BalanceBuilder struct {
	AssetID string
	Amount  int
}

func NewBalance(builder BalanceBuilder) (*balance, error) {
	if builder.Amount <= 0 {
		return nil, errorc.NewDomain("amount")
	}
	return &balance{
		assetID: builder.AssetID,
		amount:  uint64(builder.Amount),
	}, nil
}

func (b *balance) AssetID() string {
	return b.assetID
}

func (b *balance) Amount() uint64 {
	return b.amount
}

func (b *balance) Deposit(amount int) error {
	if amount <= 0 {
		return errorc.NewDomain("amount")
	}
	b.amount += uint64(amount)
	return nil
}

func (b *balance) Withdraw(amount int) error {
	if amount <= 0 {
		return errorc.NewDomain("amount")
	}
	if b.amount < uint64(amount) {
		return errorc.NewDomain("amount")
	}
	b.amount -= uint64(amount)
	return nil
}
