package commands

import (
	"bot/internal/service/config"
	"bot/internal/service/logistic/product"
	"bot/internal/service/paginator"
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProductCommander interface {
	Help(inputMessage *tgbotapi.Message)
	Get(inputMessage *tgbotapi.Message)
	List(inputMessage *tgbotapi.Message)
	Delete(inputMessage *tgbotapi.Message)
	New(inputMessage *tgbotapi.Message)
	Edit(inputMessage *tgbotapi.Message)
	HandleUpdate(tgbotapi.Update)
}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService product.ProductService
	config         *config.Config
	paginator      *paginator.Paginator[*product.Product]
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService product.ProductService,
	config *config.Config,
	paginator *paginator.Paginator[*product.Product],
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
		config:         config,
		paginator:      paginator,
	}

}
func (c *Commander) getCommandName(s string) string {
	return s + "__" + *c.config.Domen + "__" + *c.config.Subdomen
}
func (c *Commander) HandleCommand(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	switch inputMessage.Command() {

	case c.getCommandName("help"), "help":
		c.Help(inputMessage)
	case c.getCommandName("list"), "list":
		c.List(inputMessage)
	case c.getCommandName("get"), "get":
		c.Get(inputMessage)
	case c.getCommandName("delete"), "delete":
		c.Delete(inputMessage)
	case c.getCommandName("new"), "new":
		c.New(inputMessage)
	case c.getCommandName("edit"), "edit":
		c.Edit(inputMessage)
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
		data := ButtonData{}
		err := json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
		if err != nil {
			log.Fatalf("Error unmarshal button data: %s", err)
		}
		fmt.Printf("PAGE: %v \n", data.Page)

		c.RetrieveList(update.CallbackQuery.From.ID, data.Page)
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
