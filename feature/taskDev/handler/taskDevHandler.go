package handler

import (
	"NotionTask/feature/taskDev/model"
	"NotionTask/feature/taskDev/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"net/http"
)

//Notion Task BE
//1. Get data dari notion ubah menjadi key value
//2. Pastikan property fleksibel dari notion
//3. Buat penjadwalan untuk pengambilan data per 1 menit
//4. Push yang update at nya kurant dari 1 menit ke webhook.transtrack.id

type TaskDevelopmentInterface interface {
	GetData() gin.HandlerFunc
	GetDatabasePropertiesNotion() gin.HandlerFunc
	GetDataDatabaseNotion() gin.HandlerFunc
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
	return func(c *gin.Context) {
		panic("implement me")
	}
}
func (t *taskDevelopmentHandler) GetDatabasePropertiesNotion() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := t.service.GetDatabasebNotion()
		if err != nil {
			c.JSON(http.StatusBadRequest, model.FormatResponse(err.Error(), nil))
			c.Abort()
			return
		}

		dataJSON, err := json.Marshal(result.Properties)
		if err != nil {
			log.Info(err.Error())
			c.Abort()
			return
		}

		var dataResult map[string]interface{}
		if err := json.Unmarshal([]byte(dataJSON), &dataResult); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, model.FormatResponse("Success Found Property Database Notion", dataResult))

		//
		//dataResponse := make(map[string]string)
		//for key, value := range dataResult {
		//	if prop, ok := value.(map[string]interface{}); ok {
		//		if name, exists := prop["name"]; exists {
		//			if nameStr, ok := name.(string); ok {
		//				dataResponse[key] = nameStr
		//			}
		//		}
		//	}
		//}

		//c.JSON(http.StatusOK, model.FormatResponse("Success Found Database Notion", dataResponse))
	}
}

func (t *taskDevelopmentHandler) GetDataDatabaseNotion() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic("implement me")
	}
}
