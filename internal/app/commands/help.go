package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	outputMessage := "Command list:\n"
	outputMessage += "/" + c.getCommandName("help") + " - help\n"
	outputMessage += "/" + c.getCommandName("list") + " - list of products\n"
	outputMessage += "/" + c.getCommandName("get") + " - Get product. Recieve 1 argument - Product ID\n"
	outputMessage += "/" + c.getCommandName("delete") + " - Delete product. Recieve 1 argument - Product ID\n"
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)
	c.bot.Send(msg)
}
