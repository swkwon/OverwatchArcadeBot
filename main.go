package main

import (
	"log"
	"os"
)

const (
	tokenKey        = "OVERWATCH_ARCADE_BOT_API_TOKEN"
	telegramChannel = "@overwatcharcade"
)

func main() {
	info, err := getArcadeInfo()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	msg := makeText(info)
	send(telegramChannel, os.Getenv(tokenKey), msg)
}
