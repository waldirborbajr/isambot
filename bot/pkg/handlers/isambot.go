package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/rs/zerolog/log"
)

func UpdatesHandler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.CallbackQuery != nil {
			bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, "It's cool!"))
		}

		if update.Message != nil { // If we got a message
			chatID := update.Message.Chat.ID
			messageID := update.Message.MessageID
			msg := tgbotapi.NewMessage(chatID, "")

			// Process commands from user select
			if update.Message.IsCommand() {
				switch update.Message.Text {
				case "/start":
					msg.Text = MessagesHandler("Welcome", update.Message.From.UserName)
				}
			} else if update.Message.Text != "" {
				switch update.Message.Text {
				case "oi":
					msg.Text = "Voce escreveu [oi]\n"
				default:
					if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, messageID)); err != nil {
						log.Debug()
					}
					msg.Text = "only allowed commands and messages."
				}
			}

			if _, err := bot.Send(msg); err != nil {
				log.Debug()
			}
		}
	}
}
