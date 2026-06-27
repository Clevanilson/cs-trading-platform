package main

import (
	"github.com/clevanilson/cs-trading-platform/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/server"
)

func main() {
	httpServer := server.NewHttpAdapter()
	accountRepository := infrarepository.NewAccountMemoryRepository()
	createAccount := usecase.NewCreateAccount(accountRepository)
	getAccount := usecase.NewGetAccount(accountRepository)
	controller.CreateAccountController(httpServer, createAccount, getAccount)
	httpServer.Start(3000)
}
