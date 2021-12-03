package equipment_request

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const getCommandLogTag = "GetCommand"

//Get handles /get command
func (c *equipmentRequestCommander) Get(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument id", getCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		return
	}

	equipmentRequest, err := c.equipmentRequestService.Get(ctx, idx)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to get equipment request by id", getCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
		)

		if errors.Is(err, equipment_request.ErrNoExistsEquipmentRequest) {
			c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Equipment request with this id does not exist: %d", idx))
			return
		}

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Fail to get equipment request with id: %d", idx))
		return
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, equipmentRequest.String())
}
