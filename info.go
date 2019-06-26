package main

import (
	"encoding/json"
	"fmt"
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
	telegramMsg     = `Daily 1\n%s\n%s\n\nDaily 2\n%s\n%s\n\nWeekly 1\n%s\n%s\n\nWeekly 2\n%s\n%s\n\nPermanent\n%s\n%s`
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

func makeText(info *ArcadeInfo) string {
	return fmt.Sprintf(telegramMsg,
		info.TileLarge.Players, info.TileLarge.Name,
		info.TileDaily.Players, info.TileDaily.Name,
		info.TileWeekly1.Players, info.TileWeekly1.Name,
		info.TileWeekly2.Players, info.TileWeekly2.Name,
		info.TilePermanent.Players, info.TilePermanent.Name)
}
