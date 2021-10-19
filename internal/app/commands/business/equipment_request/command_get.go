package equipment_request

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EquipmentRequestCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("invalid format of argument id %s: %v", args, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		c.sendMessage(msg)
		return
	}

	equipmentRequest, err := c.equipmentRequestService.Get(idx)
	if err != nil {
		log.Printf("fail to get equipment request with id %d: %v", idx, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Fail to get equipment request with id: %d", idx))
		c.sendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, equipmentRequest.String())
	c.sendMessage(msg)
}
