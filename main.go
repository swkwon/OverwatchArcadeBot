package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	overwatchArcade   = "https://overwatcharcade.today/api/today"
	telegramURLFormat = "https://api.telegram.org/bot%s/SendMessage"
	tokenKey          = "OVERWATCH_ARCADE_BOT_API_TOKEN"
	telegramChannel   = "@overwatcharcade"
	post              = "POST"
	contentJSON       = "application/json"
)

// ArcadeItem ...
type ArcadeItem struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players string `json:"players"`
	Code    string `json:"code"`
}

// ArcadeInfo ...
type ArcadeInfo struct {
	UpdateAt      string     `json:"updated_at"`
	TileLarge     ArcadeItem `json:"tile_large"`
	TileWeekly1   ArcadeItem `json:"tile_weekly_1"`
	TileDaily     ArcadeItem `json:"tile_daily"`
	TileWeekly2   ArcadeItem `json:"tile_weekly_2"`
	TilePermanent ArcadeItem `json:"tile_permanent"`
}

func main() {
	res, err := http.Get(overwatchArcade)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	msg := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, telegramChannel, string(data))
	telegramURL := fmt.Sprintf(telegramURLFormat, os.Getenv(tokenKey))
	log.Printf("msg: %s\n", msg)
	log.Printf("telegram url: %s\n", telegramURL)

	reader := bytes.NewBuffer([]byte(msg))
	if req, err := http.NewRequest(post, telegramURL, reader); err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		req.Header.Set("Content-Type", contentJSON)
		client := &http.Client{}
		if res, err := client.Do(req); err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			defer res.Body.Close()
		}
	}
}
