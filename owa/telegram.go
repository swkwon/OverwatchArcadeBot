package owa

import (
	"bytes"
	"fmt"
	"net/http"
)

const telegramURLFormat = "https://api.telegram.org/bot%s/SendMessage"

func Send(channel, token, text string) error {
	url := fmt.Sprintf(telegramURLFormat, token)
	msg := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, channel, text)

	reader := bytes.NewBuffer([]byte(msg))
	var err error
	if req, err := http.NewRequest("POST", url, reader); err == nil {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()
		}
	}
	return err
}
