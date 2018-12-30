package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	hbot "github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
)

// Mods is a list of all of the channel mods.
// TODO: Get this upon connection by using the /mods command
var Mods = []string{"selthor"}

func main() {
	logger := log.LvlFilterHandler(log.LvlInfo, log.StdoutHandler)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Info("Env vars loaded")

	log.Info("Initializing Database")
	err = initDatabase()
	if err != nil {
		panic(err)
	}

	botName := os.Getenv("TWITCH_BOT_NAME")
	bot, err := hbot.NewBot("irc.chat.twitch.tv:6667", botName, options)
	if err != nil {
		panic(err)
	}
	bot.Logger.SetHandler(logger)

	log.Info(fmt.Sprintf("Bot created: %s", bot.String()))

	log.Info("Connecting to Twitch chat")
	bot.AddTrigger(CreateCommandTrigger)
	bot.AddTrigger(CommandTrigger)
	bot.AddTrigger(ListTrigger)
	bot.AddTrigger(RemoveTrigger)
	bot.Run()
	log.Info("Disconnected")
}

func options(bot *hbot.Bot) {
	password := os.Getenv("TWITCH_OAUTH_PASSWORD")
	channel := os.Getenv("TWITCH_CHANNEL")
	bot.Password = password
	bot.Channels = []string{channel}
}
