package equipment_request

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

func (c *EquipmentRequestCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var parsedData business.EquipmentRequest
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		log.Printf("invalid format of equipmen request json entity %s: %v", args, err)
		c.sendMessage(inputMessage.Chat.ID, "Invalid format of equipment request json entity")
		return
	}

	insertedId, err := c.equipmentRequestService.Create(parsedData)
	if err != nil {
		log.Printf("fail to create equipment request with:%v, %v", parsedData, err)
		c.sendMessage(inputMessage.Chat.ID, fmt.Sprintf("Fail to create equipment request with data: %v", parsedData))
		return
	}

	c.sendMessage(inputMessage.Chat.ID, fmt.Sprintf("Equipment request with id %d has been created", insertedId))
}
