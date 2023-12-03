package routes

import (
	"NotionTask/feature/taskDev/handler"
	"github.com/gin-gonic/gin"
)

func InitTaskDev(router *gin.RouterGroup, handlerInterface handler.TaskDevelopmentInterface) {
	router.GET("/get-data", handlerInterface.GetData())

	router.GET("/get-db", handlerInterface.GetDatabasePropertiesNotion())
	router.GET("/retrieve-db", handlerInterface.GetDataDatabaseNotion())
	router.GET("/page-notion/:id", handlerInterface.GetDataPageNotion())
	router.GET("/filter", handlerInterface.GetDataFilter())

	router.GET("/scedule", handlerInterface.GetScedule())
	router.GET("/stop", handlerInterface.StopScedule())
	router.GET("/check-scedulu", handlerInterface.CheckScedule())

}
