package entity

import (
	"regexp"
	"strings"

	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type Order interface {
	ID() string
	MarketID() string
	Side() string
	Price() float64
	Amount() float64
	MainAsset() string
	PaymentAsset() string
}

type OrderBuilder struct {
	ID        *string
	AccountID string
	MarketID  string
	Side      string
	Price     float64
	Amount    float64
}

type order struct {
	id           valueobject.ID
	marketID     string
	accountID    string
	side         string
	price        float64
	amount       float64
	mainAsset    string
	paymentAsset string
}

func NewOrder(builder OrderBuilder) (*order, error) {
	assets := strings.Split(builder.MarketID, "-")
	if builder.Price < 0 {
		return nil, errorc.NewDomain("Invalid price")
	}
	if builder.Amount < 0 {
		return nil, errorc.NewDomain("Invalid amount")
	}
	if builder.Side != "buy" && builder.Side != "sell" {
		return nil, errorc.NewDomain("Invalid side")
	}
	if builder.AccountID == "" {
		return nil, errorc.NewDomain("Invalid accountID")
	}
	if !regexp.MustCompile(`[A-Z]{3}-[A-Z]{3}`).MatchString(builder.MarketID) {
		return nil, errorc.NewDomain("Invalid marketID")
	}
	return &order{
		id:           valueobject.NewID(builder.ID),
		marketID:     builder.MarketID,
		accountID:    builder.AccountID,
		side:         builder.Side,
		price:        builder.Price,
		amount:       builder.Amount,
		mainAsset:    assets[0],
		paymentAsset: assets[1],
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
	return o.amount
}

func (o *order) MainAsset() string {
	return o.mainAsset
}

func (o *order) PaymentAsset() string {
	return o.paymentAsset
}
