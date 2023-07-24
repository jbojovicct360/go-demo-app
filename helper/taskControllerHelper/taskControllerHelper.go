package taskControllerHelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/configuration"
	"go-blog/dto"
	"go-blog/model"
)

func GetTaskById(ctx *gin.Context) model.Task {
	id := ctx.Query("id")
	var task model.Task
	configuration.DB.First(&task, id)
	return task
}

func TaskResponseHandler(ctx *gin.Context, task model.Task) {
	if task.ID == 0 {
		ctx.JSON(404, gin.H{
			"message": "Task not found!",
		})
		return
	}
	ctx.JSON(200, task)
}

func ExtractTaskDataFromRequest(ctx *gin.Context) dto.CreateTaskDTO {
	var taskCreationDTO dto.CreateTaskDTO
	if err := ctx.ShouldBind(&taskCreationDTO); err != nil {
		fmt.Println("Error binding task data!")
	}
	return taskCreationDTO
}
