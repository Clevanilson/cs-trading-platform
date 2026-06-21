package main

import (
	"fmt"
	"net/http"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
	"github.com/clevanilson/cs-trading-platform/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	accountRepository := infrarepository.NewAccountMemoryRepository()

	httpServer := gin.Default()
	httpServer.POST("/signup", func(c *gin.Context) {
		var input usecase.CreateAccountInput
		if err := c.ShouldBindBodyWithJSON(&input); err != nil {
			server.HandleError(c, err)
			return
		}
		useCase := usecase.NewCreateAccount(accountRepository)
		output, err := useCase.Execute(input)
		if err != nil {
			server.HandleError(c, err)
			return
		}
		c.JSON(http.StatusCreated, output)
	})
	httpServer.GET("/get_account/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			server.HandleError(c, errorc.NewDomain("id"))
			return
		}
		input := usecase.GetAccountInput{
			ID: id,
		}
		useCase := usecase.NewGetAccount(accountRepository)
		output, err := useCase.Execute(input)
		if err != nil {
			server.HandleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	})
	httpServer.Run(":3000")
	fmt.Println("Running My Test")
}
