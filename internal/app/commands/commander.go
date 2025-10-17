package commands

import (
	"bot/internal/service/config"
	"bot/internal/service/product"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
	config         *config.Config
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
	config *config.Config,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
		config:         config,
	}

}
func (c *Commander) getCommandName(s string) string {
	return s + "__" + *c.config.Domen + "__" + *c.config.Subdomen
}
func (c *Commander) HandleCommand(inputMessage *tgbotapi.Message) {
	switch inputMessage.Command() {
	case c.getCommandName("help"), "help":
		c.Help(inputMessage)
	case c.getCommandName("list"), "list":
		c.List(inputMessage)
	case c.getCommandName("get"), "get":
		c.Get(inputMessage)
	default:
		c.UnknownCommand(inputMessage)
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	// defer func() {
	// 	if panicValue := recover(); panicValue != nil {
	// 		fmt.Printf("recover from panic: %v", panicValue)
	// 	}
	// }()
	fmt.Printf("HandleUpdate\n CallbackQuery?: %v\n Message?: %v\n", update.CallbackQuery != nil, update.Message != nil)
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Data: "+update.CallbackQuery.Data)
		c.bot.Send(msg)
		return
	}
	if update.Message == nil {
		return
	}
	if update.Message.IsCommand() {
		c.HandleCommand(update.Message)
	} else {
		c.HandleMessage(update.Message)
	}
}
