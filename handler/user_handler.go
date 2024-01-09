package handler

import (
	"net/http"
	"sayembara/entity/model"
	"sayembara/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	PostUserHandler(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, _ := h.userService.GetUsers()

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": []model.User{},
		})
		return
	}

	baseURL := "http://127.0.0.1:8080"
	for i := range users {
		users[i].Profile = baseURL + "/" + users[i].Profile
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
