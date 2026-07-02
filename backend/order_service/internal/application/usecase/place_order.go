package usecase

import (
	"encoding/json"

	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
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
	OrderID   string  `json:"order_id"`
	AccountID string  `json:"account_id"`
	MarketID  string  `json:"market_id"`
	Side      string  `json:"side"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
}

type placeOrder struct {
	walletRepository repository.WalletRepository
	orderRepository  repository.OrderRepository
	queue            pkgqueue.Queue
}

func NewPlaceOrder(
	walletRepository repository.WalletRepository,
	orderRepository repository.OrderRepository,
	queue pkgqueue.Queue,
) *placeOrder {
	return &placeOrder{
		walletRepository: walletRepository,
		orderRepository:  orderRepository,
		queue:            queue,
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
	output := PlaceOrderOutput{
		OrderID:   order.ID(),
		AccountID: order.AccountID(),
		MarketID:  order.MarketID(),
		Side:      order.Side(),
		Amount:    order.Amount(),
		Price:     order.Price(),
	}
	event, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	if err := u.queue.Publish("orderPlaced", event); err != nil {
		return nil, err
	}
	return &output, nil
}
