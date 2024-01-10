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
	Draft(c *gin.Context)
	GetJobs(c *gin.Context)
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
		imgUrl, err := utils.Base64ToFile("jpg", img)
		if err != nil {
			custom_error.BadRequestError(c, err)
		}
		bodyRequest.Image[i] = imgUrl
	}

	result, err := h.jobService.Create(userId, bodyRequest)
	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	id := response.PostResponse{
		Id: result,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job created",
		"data":    id,
	})
}

func (h *jobHandler) Draft(c *gin.Context) {
	var userId string

	user, _ := c.Get("user")

	if user, ok := user.(model.User); ok {
		userId = user.Id
	}

	var bodyRequest request.DraftPostRequest

	if err := c.BindJSON(&bodyRequest); err != nil {
		custom_error.ValidationError(c, err)
		return
	}

	for i, img := range bodyRequest.Image {
		imgUrl, err := utils.Base64ToFile("jpg", img)
		if err != nil {
			custom_error.BadRequestError(c, err)
		}
		bodyRequest.Image[i] = imgUrl
	}

	result, err := h.jobService.Draft(userId, bodyRequest)
	if err != nil {
		custom_error.BadRequestError(c, err)
		return
	}

	id := response.PostResponse{
		Id: result,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "draft created",
		"data":    id,
	})
}

func (h *jobHandler) GetJobs(c *gin.Context) {
	jobs, _ := h.jobService.GetJobs()

	if len(jobs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": []model.JobUser{},
		})
		return
	}

	// jobMapped := []response.GetJobs{}
}
