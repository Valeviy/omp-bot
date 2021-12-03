package equipment_request

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

const newCommandLogTag = "NewCommand"

//New handles /new command
func (c *equipmentRequestCommander) New(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var parsedData business.EquipmentRequest
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of equipmen request json entity", newCommandLogTag),
			"err", err,
			"args", args,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of equipment request json entity")
		return
	}

	insertedID, err := c.equipmentRequestService.Create(ctx, parsedData)
	if err != nil {

		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to create equipment request", newCommandLogTag),
			"err", err,
			"parsedData", parsedData,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Fail to create equipment request with data: %v", parsedData))
		return
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Equipment request with id %d has been created", insertedID))
}
