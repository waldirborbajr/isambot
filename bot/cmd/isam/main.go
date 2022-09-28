package main

import (
	"localhost/isambot/bot/pkg/config"
	"localhost/isambot/bot/pkg/handlers"
	"localhost/isambot/bot/pkg/models"
	"net/http"

	"github.com/rs/zerolog/log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var cfg *models.ISamConfig

func init() {

	cfg = config.ReadConfig()

	log.Info().Msg("starting REDIS...")
	config.InitRedis()
	log.Info().Msg("Redis started.")

}

func main() {

	var webhook tgbotapi.WebhookConfig
	var updates tgbotapi.UpdatesChannel

	bot, err := tgbotapi.NewBotAPI(cfg.ISamToken)
	if err != nil {
		log.Fatal().Err(err)
	}

	bot.Debug = cfg.ISamDebug

	log.Info().Msg("Authorized on account - " + bot.Self.UserName)

	if cfg.ISamTLS {
		webhook, _ = tgbotapi.NewWebhookWithCert(cfg.ISamHookURL+bot.Token, tgbotapi.FilePath(cfg.ISamCert))
	} else {
		webhook, _ = tgbotapi.NewWebhook(cfg.ISamHookURL + bot.Token)
	}

	if _, err = bot.Request(webhook); err != nil {
		log.Fatal().Err(err)
	}

	updates = bot.ListenForWebhook("/" + bot.Token)

	if cfg.ISamTLS {
		go http.ListenAndServeTLS(":9090", "cert.pem", "key.pem", nil)
	} else {
		go http.ListenAndServe(":9090", nil)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal().Err(err)
	}

	if info.LastErrorDate != 0 {
		log.Info().Msg("Telegram callback failed: " + info.LastErrorMessage)
	}

	handlers.UpdatesHandler(bot, updates)
}
