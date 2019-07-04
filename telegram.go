package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const telegramURLFormat = "https://api.telegram.org/bot%s/SendMessage"

func send(channel, token, text string) error {
	url := fmt.Sprintf(telegramURLFormat, token)
	msg := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, channel, text)

	reader := bytes.NewBuffer([]byte(msg))
	if req, err := http.NewRequest("POST", url, reader); err != nil {
		log.Fatal(err)
	} else {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if res, err := client.Do(req); err != nil {
			log.Fatal(err)
		} else {
			defer res.Body.Close()
		}
	}
	return nil
}
