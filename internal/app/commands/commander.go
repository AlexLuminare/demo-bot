package commands

import (
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommandRouter(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: service}
}
