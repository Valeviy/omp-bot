package equipment_request

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"github.com/pkg/errors"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const editStatusCommandLogTag = "EditStatusCommand"

//EditStatus handles /edit_status command
func (c *equipmentRequestCommander) EditStatus(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	updateParts := strings.SplitN(args, " ", 2)
	if len(updateParts) != 2 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid number of arguments", editStatusCommandLogTag),
			"args", args,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid number of arguments: it should be id (positive integer number) and status (string value)")
		return
	}

	idx, err := strconv.ParseUint(updateParts[0], 10, 64)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument id", editStatusCommandLogTag),
			"err", err,
			"equipmentRequestId", updateParts[0],
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		return
	}

	status := business.EquipmentRequestStatus(updateParts[1])

	if exists := business.EquipmentRequestStatusesContains(status); exists == false {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument status", editStatusCommandLogTag),
			"err", err,
			"equipmentRequestStatus", status,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Invalid format of argument status: it should be one of these values: %v", business.EquipmentRequestStatuses))
		return
	}

	result, err := c.equipmentRequestService.UpdateStatus(ctx, idx, status)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to update equipment request status", editStatusCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
			"status", status,
		)

		if errors.Is(err, equipment_request.ErrNoExistsEquipmentRequest) {
			c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Equipment request with this id does not exist: %d", idx))
			return
		}

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Fail to update equipment request with id and status: %d, %s", idx, status))
		return
	}

	resultMsg := fmt.Sprintf("Equipment request with id %d and status %s has been updated", idx, status)

	if result == false {
		resultMsg = fmt.Sprintf("Equipment request with id %d and status %s has not been updated", idx, status)
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, resultMsg)
}
