package equipment_request

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EquipmentRequestCommander) Remove(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("invalid format of argument id %s: %v", args, err)
		c.sendMessage(inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		return
	}

	result, err := c.equipmentRequestService.Remove(idx)
	if err != nil {
		log.Printf("fail to remove equipment request with id %d: %v", idx, err)
		c.sendMessage(inputMessage.Chat.ID, fmt.Sprintf("Fail to remove equipment request with id: %d", idx))
		return
	}

	resultMsg := fmt.Sprintf("Equipment request with id %d has been deleted", idx)

	if result == false {
		resultMsg = fmt.Sprintf("Equipment request with id %d has not been deleted", idx)
	}

	c.sendMessage(inputMessage.Chat.ID, resultMsg)
}
