package AdsBot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	AdsBot "github.com/lexkanev/AdsBot/src"
)

type Config struct {
	TelegramBotToken string
}

//read config file
file, _ := os.Open("../config.json")
decoder := json.NewDecoder(file)
configuration := Config{}
err := decoder.Decode(&configuration)
if err != nil {
	log.Panic(err)
}
fmt.Println(configuration.TelegramBotToken)

AdsBot.SkeyBot(configuration.TelegramBotToken)
