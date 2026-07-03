package controller

import (
	"encoding/json"
	"fmt"

	pkgdto "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/dto"
	pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"
	pkgqueue "github.com/clevanilson/cs-trading-platform/devpack/pkg/queue"
	"github.com/clevanilson/cs-trading-platform/matching_service/internal/domain/entity"
)

func CreateBookController(queue pkgqueue.Queue) error {
	books := make(map[string]entity.Book)
	return queue.Consume("orderPlaced.insertOrder", func(data []byte) error {
		orderEvent := pkgdto.PlaceOrderEvent{}
		err := json.Unmarshal(data, &orderEvent)
		if err != nil {
			return err
		}
		book := books[orderEvent.MarketID]
		if book == nil {
			book = entity.NewBook(orderEvent.MarketID)
			books[orderEvent.MarketID] = book
		}
		order, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			ID:        &orderEvent.OrderID,
			AccountID: orderEvent.AccountID,
			MarketID:  orderEvent.MarketID,
			Side:      orderEvent.Side,
			Amount:    orderEvent.Amount,
			Price:     orderEvent.Price,
		})
		if err != nil {
			return err
		}
		book.Insert(order)
		fmt.Println(string(data))
		return nil
	})
}
