package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsSlice := strings.Split(args, " ")
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	if len(argsSlice) == 0 || (len(argsSlice) == 1 && len(argsSlice[0]) == 0) {
		err_msg := "Arguments not provided!"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	} else if len(argsSlice) > 3 {
		err_msg := "Command expects no more than 2 arguments"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	} else if len(argsSlice) != 3 {
		err_msg := "ID, Name and Price required!"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}
	fmt.Println(argsSlice, len(argsSlice), argsSlice[0])
	id := argsSlice[0]
	title := argsSlice[1]
	price := argsSlice[2]

	//id
	validID, err := strconv.Atoi(id)
	if err != nil {
		err_msg := fmt.Sprintf("ID is not valid: %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}
	// title
	title = strings.TrimSpace(title)
	if len(title) == 0 {
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

	editedProduct, err := c.productService.Get(validID)
	if err != nil {
		err_msg := fmt.Sprintf("Product with ID: %v not found", validID)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err_msg)
		c.bot.Send(msg)
		return
	}

	c.productService.Edit(editedProduct.ID, title, validPrice)
	otputMessage := fmt.Sprintf("Product with ID: %v changed\n", editedProduct.ID)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, otputMessage)
	c.bot.Send(msg)
}
