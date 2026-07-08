package main

import (
	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
	"github.com/clevanilson/cs-trading-platform/matching_service/internal/application/controller"
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
	if err := queue.SetupQueue("orderFilled.updateOrder"); err != nil {
		panic(err)
	}
	if err := controller.CreateBookController(queue); err != nil {
		panic(err)
	}

}
