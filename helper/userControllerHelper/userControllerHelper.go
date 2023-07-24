package userControllerHelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/configuration"
	"go-blog/dto"
	"go-blog/model"
)

func GetUserByUsername(username string) model.User {
	var user model.User
	configuration.DB.First(&user, "username = ?", username)
	return user
}

func ExtractUserDataFromRequest(ctx *gin.Context) dto.UserCreateDTO {
	var userCreationDto dto.UserCreateDTO
	if err := ctx.ShouldBind(&userCreationDto); err != nil {
		fmt.Println("Error binding data!")
	}
	return userCreationDto
}

func GetUserById(ctx *gin.Context) model.User {
	id := ctx.Query("id")
	return GetUserByStringID(id)
}

func GetUserByStringID(id string) model.User {
	var user model.User
	configuration.DB.First(&user, id)
	return user
}

func UserResponseHandler(user model.User, ctx *gin.Context) {
	if user.Id == 0 {
		ctx.JSON(404, gin.H{
			"message": "User not found!",
		})
		return
	}
	ctx.JSON(200, user)
}

func CreateUser(user model.User) bool {
	result := configuration.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("Error creating user!")
		return false
	}
	return true
}

func DeleteUserById(user model.User) {
	configuration.DB.Delete(&user)
}

func UpdateUser(user model.User) {
	configuration.DB.Save(&user)
}
