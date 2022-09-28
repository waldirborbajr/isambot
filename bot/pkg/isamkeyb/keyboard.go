package isamkeyb

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Define START button
var StartKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Iniciar"),
	),
)

// Defining keyboard mapping
var MainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Pedidos"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Pagamento"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Encerrar"),
	),
)

var GraphKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("15"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Voltar"),
	),
)
