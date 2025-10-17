package commands

import (
	"bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(commander *Commander, message *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot: bot,
	}

}

func (c *Commander) HandleUpdate(inputUpdate tgbotapi.Update) {
	inputMessage := inputUpdate.Message
	if inputMessage == nil {
		return
	}
	if inputMessage.IsCommand() {
		c.HandleCommand(inputMessage)
	} else {
		c.messageHandler(inputMessage)
	}

}

func (c *Commander) HandleCommand(inputMessage *tgbotapi.Message) {

	handler, ok := registeredCommands[inputMessage.Command()]
	if !ok {
		c.UnknownCommand(inputMessage)
	}
	handler(c, inputMessage)
}
