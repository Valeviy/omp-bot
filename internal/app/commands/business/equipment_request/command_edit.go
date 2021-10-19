package equipment_request

import (
	"encoding/json"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EquipmentRequestCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	updateParts := strings.SplitN(args, " ", 2)
	if len(updateParts) != 2 {
		log.Printf("invalid number of arguments %s", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid number of arguments: it should be id (positive integer number) and json with equipment request data")
		c.sendMessage(msg)
		return
	}

	idx, err := strconv.ParseUint(updateParts[0], 10, 64)
	if err != nil {
		log.Printf("invalid format of argument id %s: %v", args, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		c.sendMessage(msg)
		return
	}

	updateData := business.EquipmentRequest{}
	err = json.Unmarshal([]byte(updateParts[1]), &updateData)
	if err != nil {
		log.Printf("invalid format of equipmen request json entity %s: %v", args, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid format of equipment request json entity")
		c.sendMessage(msg)
		return
	}

	err = c.equipmentRequestService.Update(idx, updateData)
	if err != nil {
		log.Printf("fail to update equipment request with id %d: %v", idx, err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Fail to update equipment request with id: %d", idx))
		c.sendMessage(msg)
		return
	}

	resultMsg := fmt.Sprintf("Equipment request with id %d has been updated", idx)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, resultMsg)
	c.sendMessage(msg)
}
