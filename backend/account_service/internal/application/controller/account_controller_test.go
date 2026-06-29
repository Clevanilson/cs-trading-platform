package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/account_service/internal/infra/repository"
	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgclient "github.com/clevanilson/cs-trading-platform/devpack/pkg/client"
	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
)

func TestAccountController(t *testing.T) {
	httpServer := pkgserver.NewHttpAdapter()
	httpClient := pkgclient.NewHttpClient()
	accountRepository := infrarepository.NewAccountMemoryRepository()
	controller.CreateAccountController(
		httpServer,
		usecase.NewCreateAccount(accountRepository),
		usecase.NewGetAccount(accountRepository),
	)
	go httpServer.Start(3000)
	time.Sleep(300 * time.Millisecond)

	t.Run("Create account", func(t *testing.T) {
		res, err := httpClient.Post("http://localhost:3000/signup", map[string]string{"name": "John Doe"})
		if err != nil {
			t.Fatal(err)
		}
		pkgassert.Equals(t, res.StatusCode, http.StatusCreated)
		bodyOutput := usecase.CreateAccountOutput{}
		err = json.Unmarshal(res.Body, &bodyOutput)
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, bodyOutput.ID, "")
	})

	t.Run("Get account", func(t *testing.T) {
		account, _ := entity.NewAccount(entity.AccountBuilder{Name: "John Doe"})
		accountRepository.Save(account)
		res, err := httpClient.Get("http://localhost:3000/get_account/" + account.ID())
		sutOutput := usecase.GetAccountOutput{}
		json.Unmarshal(res.Body, &sutOutput)
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, res, nil)
		pkgassert.Equals(t, res.StatusCode, http.StatusOK)
		pkgassert.Equals(t, sutOutput.ID, account.ID())
		pkgassert.Equals(t, sutOutput.Name, account.Name())
	})

	httpServer.Stop()
}
