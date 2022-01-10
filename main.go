package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeAPI")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/common":
				var newItem string = createCommon()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, newItem)
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/uncommon":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createUncommon())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/rare":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createRare())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/very_rare":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, createVeryRare())
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "/legendary":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Легендарных предметов пока не подвезли")
				msg.ParseMode = "markdown"
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Undefined command")
				bot.Send(msg)
			}

		}
	}
}
