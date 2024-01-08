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

	// router.POST("/upload", func(c *gin.Context) {
	// 	file, err := c.FormFile("image")
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	uploadDir := "public"
	// 	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	filename := filepath.Join(uploadDir, file.Filename)
	// 	fmt.Println(filename)
	// 	if err := c.SaveUploadedFile(file, filename); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
	// })

	// router.GET("/upload", func(c *gin.Context) {
	// 	c.File(filepath.Join("uploads/1.png"))
	// })

	api := router.Group("api")
	firstVersion := api.Group("v1")

	// user auth
	firstVersion.POST("users", userHandler.Create)
	firstVersion.POST("login", userHandler.Login)

	// users public
	firstVersion.GET("users", userHandler.GetUsers)

	router.Run()
}
