package main

import (
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"help - help\n"+
			"list - список продуктов")
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productService *product.Service) {
	var outputMsg string = "Here all the products:"
	for _, product := range productService.List() {
		outputMsg += "\n" + product.Tittle
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)
	//msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}

func DefaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	bot.Send(msg)
}
