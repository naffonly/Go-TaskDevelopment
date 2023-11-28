package handler

import (
	"NotionTask/feature/taskDev/service"
	"github.com/gin-gonic/gin"
)

//Notion Task BE
//1. Get data dari notion ubah menjadi key value
//2. Pastikan property fleksibel dari notion
//3. Buat penjadwalan untuk pengambilan data per 1 menit
//4. Push yang update at nya kurant dari 1 menit ke webhook.transtrack.id

type TaskDevelopmentInterface interface {
	GetData() gin.HandlerFunc
}

type taskDevelopmentHandler struct {
	service service.TaskDevelopmentServiceInterface
}

func NewTaskDevelopmentHandler(service service.TaskDevelopmentServiceInterface) TaskDevelopmentInterface {
	return &taskDevelopmentHandler{
		service: service,
	}
}

func (t *taskDevelopmentHandler) GetData() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
