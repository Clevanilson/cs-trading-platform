package entity

import (
	"regexp"
	"strings"
	"time"

	pkgvalueobject "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/value_object"
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type Order interface {
	ID() string
	MarketID() string
	Side() string
	Price() float64
	Amount() float64
	MainAsset() string
	PaymentAsset() string
	Status() string
	CreatedAt() time.Time
	Fill(amount float64, price float64) error
}

type OrderBuilder struct {
	ID        *string
	AccountID string
	MarketID  string
	Side      string
	Price     float64
	Amount    float64
	Status    string
	CreatedAt time.Time
}

type order struct {
	id           pkgvalueobject.ID
	marketID     string
	accountID    string
	side         string
	price        float64
	amount       float64
	mainAsset    string
	paymentAsset string
	status       string
	createdAt    time.Time
	filledAmount float64
	filledPrice  float64
}

func NewOrder(builder OrderBuilder) (*order, error) {
	assets := strings.Split(builder.MarketID, "-")
	if builder.Price < 0 {
		return nil, pkgerror.NewDomain("Invalid price")
	}
	if builder.Amount < 1 {
		return nil, pkgerror.NewDomain("Invalid amount")
	}
	if builder.Side != "buy" && builder.Side != "sell" {
		return nil, pkgerror.NewDomain("Invalid side")
	}
	if builder.AccountID == "" {
		return nil, pkgerror.NewDomain("Invalid accountID")
	}
	if !regexp.MustCompile(`[A-Z]{3}-[A-Z]{3}`).MatchString(builder.MarketID) {
		return nil, pkgerror.NewDomain("Invalid marketID")
	}
	if builder.Status == "" {
		builder.Status = "open"
	}
	if builder.Status != "open" && builder.Status != "closed" && builder.Status != "canceled" {
		return nil, pkgerror.NewDomain("Invalid status")
	}
	if builder.CreatedAt.IsZero() {
		builder.CreatedAt = time.Now()
	}
	return &order{
		id:           pkgvalueobject.NewID(builder.ID),
		marketID:     builder.MarketID,
		accountID:    builder.AccountID,
		side:         builder.Side,
		price:        builder.Price,
		amount:       builder.Amount,
		mainAsset:    assets[0],
		paymentAsset: assets[1],
		status:       builder.Status,
		createdAt:    builder.CreatedAt,
		filledAmount: 0,
		filledPrice:  0,
	}, nil
}

func (o *order) ID() string {
	return o.id.Value()
}

func (o *order) MarketID() string {
	return o.marketID
}

func (o *order) AccountID() string {
	return o.accountID
}

func (o *order) Side() string {
	return o.side
}

func (o *order) Price() float64 {
	return o.price
}

func (o *order) Amount() float64 {
	return o.amount - o.filledAmount
}

func (o *order) MainAsset() string {
	return o.mainAsset
}

func (o *order) PaymentAsset() string {
	return o.paymentAsset
}

func (o *order) Status() string {
	return o.status
}

func (o *order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *order) Fill(amount float64, price float64) error {
	if amount <= 0 {
		return pkgerror.NewDomain("Invalid amount")
	}
	if price <= 0 {
		return pkgerror.NewDomain("Invalid price")
	}
	o.filledAmount += amount
	o.filledPrice = (o.filledPrice*o.filledAmount + price*amount) / (o.filledAmount + amount)
	return nil
}
