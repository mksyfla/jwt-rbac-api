package handler

import (
	"net/http"
	"sayembara/entity/request"
	"sayembara/entity/response"
	"sayembara/service"
	"sayembara/utils/custom_error"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) Create(c *gin.Context) {
	var bodyRequest request.UserRegisterRequest

	if err := c.BindJSON(&bodyRequest); err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	result, err := h.authService.Create(bodyRequest)

	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	id := response.PostResponse{
		Id: result,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    id,
	})
}

func (h *authHandler) Login(c *gin.Context) {
	var bodyRequest request.UserLoginRequest

	if err := c.BindJSON(&bodyRequest); err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	token, err := h.authService.Login(bodyRequest)

	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*2, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}
