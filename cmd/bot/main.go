package main

import (
	"bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	productService := product.NewService()

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			handleCommand(bot, update.Message, productService)
		} else {
			messageHandler(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help"+"\n/list - list of products")
	bot.Send(msg)
}

func unknownCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "unknown command")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productServ *product.Service) {
	outputMsg := "Here all the products: \n\n"
	products := productServ.List()
	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	bot.Send(msg)
}
func handleCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productServ *product.Service) {
	switch inputMessage.Command() {
	case "help":
		helpCommand(bot, inputMessage)
	case "list":
		listCommand(bot, inputMessage, productServ)
	default:
		unknownCommand(bot, inputMessage)
	}
}

func messageHandler(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote message: "+inputMessage.Text)
	bot.Send(msg)
}
