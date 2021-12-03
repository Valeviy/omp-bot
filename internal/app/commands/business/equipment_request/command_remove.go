package equipment_request

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"github.com/pkg/errors"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const removeCommandLogTag = "RemoveCommand"

//Remove handles remove command
func (c *equipmentRequestCommander) Remove(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument id", removeCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		return
	}

	result, err := c.equipmentRequestService.Remove(ctx, idx)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to remove equipment request by id", removeCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
		)

		if errors.Is(err, equipment_request.ErrNoExistsEquipmentRequest) {
			c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Equipment request with this id does not exist: %d", idx))
			return
		}

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Fail to remove equipment request with id: %d", idx))
		return
	}

	resultMsg := fmt.Sprintf("Equipment request with id %d has been deleted", idx)

	if result == false {
		resultMsg = fmt.Sprintf("Equipment request with id %d has not been deleted", idx)
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, resultMsg)
}
