package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/am3o/discordbot/pkg/service"
)

const DictonaryPath = "./resources/dictonary.json"

func main() {
	logger := logrus.StandardLogger()

	token, exists := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !exists {
		logger.Fatal("Could not start discord bot without any token")
		return
	}

	bot, err := service.New(token, DictonaryPath, logger)
	if err != nil {
		logger.WithError(err).Error("Could not initialize the bot")
	}
	defer bot.Close()

	if err := bot.ListenAndServe(context.Background()); err != nil {
		logger.WithError(err).Error("Could not listen any more to the discord session ")
		return
	}
	logger.Info("Discord bot is stopped")
}
