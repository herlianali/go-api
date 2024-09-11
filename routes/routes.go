package routes

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/repositories"
	"go-api/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	route := gin.Default()

	config.InitDB()

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	route.GET("/users", userController.GetAllUsers)
	route.GET("/users/:id", userController.GetUserByID)
	route.POST("/users", userController.CreateUser)
	route.PUT("/users/:id", userController.UpdateUser)
	route.DELETE("/users/:id", userController.DeleteUser)

	return route
}
