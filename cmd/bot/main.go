package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable not set")
	}
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bot is now running")
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updates := botApi.GetUpdatesChan(update)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		log.Printf("[%s] %s", update.Message.From.UserName, text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		_, err := botApi.Send(msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
