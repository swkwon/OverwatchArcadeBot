package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

func TestDiscord(t *testing.T) {
	webhookURL := "OVERWATCH_ARCADE_DISCORD_WEBHOOK"
	method := "POST"
	contentType := "application/json"
	body := `{"content":"hello world", "username":"OverwatchArcade"}`

	reader := bytes.NewBuffer([]byte(body))
	if req, err := http.NewRequest(method, os.Getenv(webhookURL), reader); err == nil {
		req.Header.Set("Content-Type", contentType)
		client := &http.Client{}
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()
		} else {
			t.Error(err)
		}
	}
}