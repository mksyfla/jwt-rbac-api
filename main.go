package main

import (
	"sayembara/db"
	"sayembara/handler"
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

	router := gin.Default()

	api := router.Group("api")
	firstVersion := api.Group("v1")

	firstVersion.POST("users", userHandler.Create)

	router.Run()
}
