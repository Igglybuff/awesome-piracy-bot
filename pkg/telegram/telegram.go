package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mvdan/xurls"
	"log"
)

type TelegramConfig struct {
	apiToken string
}

func (c TelegramConfig) watchTelegram() {
	bot, err := tgbotapi.NewBotAPI(c.apiToken)

	if err != nil {
		log.Panicf("[TELEGRAM] Fatal error! %s", err)
	}

	bot.Debug = false
	log.Printf("[TELEGRAM] Logged in successfully as %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		xurl := xurls.Relaxed()
		url := xurl.FindString(update.Message.Text)
		if url != "" {
			log.Printf("[TELEGRAM] Valid URL found: %s", url)
		}
	}
}

func Run(apiToken string) {
	c := TelegramConfig{
		apiToken: apiToken,
	}

	log.Printf("[TELEGRAM] Starting Telegram watcher...")
	c.watchTelegram()
}
