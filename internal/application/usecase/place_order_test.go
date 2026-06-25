package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestPlaceOrder(t *testing.T) {
	var account entity.Account
	var accountRepository repository.AccountRepository
	var orderRepository repository.OrderRepository
	var sut usecase.PlaceOrder

	setup := func() {
		var err error
		account, err = entity.NewAccount(entity.AccountBuilder{Name: "Lune"})
		err = account.Deposit("USD", 100_000)
		accountRepository = infrarepository.NewAccountMemoryRepository()
		err = accountRepository.Save(account)
		orderRepository = infrarepository.NewOrderMemoryRepository()
		sut = usecase.NewPlaceOrder(accountRepository, orderRepository)
		assert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: account.ID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     78_000,
		}
		output, err := sut.Execute(input)
		assert.Equals(t, err, nil)
		assert.NotEquals(t, output, nil)
		savedOrder, err := orderRepository.GetById(output.OrderID)
		assert.NotEquals(t, savedOrder, nil)
		assert.Equals(t, savedOrder.ID(), output.OrderID)
		assert.Equals(t, savedOrder.MarketID(), input.MarketID)
		assert.Equals(t, savedOrder.Side(), input.Side)
		assert.Equals(t, savedOrder.Amount(), input.Amount)
		assert.Equals(t, savedOrder.Price(), input.Price)
	})

	t.Run("With a non-existent account", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: "cbce8b3e-c5fb-4118-87dc-db0897241c48",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     78_000,
		}
		output, err := sut.Execute(input)
		assert.Equals(t, output, nil)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Account not found")
	})

	t.Run("With insufficient funds", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: account.ID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     120_000,
		}
		output, err := sut.Execute(input)
		assert.Equals(t, output, nil)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Insufficient funds")
	})

	t.Run("With insufficient funds for a second order", func(t *testing.T) {
		setup()
		input := usecase.PlaceOrderInput{
			AccountID: account.ID(),
			MarketID:  "BTC-USD",
			Side:      "buy",
			Amount:    1,
			Price:     100_000,
		}
		output, err := sut.Execute(input)
		assert.Equals(t, err, nil)
		assert.NotEquals(t, output, nil)
		output, err = sut.Execute(input)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, output, nil)
		assert.Equals(t, err.Error(), "Insufficient funds")
	})
}
