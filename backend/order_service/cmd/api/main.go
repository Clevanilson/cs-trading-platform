package main

import (
	"os"
	"strconv"

	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/controller"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"

	infrarepository "github.com/clevanilson/cs-trading-platform/order_service/internal/infra/repository"
)

func main() {
	queue, err := pkgqueue.NewRabbitAdapter()
	if err != nil {
		panic(err)
	}
	defer queue.Close()
	if err := queue.SetupQueue("orderPlaced.insertOrder"); err != nil {
		panic(err)
	}
	httpServer := pkgserver.NewHttpAdapter()
	walletRepository := infrarepository.NewWalletMemoryRepository()
	controller.WalletController(
		httpServer,
		usecase.NewDeposit(walletRepository),
		usecase.NewWithdraw(walletRepository),
		usecase.NewGetWallet(walletRepository),
	)
	orderRepository := infrarepository.NewOrderMemoryRepository()
	controller.OrderController(
		httpServer,
		usecase.NewPlaceOrder(walletRepository, orderRepository, queue),
	)
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
