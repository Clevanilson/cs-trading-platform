package usecase_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func TestGetWallet(t *testing.T) {
	var repository repository.WalletRepository
	var sut usecase.GetWallet

	setup := func() {
		repository = infrarepository.NewWalletMemoryRepository()
		deposit := usecase.NewDeposit(repository)
		if err := deposit.Execute(usecase.DepositInput{AccountID: "123", AssetID: "USD", Amount: 100}); err != nil {
			t.Fatal(err)
		}
		sut = usecase.NewGetWallet(repository)
	}

	t.Run("GetWallet", func(t *testing.T) {
		setup()
		input := usecase.GetWalletInput{AccountID: "123"}
		output, err := sut.Execute(input)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, len(output.Balances), 1)
		pkgassert.Equals(t, output.Balances[0].Amount, 100)
		pkgassert.Equals(t, output.Balances[0].AssetID, "USD")
	})	


}