package usecase

import (
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
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
	walletRepository repository.WalletRepository
	orderRepository  repository.OrderRepository
}

func NewPlaceOrder(
	walletRepository repository.WalletRepository,
	orderRepository repository.OrderRepository,
) *placeOrder {
	return &placeOrder{
		walletRepository,
		orderRepository,
	}
}

func (u *placeOrder) Execute(input PlaceOrderInput) (*PlaceOrderOutput, error) {
	wallet, err := u.walletRepository.GetByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, pkgerror.NewNotFound("Wallet")
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
	if err := wallet.LockAmount(order); err != nil {
		return nil, err
	}
	if err := u.walletRepository.Update(wallet); err != nil {
		return nil, err
	}
	if err := u.orderRepository.Save(order); err != nil {
		return nil, err
	}
	return &PlaceOrderOutput{OrderID: order.ID()}, nil
}
