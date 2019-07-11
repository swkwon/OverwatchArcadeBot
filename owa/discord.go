package owa

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

const URL = "OVERWATCH_ARCADE_DISCORD_WEBHOOK"

type Discord struct {

}

func (d *Discord) Send(text string) error {
	body := `{"content":"%s", "username":"OverwatchArcade", "avatar_url":"https://cdn.discordapp.com/icons/592965021051650058/7c85a56a8741273862df36d3c72973d6.png"}`

	reader := bytes.NewBuffer([]byte(fmt.Sprintf(body, text)))
	if req, err := http.NewRequest("POST", os.Getenv(URL), reader); err == nil {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()
		}
	}
	return nil
}