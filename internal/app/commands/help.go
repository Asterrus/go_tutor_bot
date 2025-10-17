package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	outputMessage := "Command list:\n"
	outputMessage += "/help - help\n"
	outputMessage += "/list - list of products\n"
	outputMessage += "/get - get product by id\n"
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)
	c.bot.Send(msg)
}
