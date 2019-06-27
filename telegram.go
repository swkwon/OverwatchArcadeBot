package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const telegramURLFormat = "https://api.telegram.org/bot%s/SendMessage"

func send(channel, token, text string) error {
	log.Printf("channel: %s, token: %s, text: %s\n", channel, token, text)
	url := fmt.Sprintf(telegramURLFormat, token)
	msg := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, channel, text)

	reader := bytes.NewBuffer([]byte(msg))
	if req, err := http.NewRequest("POST", url, reader); err != nil {
		log.Fatal(err)
		return err
	} else {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if res, err := client.Do(req); err != nil {
			return err
		} else {
			defer res.Body.Close()
		}
	}
	return nil
}
