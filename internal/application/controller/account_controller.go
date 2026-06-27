package controller

import (
	"encoding/json"
	"net/http"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/pkg/server"
)

func CreateAccountController(
	httpServer server.HttpServer,
	createAccount usecase.CreateAccount,
	getAccount usecase.GetAccount,
) {
	httpServer.POST("/signup", func(input *server.HandlerInput) (*server.Response, error) {
		var createAccountInput usecase.CreateAccountInput
		err := json.Unmarshal(input.Body, &createAccountInput)
		if err != nil {
			return nil, err
		}
		output, err := createAccount.Execute(createAccountInput)
		if err != nil {
			return nil, err
		}
		return &server.Response{StatusCode: http.StatusCreated, Body: output}, nil
	})
	httpServer.GET("/get_account/:id", func(input *server.HandlerInput) (*server.Response, error) {
		id := input.Params["id"]
		getAccountInput := usecase.GetAccountInput{ID: id}
		output, err := getAccount.Execute(getAccountInput)
		if err != nil {
			return nil, err
		}
		return &server.Response{StatusCode: http.StatusOK, Body: output}, nil
	})
}
