package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/clevanilson/cs-trading-platform/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/pkg/client"
	"github.com/clevanilson/cs-trading-platform/pkg/server"
)

func TestAccountController(t *testing.T) {
	var httpServer server.HttpServer
	var httpClient client.HttpClient

	setup := func() {
		httpServer = server.NewHttpAdapter()
		httpClient = client.NewHttpClient()
		accountRepository := infrarepository.NewAccountMemoryRepository()
		createAccount := usecase.NewCreateAccount(accountRepository)
		getAccount := usecase.NewGetAccount(accountRepository)
		controller.CreateAccountController(httpServer, createAccount, getAccount)
		go httpServer.Start(3000)
		time.Sleep(100 * time.Millisecond)
	}

	tearDown := func() {
		httpServer.Stop()
	}

	t.Run("Create account", func(t *testing.T) {
		setup()
		res, err := httpClient.Post("http://localhost:3000/signup", map[string]string{"name": "John Doe"})
		if err != nil {
			t.Fatal(err)
		}
		assert.Equals(t, res.StatusCode, http.StatusCreated)
		bodyOutput := usecase.CreateAccountOutput{}
		err = json.Unmarshal(res.Body, &bodyOutput)
		assert.Equals(t, err, nil)
		assert.NotEquals(t, bodyOutput.ID, "")
		tearDown()
	})

	t.Run("Get account", func(t *testing.T) {
		setup()
		createAccountInput := usecase.CreateAccountInput{Name: "John Doe"}
		res, err := httpClient.Post("http://localhost:3000/signup", createAccountInput)
		bodyOutput := usecase.GetAccountOutput{}
		err = json.Unmarshal(res.Body, &bodyOutput)
		res, err = httpClient.Get("http://localhost:3000/get_account/" + bodyOutput.ID)
		assert.Equals(t, err, nil)
		assert.NotEquals(t, res, nil)
		assert.Equals(t, res.StatusCode, http.StatusOK)
		assert.NotEquals(t, bodyOutput.ID, "")
		tearDown()
	})
}
