package main

import (
	"OverwatchArcadeBot/date"
	"OverwatchArcadeBot/owa"
	"log"
	"time"
)

func execute() error {
	info, err := owa.GetArcadeInfo()
	if err != nil {
		return err
	}
	if msg, err := owa.MakeText(info); err != nil {
		return err
	} else {
		if e := owa.Send(msg); e != nil {
			log.Fatal(e)
		}
		return nil
	}
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
