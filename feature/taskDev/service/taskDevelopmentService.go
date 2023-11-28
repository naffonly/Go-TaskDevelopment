package service

import "github.com/dstotijn/go-notion"

type TaskDevelopmentServiceInterface interface {
	GetData() map[string]any
}

type notionIntegration struct {
	Client  *notion.Client
	Version string
}

func NewTaskDevelompentServiceImpl(client *notion.Client, ver string) TaskDevelopmentServiceInterface {
	return &notionIntegration{
		Client:  client,
		Version: ver,
	}
}

func (n *notionIntegration) GetData() map[string]any {
	//TODO implement me
	panic("implement me")
}
