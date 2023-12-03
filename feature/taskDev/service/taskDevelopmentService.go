package service

import (
	"NotionTask/util"
	"context"
	"encoding/json"
	"errors"
	"github.com/dstotijn/go-notion"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type TaskDevelopmentServiceInterface interface {
	GetData() (map[string]any, error)
	GetDatabasebNotion() (*notion.Database, error)
	GetDataDatabaseNotion() (*[]map[string]interface{}, error)
	GetDataPageNotion(id string) string
	GetDataScdule() error
	StopScedule() error
	CheckScedule() string
	GetDataBaseFilter() (*notion.DatabaseQueryResponse, error)
}

var (
	ctx           = context.Background()
	ticker        *time.Ticker
	stopSignal    chan struct{}
	tickerRunning bool
)

type notionIntegration struct {
	Client     *notion.Client
	Version    string
	DatabaseID string
	Url        string
}

func (n *notionIntegration) GetDataBaseFilter() (*notion.DatabaseQueryResponse, error) {

	oneMinuteAgo := util.GetTimeOneMinuteAgo()
	query := &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Timestamp: notion.TimestampLastEditedTime,
			DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
				LastEditedTime: &notion.DatePropertyFilter{
					OnOrAfter: oneMinuteAgo,
				},
			},
		},
	}

	result, err := n.Client.QueryDatabase(ctx, n.DatabaseID, query)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewTaskDevelompentServiceImpl(client *notion.Client, ver string, db string, url string) TaskDevelopmentServiceInterface {
	return &notionIntegration{
		Client:     client,
		Version:    ver,
		DatabaseID: db,
		Url:        url,
	}
}

func (n *notionIntegration) GetData() (map[string]any, error) {
	return nil, nil
}

func (n *notionIntegration) GetDatabasebNotion() (*notion.Database, error) {
	result, err := n.Client.FindDatabaseByID(ctx, n.DatabaseID)
	if err != nil {
		log.Fatal(err.Error())
		return &notion.Database{}, errors.New("failed get database")
	}
	return &result, nil
}

func (n *notionIntegration) GetDataDatabaseNotion() (*[]map[string]interface{}, error) {

	oneMinuteAgo := util.GetTimeOneMinuteAgo()
	query := &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Timestamp: notion.TimestampLastEditedTime,
			DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
				LastEditedTime: &notion.DatePropertyFilter{
					OnOrAfter: oneMinuteAgo,
				},
			},
		},
	}

	result, err := n.Client.QueryDatabase(ctx, n.DatabaseID, query)

	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("failed get database")
	}
	//extra object notion to json
	var dataObject []map[string]interface{}

	dataJSON, _ := json.Marshal(result.Results)

	errs := json.Unmarshal([]byte(dataJSON), &dataObject)
	if errs != nil {
		log.Fatal(errs.Error())
		return nil, err
	}

	//slice data properties
	propertiesList := make([]interface{}, len(dataObject))
	//slice data response
	var responses []map[string]interface{}

	for key, value := range dataObject {
		propertiesList[key] = value["properties"]
	}

	for _, properties := range propertiesList {
		// Cek tipe data dari properties, karena properties dalam hal ini adalah interface{}
		response := make(map[string]interface{})

		if prop, ok := properties.(map[string]interface{}); ok {
			for title, value := range prop {
				if data, isMap := value.(map[string]interface{}); isMap {
					for key, formula := range data {
						switch expr := key; expr {
						case "formula":
							if formulaMap, isMap := formula.(map[string]interface{}); isMap {
								for key, typeFormula := range formulaMap {
									if key == "number" {
										response[title] = typeFormula
									}
									if key == "string" {
										response[title] = typeFormula
									}
								}
							}
						case "date":
							if dateMap, isMap := formula.(map[string]interface{}); isMap {
								for key, Start := range dateMap {
									if key == "start" {
										response[title] = Start
									}
								}
							}
						case "checkbox":
							response[title] = formula
						case "created_time":
							rsDate, _ := util.TimeIsoToFormat(formula.(string))
							response[title] = rsDate
						case "created_by":
							if createByMap, isMap := formula.(map[string]interface{}); isMap {
								for key, name := range createByMap {
									if key == "name" {
										response[title] = name
									}
								}
							}
						case "select":
							if selectMap, isMap := formula.(map[string]interface{}); isMap {
								for key, name := range selectMap {
									if key == "name" {
										response[title] = name
									}
								}
							}
						case "email":
							response[title] = formula
						case "last_edited_by":
							if createByMap, isMap := formula.(map[string]interface{}); isMap {
								for key, name := range createByMap {
									if key == "name" {
										response[title] = name
									}
								}
							}
						case "last_edited_time":
							date, _ := util.TimeIsoToFormat(formula.(string))
							response[title] = date
						case "phone_number":
							response[title] = formula
						case "relation":
							if relations, isArray := formula.([]interface{}); isArray {
								var slice []string

								for _, relation := range relations {
									if relationMap, isMap := relation.(map[string]interface{}); isMap {
										if id, exists := relationMap["id"].(string); exists {
											rs := n.GetDataPageNotion(id)
											//rs := t.service.GetDataPageNotion(id)
											slice = append(slice, rs)
										}
									}
								}
								rs := strings.Join(slice, ", ")
								response[title] = rs
							}
						case "verification":
							if selectMap, isMap := formula.(map[string]interface{}); isMap {
								for key, verified := range selectMap {
									if key == "state" {
										if verified == "verified" || verified == "unverified" {
											response[title] = verified
										}
									}
								}
							}
						case "people":
							if relations, isArray := formula.([]interface{}); isArray {
								for _, relation := range relations {
									if relationMap, isMap := relation.(map[string]interface{}); isMap {
										if name, exists := relationMap["name"].(string); exists {
											response[title] = name
										}
									}
								}
							}
						case "unique_id":
							if uniqueMap, isMap := formula.(map[string]interface{}); isMap {
								for key, unique := range uniqueMap {
									if key == "number" {
										response[title] = unique
									}
								}
							}
						case "files":
							if files, isArray := formula.([]interface{}); isArray {
								for _, relation := range files {
									if relationMap, isMap := relation.(map[string]interface{}); isMap {
										if name, exists := relationMap["name"].(string); exists {
											response[title] = name
										}
									}
								}
							}
						case "multi_select":
							if multiSelect, isArray := formula.([]interface{}); isArray {
								var slice []string
								for _, selected := range multiSelect {
									if selectMap, isMap := selected.(map[string]interface{}); isMap {
										if name, exists := selectMap["name"].(string); exists {
											slice = append(slice, name)
										}
									}
								}
								rs := strings.Join(slice, ", ")
								response[title] = rs
							}
						case "url":
							response[title] = formula
						case "title":
							if titleMap, isArray := formula.([]interface{}); isArray {
								for _, titleNested := range titleMap {
									if relationMap, isMap := titleNested.(map[string]interface{}); isMap {
										if plainTtext, exists := relationMap["plain_text"].(string); exists {
											response[title] = plainTtext
										}
									}
								}
							}
						case "rich_text":
							if richTextMap, isArray := formula.([]interface{}); isArray {
								var slice []string
								for _, titleNested := range richTextMap {
									if relationMap, isMap := titleNested.(map[string]interface{}); isMap {
										if plainTtext, exists := relationMap["plain_text"].(string); exists {
											slice = append(slice, plainTtext)
										}
									}
								}
								rs := strings.Join(slice, "")
								response[title] = rs
							}
						}

					}
				}
			}
		}

		responses = append(responses, response)
	}

	return &responses, nil
}

