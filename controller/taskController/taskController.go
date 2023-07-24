package taskController

import (
	"github.com/gin-gonic/gin"
	"go-blog/configuration"
	"go-blog/helper/taskControllerHelper"
	"go-blog/helper/userControllerHelper"
	"go-blog/model"
	"strconv"
)

func CreateTask(ctx *gin.Context) {
	taskCreationDTO := taskControllerHelper.ExtractTaskDataFromRequest(ctx)
	var task model.Task
	task.Title = taskCreationDTO.Title
	task.Description = taskCreationDTO.Description
	task.UserID = taskCreationDTO.UserID
	task.User = userControllerHelper.GetUserByStringID(strconv.Itoa(int(taskCreationDTO.UserID)))
	result := configuration.DB.Create(&task)
	if result.Error != nil {
		ctx.JSON(500, "Error creating task!")
		return
	}
	ctx.JSON(200, task)
}

func UpdateTask(ctx *gin.Context) {
	taskCreationDTO := taskControllerHelper.ExtractTaskDataFromRequest(ctx)
	task := taskControllerHelper.GetTaskById(ctx)
	task.Title = taskCreationDTO.Title
	task.Description = taskCreationDTO.Description
	result := configuration.DB.Save(&task)
	if result.Error != nil {
		ctx.JSON(500, "Error updating task!")
		return
	}
	ctx.JSON(200, task)
}

func DeleteTask(ctx *gin.Context) {
	task := taskControllerHelper.GetTaskById(ctx)
	configuration.DB.Delete(&task)
	ctx.JSON(200, gin.H{
		"message": "Task deleted successfully!",
	})
}

func GetTaskById(ctx *gin.Context) {
	task := taskControllerHelper.GetTaskById(ctx)
	task.User = userControllerHelper.GetUserByStringID(strconv.Itoa(int(task.UserID)))
	taskControllerHelper.TaskResponseHandler(ctx, task)
}

func GetTaskByUserId(ctx *gin.Context) {
	userID := ctx.Query("userID")
	var tasks []model.Task
	configuration.DB.Where("user_id = ?", userID).Find(&tasks)
	if len(tasks) == 0 {
		ctx.JSON(404, "Tasks not found!")
		return
	}
	ctx.JSON(200, tasks)
}
