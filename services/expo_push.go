package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendExpoPush(token string, title string, body string) error {
	expoEndpoint := "https://exp.host/--/api/v2/push/send"

	payload := map[string]interface{}{
		"to":    token,
		"title": title,
		"body":  body,
	}

	data, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", expoEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
