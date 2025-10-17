package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsSlice := strings.Split(args, " ")
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	if len(argsSlice) == 0 {
		err_msg := "Product ID not provided!"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	} else if len(argsSlice) > 1 {
		err_msg := "Command expects 1 argument"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}

	arg, err := strconv.Atoi(argsSlice[0])
	log.Printf("Atoi result: %v, %v", arg, err)
	if err != nil {
		err_msg := fmt.Sprintf("wrong first arg, must be int: '%v'", arg)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}
	item, err := c.productService.Get(arg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	otputMessage := fmt.Sprintf("Find product %v \n", item.Title)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, otputMessage)
	c.bot.Send(msg)
}
