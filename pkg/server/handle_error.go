package server

import (
	"errors"
	"net/http"

	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var customError errorc.ErrorC
	if errors.As(err, &customError) {
		errorResponse := gin.H{
			"message": err.Error(),
		}
		if customError.Code() == errorc.DomainErrorCode {
			c.JSON(http.StatusBadRequest, errorResponse)
		} else if customError.Code() == errorc.NotFoundErrorCode {
			c.JSON(http.StatusNotFound, errorResponse)
		}
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Unepected error",
	})
}
