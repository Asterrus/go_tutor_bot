package main

import (
	"bot/internal/app/commands"
	"bot/internal/service/config"
	"bot/internal/service/logistic/product"
	"bot/internal/service/paginator"
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
	domen := os.Getenv("DOMEN")
	subdomen := os.Getenv("SUBDOMEN")
	if len(domen) == 0 || len(subdomen) == 0 {
		log.Panic("DOMEN or SUBDOMEN not provided")
	}
	config := config.NewConfig(&domen, &subdomen)

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	paginator := paginator.NewPaginator[*product.Product](5)
	productService := product.NewService()

	var commander commands.ProductCommander = commands.NewCommander(bot, productService, config, paginator)

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		commander.HandleUpdate(update)
	}
}
