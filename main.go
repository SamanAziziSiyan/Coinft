package main

import (
	"coinft/bot"
	"log"
)

func main() {
	log.Println("Starting Telegram Bot...")

	// Replace with your actual bot token
	botToken := "6753698762:AAG8kUyZqfnsPtRz4bAga1y-l5WbW0q3vG8"

	telegramBot, err := bot.NewTelegramBot(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	telegramBot.Start()
}
