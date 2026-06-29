package usecase_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func TestDeposit(t *testing.T) {
	var repository repository.WalletRepository
	var sut usecase.Deposit
	var wallet entity.Wallet

	setup := func() {
		var err error
		repository = infrarepository.NewWalletMemoryRepository()
		sut = usecase.NewDeposit(repository)
		wallet, err = entity.NewWallet(entity.WalletBuilder{AccountID: "abvoafd-8765-4321-0987-654321098765"})
		err = repository.Update(wallet)
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With valid value", func(t *testing.T) {
		setup()
		err := sut.Execute(usecase.DepositInput{
			AccountID: wallet.AccountID(),
			AssetID:   "USD",
			Amount:    1000,
		})
		wallet, err = repository.GetByAccountID(wallet.AccountID())
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, len(wallet.Balances()), 1)
		pkgassert.Equals(t, wallet.Balances()[0].AssetID(), "USD")
		pkgassert.Equals(t, wallet.Balances()[0].Amount(), 1000)
	})

	t.Run("With invalid value", func(t *testing.T) {
		err := sut.Execute(usecase.DepositInput{
			AccountID: wallet.AccountID(),
			AssetID:   "USD",
			Amount:    -1000,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
	})
}
