# awesome-piracy-bot

This is a WIP set of bots which scrape URLs from messages sent to Telegram groups/channels, Discord server channels, subreddit posts/comments, and IRC channels. Written in GoLang.

## Installation

```
$ git clone https://github.com/igglybuff/awesome-piracy-bot.git
$ cd awesome-piracy-bot
$ go build
```

## Configuration

`awesome-piracy-bot` takes a YAML file for config, ideally stored in `$HOME/.config/awesome-piracy-bot/config.yaml`.

The structure is as follows:

```
discord:
  enabled: true
  apiToken: "YOUR_BOT_TOKEN_HERE"

telegram:
  enabled: true
  apiToken: "YOUR_BOT_TOKEN_HERE"

irc:
  enabled: false

reddit:
  enabled: false
```

Currently only Discord and Telegram have been implemented.