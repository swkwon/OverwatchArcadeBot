package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	overwatchArcade = "https://overwatcharcade.today/api/today"
)

type Item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players string `json:"players"`
	Code    string `json:"code"`
}

type ArcadeInfo struct {
	UpdateAt      string `json:"updated_at"`
	TileLarge     Item   `json:"tile_large"`
	TileWeekly1   Item   `json:"tile_weekly_1"`
	TileDaily     Item   `json:"tile_daily"`
	TileWeekly2   Item   `json:"tile_weekly_2"`
	TilePermanent Item   `json:"tile_permanent"`
}

func main() {
	res, err := http.Get(overwatchArcade)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	arcadeInfo := &ArcadeInfo{}

	if e := json.Unmarshal(bytes, arcadeInfo); e != nil {
		fmt.Printf("error on unmarshal: %#v", e)
	}

	fmt.Printf("%#v\n", arcadeInfo)
}
