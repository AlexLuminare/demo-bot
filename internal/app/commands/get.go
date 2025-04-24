package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	//Строка с аргументами команды /Get
	args := inputMsg.CommandArguments()

	arg, ok := strconv.Atoi(args)
	if ok != nil {
		log.Println("Wrong args with /get command")
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Successfully parsed arg: %v", arg))
	//msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}

func init() {
	registredCommands["get"] = (*Commander).Get
}
