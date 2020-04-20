package main

import (
	"awesome-piracy-bot/pkg/discord"
	"awesome-piracy-bot/pkg/telegram"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var goroutineDelta = make(chan int)

type Config struct {
	Telegram struct {
		Enabled  bool
		APIToken string
	}
	Discord struct {
		Enabled  bool
		APIToken string
	}
	Reddit struct {
		Enabled bool
	}
	IRC struct {
		Enabled bool
	}
}

func main() {
	// get configuration
	var name = "awesome-piracy-bot"
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", name))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", name))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Fatal error in config file: %s \n", err)
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		log.Panicf("Error parsing config file, %v \n", err)
	}

	// TODO: send URLs back to main() via a channel
	// TODO: add elasticsearch URL destination with metadata
	// TODO: add metadata to URLs, e.g. HTTP response, HTML title, protocol
	// start watchers
	go startTelegram(config.Telegram.APIToken, config.Telegram.Enabled)
	go startDiscord(config.Discord.APIToken, config.Discord.Enabled)
	numGoroutines := 0
	for diff := range goroutineDelta {
		numGoroutines += diff
		if numGoroutines == 0 {
			os.Exit(0)
		}
	}
}

func startTelegram(apiToken string, enabled bool) {
	if enabled != true {
		log.Printf("[INFO] Telegram disabled - skipping")
	} else {
		telegram.Run(apiToken)
	}
	goroutineDelta <- +1
	go f()
}

func startDiscord(apiToken string, enabled bool) {
	if enabled != true {
		log.Printf("[INFO] Discord disabled - skipping")
	} else {
		discord.Run(apiToken)
	}
	goroutineDelta <- +1
	go f()
}

func f() {
	goroutineDelta <- -1
}