func (n *notionIntegration) GetDataPageNotion(id string) string {
	result, err := n.Client.FindPageByID(ctx, id)
	if err != nil {
		log.Fatal(err.Error())
		return err.Error()
	}

	var dataObject map[string]interface{}

	dataJSON, _ := json.Marshal(result.Properties)

	errs := json.Unmarshal([]byte(dataJSON), &dataObject)
	if errs != nil {
		log.Fatal(errs.Error())
		return err.Error()
	}

	var namePage string

	for key, properties := range dataObject {
		if key == "Name" {
			if prop, ok := properties.(map[string]interface{}); ok {
				for keyMap, valueMap := range prop {
					if keyMap == "title" {
						if relations, isArray := valueMap.([]interface{}); isArray {
							var slice []string
							for _, relation := range relations {
								if relationMap, isMap := relation.(map[string]interface{}); isMap {
									if plainTtext, exists := relationMap["plain_text"].(string); exists {
										slice = append(slice, plainTtext)
									}
								}
							}
							rs := strings.Join(slice, ", ")
							namePage = rs
						}
					}
				}
			}
		}

	}
	return namePage
}

func (n *notionIntegration) CheckScedule() string {
	if tickerRunning {
		return "Scedule sedang berjalan"
	} else {
		return "Scedule tidak sedang berjalan"
	}
}

func (n *notionIntegration) StopScedule() error {
	if tickerRunning {
		ticker.Stop()
		stopSignal <- struct{}{}
		tickerRunning = false
		return nil
	}
	return errors.New("Scedule tidak sedang berjalan")
}

func (n *notionIntegration) GetDataScdule() error {
	client := &http.Client{}
	if !tickerRunning {
		ticker = time.NewTicker(1 * time.Minute)
		tickerRunning = true
		stopSignal = make(chan struct{})

		go func() {
			for {
				select {
				case <-ticker.C:
					log.Info("Scedule sedang berjalan : ", time.Now())
					rs, _ := n.GetDataDatabaseNotion()
					requestBody, err := json.Marshal(rs)
					if err != nil {
						log.Fatal("Error encoding JSON:", err)
						return
					}

					errs := util.SendToWebhook(*client, n.Url, requestBody)
					if errs != nil {
						log.Fatal("Error encoding JSON:", errs)
						return
					}

				case <-stopSignal:
					log.Info("Scedule telah dihentikan : ", time.Now())
					tickerRunning = false
					return
				}
			}
		}()
		return nil
	}
	log.Info("Scedule Sedang Berjalan")
	return errors.New("Scedule Sedang Berjalan")
}
