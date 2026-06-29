package main

import (
	"fmt"

	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
)

func main() {
	queue, err := pkgqueue.NewRabbitAdapter()
	if err != nil {
		panic(err)
	}
	defer queue.Close()
	if err := queue.SetupExchange("orderPlaced"); err != nil {
		panic(err)
	}
	if err := queue.SetupQueue("orderPlaced.insertOrder"); err != nil {
		panic(err)
	}
	if err := queue.BindQueue("orderPlaced.insertOrder"); err != nil {
		panic(err)
	}
	if err := queue.Publish("orderPlaced", []byte(`{ "message": "ttestt"}`)); err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}
