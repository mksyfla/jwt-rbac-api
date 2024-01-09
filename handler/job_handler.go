package handler

import (
	"net/http"
	"sayembara/entity/model"
	"sayembara/entity/request"
	"sayembara/entity/response"
	"sayembara/service"
	"sayembara/utils"
	"sayembara/utils/custom_error"

	"github.com/gin-gonic/gin"
)

type JobHandler interface {
	Create(c *gin.Context)
}

type jobHandler struct {
	jobService service.JobService
}

func NewJobHandler(jobService service.JobService) *jobHandler {
	return &jobHandler{jobService}
}

func (h *jobHandler) Create(c *gin.Context) {
	var userId string

	user, _ := c.Get("user")

	if user, ok := user.(model.User); ok {
		userId = user.Id
	}

	var bodyRequest request.JobPostRequest

	if err := c.BindJSON(&bodyRequest); err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	for i, img := range bodyRequest.Image {
		imgUrl, err := utils.Base64ToJpg(img)
		if err != nil {
			custom_error.BadRequestError(c, err)
		}
		bodyRequest.Image[i] = imgUrl
	}

	result, err := h.jobService.Create(userId, bodyRequest)
	if err != nil {
		custom_error.BadRequestError(c, err)
	}

	id := response.JobPostResponse{
		Id: result,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job created",
		"data":    id,
	})
}
