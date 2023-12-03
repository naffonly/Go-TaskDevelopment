package model

import "encoding/json"

type TaskNotionRequest struct {
}

type TaskNotionModel struct {
}

type ModelProperty struct {
	Properties []interface{} `json:"properties"`
}

func FormatResponse(message string, data any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}

func PropertiesToResponse(data []interface{}) (string, error) {
	// Mengonversi data menjadi representasi JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Mengembalikan data dalam bentuk string
	return string(jsonData), nil
}
