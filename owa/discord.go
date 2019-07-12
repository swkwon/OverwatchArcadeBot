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
	body := `{"content":"%s", "username":"Overwatch today's arcade"}`

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