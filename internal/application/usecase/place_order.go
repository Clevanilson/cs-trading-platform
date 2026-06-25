package usecase

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

type PlaceOrder interface {
	Execute(input PlaceOrderInput) (*PlaceOrderOutput, error)
}

type PlaceOrderInput struct {
	AccountID string  `json:"account_id"`
	MarketID  string  `json:"market_id"`
	Side      string  `json:"side"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
}

type PlaceOrderOutput struct {
	OrderID string `json:"order_id"`
}

type placeOrder struct {
	accountRepository repository.AccountRepository
	orderRepository   repository.OrderRepository
}

func NewPlaceOrder(
	accountRepository repository.AccountRepository,
	orderRepository repository.OrderRepository,
) *placeOrder {
	return &placeOrder{
		accountRepository,
		orderRepository,
	}
}

func (u *placeOrder) Execute(input PlaceOrderInput) (*PlaceOrderOutput, error) {
	account, err := u.accountRepository.GetByID(input.AccountID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errorc.NewNotFound("Account")
	}
	order, err := entity.NewOrder(entity.OrderBuilder{
		AccountID: input.AccountID,
		MarketID:  input.MarketID,
		Side:      input.Side,
		Price:     input.Price,
		Amount:    input.Amount,
	})
	if err != nil {
		return nil, err
	}
	if err := account.LockAmount(order); err != nil {
		return nil, err
	}
	if err := u.accountRepository.Save(account); err != nil {
		return nil, err
	}
	if err := u.orderRepository.Save(order); err != nil {
		return nil, err
	}
	return &PlaceOrderOutput{OrderID: order.ID()}, nil
}
