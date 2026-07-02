package usecase_test

import (
	"encoding/json"
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func TestPlaceOrder(t *testing.T) {
	var wallet entity.Wallet
	var walletRepository repository.WalletRepository
	var orderRepository repository.OrderRepository
	var queue *pkgqueue.MockQueue
	var sut usecase.PlaceOrder

	setup := func() {
		var err error
		queue = pkgqueue.NewMockQueue()
		wallet, err = entity.NewWallet(entity.WalletBuilder{AccountID: "Lune"})
		err = wallet.Deposit("USD", 100_000)
		walletRepository = infrarepository.NewWalletMemoryRepository()
		err = walletRepository.Update(wallet)
		orderRepository = infrarepository.NewOrderMemoryRepository()
		sut = usecase.NewPlaceOrder(walletRepository, orderRepository, queue)
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: wallet.AccountID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     78_000,
		}
		output, err := sut.Execute(input)
		outputJSON, err := json.Marshal(output)
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, output, nil)
		pkgassert.NotEquals(t, output, nil)
		pkgassert.Equals(t, output.OrderID, output.OrderID)
		pkgassert.Equals(t, output.MarketID, input.MarketID)
		pkgassert.Equals(t, output.Side, input.Side)
		pkgassert.Equals(t, output.Amount, input.Amount)
		pkgassert.Equals(t, output.Price, input.Price)
		pkgassert.Equals(t, queue.PublishCalls, 1)
		pkgassert.Equals(t, queue.PublishCalledWithExchange, "orderPlaced")
		pkgassert.Each(t, queue.PublishCalledWithPayload, func(value byte, index int) {
			pkgassert.Equals(t, value, outputJSON[index])
		})
	})

	t.Run("With insufficient funds", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: wallet.AccountID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     120_000,
		}
		output, err := sut.Execute(input)
		pkgassert.Equals(t, output, nil)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Insufficient funds")
	})

	t.Run("With insufficient funds for a second order", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: wallet.AccountID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     100_000,
		}
		output, err := sut.Execute(input)
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, output, nil)
		output, err = sut.Execute(input)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, output, nil)
		pkgassert.Equals(t, err.Error(), "Insufficient funds")
	})
}
