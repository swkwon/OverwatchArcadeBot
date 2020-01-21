package owa

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

const URL = "OVERWATCH_ARCADE_DISCORD_WEBHOOK"

type Discord struct {
}

// Send ...
func (d *Discord) Send(text string) error {
	body := `{"content":"%s", "username":"Overwatch today's arcade"}`
	sendData := []byte(fmt.Sprintf(body, text))
	log.Println(string(sendData))
	reader := bytes.NewBuffer(sendData)
	log.Println(os.Getenv(URL))
	if req, err := http.NewRequest("POST", os.Getenv(URL), reader); err == nil {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()
		}
	}
	return nil
}
