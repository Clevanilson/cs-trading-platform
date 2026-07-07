package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
)

func WalletController(
	httpServer pkgserver.HttpServer,
	deposit usecase.Deposit,
	withdraw usecase.Withdraw,
	getWallet usecase.GetWallet,
) {
	httpServer.POST("/deposit", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		var depositInput usecase.DepositInput
		err := json.Unmarshal(input.Body, &depositInput)
		if err != nil {
			return nil, err
		}
		err = deposit.Execute(depositInput)
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusOK, Body: nil}, nil
	})

	httpServer.POST("/withdraw", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		var withdrawInput usecase.WithdrawInput
		err := json.Unmarshal(input.Body, &withdrawInput)
		if err != nil {
			return nil, err
		}
		err = withdraw.Execute(withdrawInput)
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusOK, Body: nil}, nil
	})

	httpServer.GET("/wallet/:accountID", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		accountID := input.Params["accountID"]
		fmt.Println("accountID", accountID)
		wallet, err := getWallet.Execute(usecase.GetWalletInput{AccountID: accountID})
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusOK, Body: wallet}, nil
	})
}
