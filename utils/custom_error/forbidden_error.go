package custom_error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForbiddenError(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "Forbidden",
	})
	c.Abort()
}
