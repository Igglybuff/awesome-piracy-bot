package main

import (
	"awesome-piracy-bot/pkg/discord"
	"awesome-piracy-bot/pkg/telegram"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"reflect"
)

var goroutineDelta = make(chan int)

type Config struct {
	Watchers struct {
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
}

type WatcherConfig struct {
	Type     string
	APIToken string
}

type WatcherConfigs []WatcherConfig

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

	v := reflect.ValueOf(config.Watchers).Type()
	var watcherConfigs WatcherConfigs

	// generate a list of WatcherConfigs
	for i := 0; i < v.NumField(); i++ {
		c := reflect.ValueOf(config.Watchers).Field(i)
		enabled := c.FieldByName("Enabled")
		token := c.FieldByName("APIToken")
		// create a WatcherConfig with this iteration's watcher type
		if enabled.Bool() == true {
			watcherConfig := WatcherConfig{
				Type:     v.Field(i).Name,
				APIToken: token.String(),
			}
			watcherConfigs = append(watcherConfigs, watcherConfig)
		}
	}

	numGoroutines := 0

	//urls := make(chan string)

	// start watchers
	for _, w := range watcherConfigs {
		go w.startWatcher()
	}

	// start listening to watcher channels for URLs

	// run metadata tasks for incoming URLs

	// upload URL with metadata to ElasticSearch

	for diff := range goroutineDelta {
		numGoroutines += diff
		if numGoroutines == 0 {
			os.Exit(0)
		}
	}
}

func (c WatcherConfig) startWatcher() {
	if c.Type == "Telegram" {
		telegram.Run(c.APIToken)
	}
	if c.Type == "Discord" {
		discord.Run(c.APIToken)
	}
	goroutineDelta <- +1
	go f()
}

func f() {
	goroutineDelta <- -1
}
