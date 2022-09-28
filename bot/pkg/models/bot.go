package models

import (
	"github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ISamBot struct {
	BotApi *tgbotapi.BotAPI
	Msg    *tgbotapi.MessageConfig
	Upd    *tgbotapi.Update
	Db     database
}

type database struct {
	Db *redis.Client
}
