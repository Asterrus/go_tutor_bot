package commands

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Direction int

const (
	Next int = iota
	Previous
)

type ButtonData struct {
	Page      int
	Direction Direction
}

func (c *Commander) List(userID int64, page int) {
	outputMsg := fmt.Sprintf("Here all the products on page %d: \n\n", page)
	allProducts := c.productService.List()
	products := c.paginator.GetPaginatedItems(c.productService.List(), page)
	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(userID, outputMsg)

	nextButtonRequired := c.paginator.TotalPages(allProducts) > page
	previousButtonRequired := page != 1
	fmt.Println("Next", nextButtonRequired, "Prev", previousButtonRequired)
	if nextButtonRequired || previousButtonRequired {

		if nextButtonRequired && !previousButtonRequired {
			fmt.Println(1)
			nextButton := ButtonData{
				Page:      page + 1,
				Direction: Direction(Next),
			}
			nextButtonData, err := json.Marshal(nextButton)

			if err != nil {
				log.Panicf("Button json encoding error: %s", err)
			}
			replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Next page", string(nextButtonData)),
				))
			msg.ReplyMarkup = replyMarkup

		} else if !nextButtonRequired && previousButtonRequired {
			fmt.Println(2)
			previousButton := ButtonData{
				Page:      page - 1,
				Direction: Direction(Previous),
			}
			previousButtonData, err := json.Marshal(previousButton)

			if err != nil {
				log.Panicf("Button json encoding error: %s", err)
			}
			replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Previous page", string(previousButtonData)),
				))
			msg.ReplyMarkup = replyMarkup

		} else {
			fmt.Println(3)
			nextButton := ButtonData{
				Page:      page + 1,
				Direction: Direction(Next),
			}
			nextButtonData, next_err := json.Marshal(nextButton)
			if next_err != nil {
				log.Panicf("Button json encoding error: %s", next_err)
			}
			previousButton := ButtonData{
				Page:      page - 1,
				Direction: Direction(Previous),
			}

			previousButtonData, prev_err := json.Marshal(previousButton)
			if prev_err != nil {
				log.Panicf("Button json encoding error: %s", prev_err)
			}
			replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Previous page", string(previousButtonData)),
					tgbotapi.NewInlineKeyboardButtonData("Next page", string(nextButtonData)),
				))
			msg.ReplyMarkup = replyMarkup

		}
	}
	c.bot.Send(msg)

}
