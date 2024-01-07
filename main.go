package main

import (
	"net/http"
	"sayembara/db"
	"sayembara/handler"
	"sayembara/middleware"
	"sayembara/repository"
	"sayembara/service"
	"sayembara/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db := db.NewDB()
	id := utils.NewIdGenerator

	userRepository := repository.NewUserRepository(id, db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	middleware := middleware.NewMiddleware(userRepository)

	router := gin.Default()

	api := router.Group("api")
	firstVersion := api.Group("v1")

	firstVersion.POST("users", userHandler.Create)
	firstVersion.POST("login", userHandler.Login)
	firstVersion.GET("auth", middleware.AuthMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test auth",
		})
	})

	router.Run()
}
