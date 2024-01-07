package handler

import (
	"net/http"
	"sayembara/entity/request"
	"sayembara/entity/response"
	"sayembara/service"
	"sayembara/utils/custom_error"

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

func (h *userHandler) Create(c *gin.Context) {
	var bodyRequest request.UserRegisterRequest

	err := c.BindJSON(&bodyRequest)

	if err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	result, err := h.userService.Create(bodyRequest)

	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	id := response.UserRegisterResponse{
		Id: result,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    id,
	})
}
