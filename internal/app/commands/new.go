package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsSlice := strings.Split(args, " ")
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	if len(argsSlice) == 0 || (len(argsSlice) == 1 && len(argsSlice[0]) == 0) {
		err_msg := "Arguments not provided!"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	} else if len(argsSlice) > 2 {
		err_msg := "Command expects no more than 2 arguments"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}
	fmt.Println(argsSlice, len(argsSlice), argsSlice[0])
	name := argsSlice[0]
	price := argsSlice[1]

	// name
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		err_msg := "Name not provided!"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}

	// price
	validPrice, err := strconv.ParseFloat(price, 64)
	log.Printf("Float parse result: %v, %v", validPrice, err)
	if err != nil {
		err_msg := fmt.Sprintf("Price is not valid: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}

	newProductID := c.productService.New(name, validPrice)
	newProduct, _ := c.productService.Get(newProductID)
	otputMessage := fmt.Sprintf("Created %s ID: %v \n", newProduct.Title, newProduct.ID)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, otputMessage)
	c.bot.Send(msg)
}
