package main

import (
	"os"
	"strconv"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/account_service/internal/infra/repository"
	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
)

func main() {
	httpServer := pkgserver.NewHttpAdapter()
	accountRepository := infrarepository.NewAccountMemoryRepository()
	createAccount := usecase.NewCreateAccount(accountRepository)
	getAccount := usecase.NewGetAccount(accountRepository)
	controller.CreateAccountController(httpServer, createAccount, getAccount)
	PORT_ENV := os.Getenv("PORT")
	port := 3000
	if PORT_ENV != "" {
		converterd, err := strconv.ParseInt(PORT_ENV, 10, 64)
		if err != nil {
			panic(err)
		}
		port = int(converterd)
	}
	httpServer.Start(port)
}
