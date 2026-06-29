package usecase_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func TestPlaceOrder(t *testing.T) {
	var wallet entity.Wallet
	var walletRepository repository.WalletRepository
	var orderRepository repository.OrderRepository
	var sut usecase.PlaceOrder

	setup := func() {
		var err error
		wallet, err = entity.NewWallet(entity.WalletBuilder{AccountID: "Lune"})
		err = wallet.Deposit("USD", 100_000)
		walletRepository = infrarepository.NewWalletMemoryRepository()
		err = walletRepository.Update(wallet)
		orderRepository = infrarepository.NewOrderMemoryRepository()
		sut = usecase.NewPlaceOrder(walletRepository, orderRepository)
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
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, output, nil)
		savedOrder, err := orderRepository.GetByID(output.OrderID)
		pkgassert.NotEquals(t, savedOrder, nil)
		pkgassert.Equals(t, savedOrder.ID(), output.OrderID)
		pkgassert.Equals(t, savedOrder.MarketID(), input.MarketID)
		pkgassert.Equals(t, savedOrder.Side(), input.Side)
		pkgassert.Equals(t, savedOrder.Amount(), input.Amount)
		pkgassert.Equals(t, savedOrder.Price(), input.Price)
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
