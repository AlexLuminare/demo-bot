package main

import (
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CommandRouter struct {
	bot *tgbotapi.BotAPI
}

func NewCommandRouter(bot *tgbotapi.BotAPI) *CommandRouter {
	return &CommandRouter{bot: bot}
}

func (c *CommandRouter) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"help - help\n"+
			"list - список продуктов")
	msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}
func (c *CommandRouter) List(inputMsg *tgbotapi.Message, productService *product.Service) {
	var outputMsg string = "Here all the products:"
	for _, prod := range productService.List() {
		outputMsg += "\n" + prod.Tittle
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)
	//msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}

func (c *CommandRouter) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	c.bot.Send(msg)
}
