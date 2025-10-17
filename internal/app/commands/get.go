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
	log.Printf("args string: %v", args)
	argsSlice := strings.Split(args, " ")
	log.Printf("args slice: %v", argsSlice)
	otputMessage := "Get with args: \n"
	for _, arg := range argsSlice {
		otputMessage += fmt.Sprintf("%s \n", arg)
	}

	firstArg := argsSlice[0]

	arg, err := strconv.Atoi(firstArg)
	log.Printf("Atoi result: %v, %v", arg, err)
	if err != nil {
		otputMessage += fmt.Sprintf("wrong first arg, must be int: '%v'", firstArg)
	} else {
		otputMessage += fmt.Sprintf("First Arg is valid %v \n", arg)
		item, err := c.productService.Get(arg)
		if err != nil {
			otputMessage += fmt.Sprintf("No product with id %v \n", arg)
		} else {
			otputMessage += fmt.Sprintf("Find product %v \n", item.Title)
		}
	}

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, otputMessage)
	c.bot.Send(msg)
}
