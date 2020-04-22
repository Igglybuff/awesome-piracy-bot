package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mvdan/xurls"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type DiscordConfig struct {
	apiToken string
}

func (c DiscordConfig) watchDiscord() {
	dg, err := discordgo.New("Bot " + c.apiToken)

	if err != nil {
		log.Panicf("[DISCORD] Fatal error! %s", err)
	}

	dg.AddHandler(messageWatch)

	err = dg.Open()
	if err != nil {
		log.Panicf("[DISCORD] Error opening connection: %s", err)
	}

	log.Printf("[DISCORD] Logged in successfully as %s", dg.State.User.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Panicf("[DISCORD] Error closing connection: %s", err)
	}
}

func messageWatch(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	xurl := xurls.Relaxed()
	url := xurl.FindString(m.Content)
	if url != "" {
		log.Printf("[DISCORD] Valid URL found: %s", url)
	}
}

func Run(apiToken string) {
	c := DiscordConfig{
		apiToken: apiToken,
	}

	log.Printf("[DISCORD] Starting Discord watcher...")
	c.watchDiscord()
}
