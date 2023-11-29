package service

import (
	"context"
	_ "encoding/json"
	"errors"
	_ "fmt"
	"github.com/dstotijn/go-notion"
	log "github.com/sirupsen/logrus"
	_ "io/ioutil"
	_ "net/http"
	_ "time"
)

type TaskDevelopmentServiceInterface interface {
	GetData() (map[string]any, error)
	GetDatabasebNotion() (notion.Database, error)
	GetDataDatabaseNotion() (map[string]any, error)
}

var (
	ctx = context.Background()
)

type notionIntegration struct {
	Client     *notion.Client
	Version    string
	DatabaseID string
}

func NewTaskDevelompentServiceImpl(client *notion.Client, ver string, db string) TaskDevelopmentServiceInterface {
	return &notionIntegration{
		Client:     client,
		Version:    ver,
		DatabaseID: db,
	}
}

func (n *notionIntegration) GetData() (map[string]any, error) {
	return nil, nil
}

func (n *notionIntegration) GetDatabasebNotion() (notion.Database, error) {
	result, err := n.Client.FindDatabaseByID(ctx, n.DatabaseID)
	if err != nil {
		log.Fatal(err.Error())
		return notion.Database{}, errors.New("failed get database")
	}
	return result, nil
}

func (n *notionIntegration) GetDataDatabaseNotion() (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
