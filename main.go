package main

import (
	configHandler "NotionTask/config"
	"NotionTask/feature/taskDev/handler"
	serviceHandler "NotionTask/feature/taskDev/service"
	"NotionTask/routes"
	"NotionTask/util"
	_ "github.com/dstotijn/go-notion"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupAppRouter().Run(":8080")
}

func SetupAppRouter() *gin.Engine {
	router := gin.Default()

	config := configHandler.InitConfig()
	notionConfig := util.InitNotionApi(config.KeyNotion)
	service := serviceHandler.NewTaskDevelompentServiceImpl(notionConfig, config.Version, config.DB_ID, config.WebHook_URL)
	handlerTaskDev := handler.NewTaskDevelopmentHandler(service)

	public := router.Group("/api")
	routes.InitTaskDev(public, handlerTaskDev)
	return router
}
