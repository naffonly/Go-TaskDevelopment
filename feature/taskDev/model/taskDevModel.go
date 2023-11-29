package model

type TaskNotionRequest struct {
}

type TaskNotionResponse struct {
}

type PropertyResponse struct {
	ActualHour string `json:"actualHour"`
	ArchivedAt string `json:"archivedAt"`
}

func FormatResponse(message string, data any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}
