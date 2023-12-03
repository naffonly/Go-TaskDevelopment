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
	GetDataPageNotion() gin.HandlerFunc
	GetScedule() gin.HandlerFunc
	StopScedule() gin.HandlerFunc
	CheckScedule() gin.HandlerFunc
	GetDataFilter() gin.HandlerFunc
}

type taskDevelopmentHandler struct {
	service service.TaskDevelopmentServiceInterface
}

func (t *taskDevelopmentHandler) GetDataFilter() gin.HandlerFunc {
	return func(c *gin.Context) {

		rs, err := t.service.GetDataBaseFilter()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, rs)
	}
}

func (t *taskDevelopmentHandler) CheckScedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := t.service.CheckScedule()
		c.JSON(http.StatusOK, gin.H{
			"msg": rs,
		})
	}
}

func (t *taskDevelopmentHandler) StopScedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := t.service.StopScedule()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Jadwal fungsi telah dihentikan"})

	}
}

func (t *taskDevelopmentHandler) GetScedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := t.service.GetDataScdule()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Jadwal fungsi telah dimulai"})
	}
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

		result, err := t.service.GetDataDatabaseNotion()
		if err != nil {
			c.JSON(http.StatusBadRequest, model.FormatResponse(err.Error(), nil))
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (t *taskDevelopmentHandler) GetDataPageNotion() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		rs := t.service.GetDataPageNotion(id)

		c.JSON(http.StatusOK, rs)
	}
}
