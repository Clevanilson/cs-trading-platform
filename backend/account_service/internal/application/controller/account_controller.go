package controller

import (
	"encoding/json"
	"net/http"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/usecase"
	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
)

func CreateAccountController(
	httpServer pkgserver.HttpServer,
	createAccount usecase.CreateAccount,
	getAccount usecase.GetAccount,
) {
	httpServer.POST("/signup", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		var createAccountInput usecase.CreateAccountInput
		err := json.Unmarshal(input.Body, &createAccountInput)
		if err != nil {
			return nil, err
		}
		output, err := createAccount.Execute(createAccountInput)
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusCreated, Body: output}, nil
	})

	httpServer.GET("/get_account/:id", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		id := input.Params["id"]
		getAccountInput := usecase.GetAccountInput{ID: id}
		output, err := getAccount.Execute(getAccountInput)
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusOK, Body: output}, nil
	})
}
