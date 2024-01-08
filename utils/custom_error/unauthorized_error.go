package custom_error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnauthorizedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	c.Abort()
}
