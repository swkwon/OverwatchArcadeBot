package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

const (
	overwatchArcade = "https://overwatcharcade.today/api/today"
	telegramMsg     = `test\ntest`
)

func getArcadeInfo() (*ArcadeInfo, error) {
	res, err := http.Get(overwatchArcade)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	info := &ArcadeInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return info, nil
}

func makeText(info *ArcadeInfo) (string, error) {
	return telegramMsg, nil
}