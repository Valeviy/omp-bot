package equipment_request

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EquipmentRequestCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	c.sendMessage(inputMessage.Chat.ID, fmt.Sprintf("You wrote: %s", inputMessage.Text))
}
