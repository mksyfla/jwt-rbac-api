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

	if err := c.BindJSON(&bodyRequest); err != nil {
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

func (h *userHandler) Login(c *gin.Context) {
	var bodyRequest request.UserLoginRequest

	if err := c.BindJSON(&bodyRequest); err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	token, err := h.userService.Login(bodyRequest)

	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, _ := h.userService.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
