package userController

import (
	"go-blog/helper/userControllerHelper"
	"go-blog/model"

	"github.com/gin-gonic/gin"
)

var GetUserByIdHelper = userControllerHelper.GetUserById
var GetUserByUsernameHelper = userControllerHelper.GetUserByUsername
var CreateUserHelper = userControllerHelper.CreateUser
var DeleteUserHelper = userControllerHelper.DeleteUserById
var UpdateUserHelper = userControllerHelper.UpdateUser

func DeleteUserByID(ctx *gin.Context) {
	user := GetUserByIdHelper(ctx)
	DeleteUserHelper(user)
	ctx.JSON(200, gin.H{
		"message": "User deleted successfully!",
	})
}

func UpdateUser(ctx *gin.Context) {
	userCreationDto := userControllerHelper.ExtractUserDataFromRequest(ctx)
	user := GetUserByIdHelper(ctx)
	user.Username = userCreationDto.Username
	UpdateUserHelper(user)
	ctx.JSON(200, user)
}

func GetUserByUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	user := GetUserByUsernameHelper(username)
	userControllerHelper.UserResponseHandler(user, ctx)
}

func GetUserByID(ctx *gin.Context) {
	user := GetUserByIdHelper(ctx)
	userControllerHelper.UserResponseHandler(user, ctx)
}

func CreateUser(ctx *gin.Context) {
	userCreationDto := userControllerHelper.ExtractUserDataFromRequest(ctx)
	user := model.User{Username: userCreationDto.Username}
	result := CreateUserHelper(user)
	if result {
		ctx.JSON(201, user)
	} else {
		ctx.JSON(400, gin.H{
			"message": "User cannot be created!",
		})
	}
}
