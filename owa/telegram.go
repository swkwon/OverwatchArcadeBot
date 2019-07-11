package owa

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

const (
	tokenKey          = "OVERWATCH_ARCADE_BOT_API_TOKEN"
	telegramChannel   = "OVERWATCH_TG_CH"
	telegramURLFormat = "https://api.telegram.org/bot%s/SendMessage"
)

type Telegram struct {

}

func (t *Telegram) Send(text string) error {
	channel := os.Getenv(telegramChannel)
	token := os.Getenv(tokenKey)
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
