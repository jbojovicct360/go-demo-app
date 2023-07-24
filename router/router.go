package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller/publicController"
	"go-blog/controller/taskController"
	"go-blog/controller/userController"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/by-id", userController.GetUserByID)
		userRoutes.GET("/by-username", userController.GetUserByUsername)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/", userController.UpdateUser)
		userRoutes.DELETE("/", userController.DeleteUserByID)
	}
}

func TaskRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/task")
	{
		taskRoutes.GET("/by-id", taskController.GetTaskById)
		taskRoutes.GET("/by-userID", taskController.GetTaskByUserId)
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.PUT("/", taskController.UpdateTask)
		taskRoutes.DELETE("/", taskController.DeleteTask)
	}
}

// PublicRoutes creating method for public routes
func PublicRoutes(router *gin.Engine) {
	publicRoutes := router.Group("/")
	{
		publicRoutes.GET("/hello-world", publicController.HelloWorld)
	}
}
