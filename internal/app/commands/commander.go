package commands

import (
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// Мапа с регистрацией рабочих методов
var registredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommandRouter(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: service}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recover from panic: %#v", err)
		}
	}()

	if update.Message == nil {
		return
	}
	// Add logic here

	command, ok := registredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}

	//switch update.Message.Command() {
	//case "help":
	//	c.Help(update.Message)
	//case "list":
	//	c.List(update.Message)
	//default:
	//	c.Default(update.Message)

	//}
	//msg.ReplyToMessageID = update.Message.MessageID
}
