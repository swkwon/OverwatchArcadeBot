package main

import (
	"OverwatchArcadeBot/date"
	"OverwatchArcadeBot/owa"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	tokenKey        = "OVERWATCH_ARCADE_BOT_API_TOKEN"
	telegramChannel = "@overwatcharcade"
)

func execute() error {
	info, err := owa.GetArcadeInfo()
	if err != nil {
		log.Println(err)
		if e := owa.Send(telegramChannel, os.Getenv(tokenKey), fmt.Sprintf("ERROR:\n%s", err.Error())); e != nil {
			log.Fatal(e)
		}
		return err
	}
	msg := owa.MakeText(info)
	if e := owa.Send(telegramChannel, os.Getenv(tokenKey), msg); e != nil {
		log.Fatal(e)
	}

	return nil
}

func isTodayData() bool {
	dt := date.ReadFile()
	if dt == "" {
		return false
	}

	t, e := time.Parse(time.RFC3339, dt)
	if e != nil {
		return false
	}

	now := time.Now()
	if now.Year() == t.Year() && now.Month() == t.Month() && now.Day() == t.Day() {
		return true
	}
	return false
}

func main() {
	if isTodayData() {
		return
	}
	if e := execute(); e == nil {
		date.WriteFile()
	}
}
