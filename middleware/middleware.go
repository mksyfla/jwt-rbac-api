package middleware

import (
	"fmt"
	"os"
	"sayembara/entity/model"
	"sayembara/repository"
	"sayembara/utils/custom_error"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type middleware struct {
	repository repository.UserRepository
}

func NewMiddleware(repository repository.UserRepository) *middleware {
	return &middleware{repository}
}

func (m *middleware) AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		custom_error.UnauthorizedError(c)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			custom_error.UnauthorizedError(c)
			return
		}

		sub, ok := claims["sub"].(string)

		if !ok {
			custom_error.UnauthorizedError(c)
			return
		}

		user, _ := m.repository.GetUserById(sub)

		if user.Id == "" {
			custom_error.UnauthorizedError(c)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		custom_error.UnauthorizedError(c)
		return
	}
}

func (m *middleware) RoleBasedMiddleware(roles string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := c.Get("user")

		if !err {
			custom_error.UnauthorizedError(c)
			return
		}

		category := user.(model.User).Category

		println(category)
		if category == roles {
			c.Next()
			return
		}

		custom_error.ForbiddenError(c)
	}
}
