package commands

import (
	"bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func (c *Commander) HandleCommand(inputMessage *tgbotapi.Message) {
	switch inputMessage.Command() {
	case "help":
		c.Help(inputMessage)
	case "list":
		c.List(inputMessage)
	default:
		c.UnknownCommand(inputMessage)
	}
}
