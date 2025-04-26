package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	var outputMsg string = "Here all the products:"
	for _, prod := range c.productService.List() {
		outputMsg += "\n" + prod.Tittle
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)
	//msg.ReplyToMessageID = inputMsg.MessageID
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "Some Data"),
		),
	)

	c.bot.Send(msg)
}

func init() {
	registredCommands["list"] = (*Commander).List
}
