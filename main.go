package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-blog/configuration"
	"go-blog/router"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file!")
	}
}

func main() {
	// creating gin app
	ginApp := gin.Default()
	// connecting app to database and getting database object
	configuration.CreateDatabase()
	// creating migrations
	configuration.CreateMigrations()
	// getting app port from env file
	AppPort := os.Getenv("PORT")
	// adding routes to app
	router.PublicRoutes(ginApp)
	router.UserRoutes(ginApp)
	router.TaskRoutes(ginApp)
	// starting app
	err := ginApp.Run(":" + AppPort)
	if err != nil {
		fmt.Println("Error running app!")
		return
	}
}
