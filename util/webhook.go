package util

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendToWebhook(client http.Client, url string, data []byte) error {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
