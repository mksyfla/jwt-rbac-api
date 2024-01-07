package custom_error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}
