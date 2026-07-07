package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgclient "github.com/clevanilson/cs-trading-platform/devpack/pkg/client"
	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func TestWalletController(t *testing.T) {
	var httpServer pkgserver.HttpServer
	var httpClient pkgclient.HttpClient
	var walletRepository repository.WalletRepository

	setup := func() {
		httpClient = pkgclient.NewHttpClient()
		httpServer = pkgserver.NewHttpAdapter()
		walletRepository = infrarepository.NewWalletMemoryRepository()
		wallet, _ := entity.NewWallet(entity.WalletBuilder{AccountID: "123"})
		walletRepository.Update(wallet)
		wallet.Deposit("USD", 100)
		walletRepository.Update(wallet)
		deposit := usecase.NewDeposit(walletRepository)
		withdraw := usecase.NewWithdraw(walletRepository)
		getWallet:= usecase.NewGetWallet(walletRepository)
		controller.WalletController(httpServer, deposit, withdraw, getWallet)
		go httpServer.Start(3102)
		time.Sleep(1 * time.Second)
	}

	t.Run("Deposit", func(t *testing.T) {
		setup()
		body := usecase.DepositInput{
			AccountID: "123",
			AssetID:   "USD",
			Amount:    100,
		}
		response, err := httpClient.Post("http://localhost:3102/deposit", body)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, response.StatusCode, http.StatusOK)
		httpServer.Stop()
	})

	t.Run("Withdraw", func(t *testing.T) {
		setup()
		body := usecase.WithdrawInput{
			AccountID: "123",
			AssetID:   "USD",
			Amount:    100,
		}
		response, err := httpClient.Post("http://localhost:3102/withdraw", body)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, response.StatusCode, http.StatusOK)
		httpServer.Stop()
	})

	t.Run("Get wallet", func(t *testing.T) {
		setup()
		response, err := httpClient.Get("http://localhost:3102/wallet/123")
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, response.StatusCode, http.StatusOK)
		var body usecase.GetWalletOutput
		err = json.Unmarshal(response.Body, &body)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, len(body.Balances), 1)
		pkgassert.Equals(t, body.Balances[0].AssetID, "USD")
		pkgassert.Equals(t, body.Balances[0].Amount, 100)
		httpServer.Stop()
	})

}