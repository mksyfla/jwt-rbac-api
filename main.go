package main

import (
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
	jobRepository := repository.NewJobRepository(id, db)

	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	jobService := service.NewJobService(jobRepository)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	jobHandler := handler.NewJobHandler(jobService)

	middleware := middleware.NewMiddleware(userRepository)

	router := gin.Default()

	router.Static("/public", "./public")

	api := router.Group("api")
	v1 := api.Group("v1")

	v1.POST("users", authHandler.Create)
	v1.POST("login", authHandler.Login)

	v1.GET("users", userHandler.GetUsers)

	v1.POST("jobs", middleware.AuthMiddleware, middleware.RoleBasedMiddleware("UMKM"), jobHandler.Create)

	router.Run()
}
