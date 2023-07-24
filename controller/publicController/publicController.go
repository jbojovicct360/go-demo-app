package publicController

import "github.com/gin-gonic/gin"

// HelloWorld test method that returns "Hello World"
func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
