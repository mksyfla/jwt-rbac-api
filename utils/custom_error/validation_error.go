package custom_error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationError(c *gin.Context, err error) {
	var errorMessage []string
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage = append(errorMessage, fmt.Sprintf("error on field %s, condition, %s", e.Field(), e.ActualTag()))
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": errorMessage,
	})
}
