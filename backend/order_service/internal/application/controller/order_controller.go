package controller

import (
	"encoding/json"
	"net/http"

	pkgserver "github.com/clevanilson/cs-trading-platform/devpack/pkg/server"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/application/usecase"
)

func OrderController(
	httpServer pkgserver.HttpServer,
	placeOrder usecase.PlaceOrder,
) {
	httpServer.POST("/place_order", func(input *pkgserver.HandlerInput) (*pkgserver.Response, error) {
		var placeOrderInput usecase.PlaceOrderInput
		err := json.Unmarshal(input.Body, &placeOrderInput)
		if err != nil {
			return nil, err
		}
		output, err := placeOrder.Execute(placeOrderInput)
		if err != nil {
			return nil, err
		}
		return &pkgserver.Response{StatusCode: http.StatusOK, Body: output}, nil
	})
}
