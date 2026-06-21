package main

import (
	"fmt"
	"net/http"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	accountRepository := infrarepository.NewAccountMemoryRepository()

	server := gin.Default()
	server.POST("/signup", func(c *gin.Context) {
		var input usecase.CreateAccountInput
		if err := c.ShouldBindBodyWithJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Malformed json",
			})
		}
		useCase := usecase.NewCreateAccount(accountRepository)
		output, err := useCase.Execute(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
		}
		c.JSON(http.StatusCreated, output)
	})
	server.GET("/get_account/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
		}
		input := usecase.GetAccountInput{
			ID: id,
		}
		useCase := usecase.NewGetAccount(accountRepository)
		output, err := useCase.Execute(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
		}
		c.JSON(http.StatusOK, output)
	})
	server.Run(":3000")
	fmt.Println("Running My Test")
}
