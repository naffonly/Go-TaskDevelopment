package routes

import (
	"NotionTask/feature/taskDev/handler"
	"github.com/gin-gonic/gin"
)

func InitTaskDev(router *gin.RouterGroup, handlerInterface handler.TaskDevelopmentInterface) {
	router.GET("/get-data", handlerInterface.GetData())
}