package equipment_request

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const editEquipmentIDStatusCommandLogTag = "EditEquipmentIdStatusCommand"

//EditEquipmentID handles /edit_equipment_id command
func (c *equipmentRequestCommander) EditEquipmentID(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	updateParts := strings.SplitN(args, " ", 2)
	if len(updateParts) != 2 {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid number of arguments", editEquipmentIDStatusCommandLogTag),
			"args", args,
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid number of arguments: it should be id (positive integer number) and equipment id (positive integer number)")
		return
	}

	idx, err := strconv.ParseUint(updateParts[0], 10, 64)
	if err != nil {
		log.Printf("invalid format of argument id %s: %v", args, err)
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument id", editEquipmentIDStatusCommandLogTag),
			"err", err,
			"equipmentRequestId", updateParts[0],
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of argument id: it should be positive integer number")
		return
	}

	equipmentIdx, err := strconv.ParseUint(updateParts[1], 10, 64)
	if err != nil {
		log.Printf("invalid format of argument equipment id %s: %v", args, err)
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument equipment id", editEquipmentIDStatusCommandLogTag),
			"err", err,
			"equipmentId", updateParts[1],
		)

		c.sendMessage(ctx, inputMessage.Chat.ID, "Invalid format of argument equipment id: it should be positive integer number")
		return
	}

	result, err := c.equipmentRequestService.UpdateEquipmentID(ctx, idx, equipmentIdx)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to update equipment id", editStatusCommandLogTag),
			"err", err,
			"equipmentRequestId", idx,
			"equipmentId", equipmentIdx,
		)

		if errors.Is(err, equipment_request.ErrNoExistsEquipmentRequest) {
			c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Equipment request with this id does not exist: %d", idx))
			return
		}

		c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("Fail to update equipment request with id and equipment id: %d, %d", idx, equipmentIdx))
		return
	}

	resultMsg := fmt.Sprintf("Equipment request with id %d and equipment id %d has been updated", idx, equipmentIdx)

	if result == false {
		resultMsg = fmt.Sprintf("Equipment request with id %d and equipment id %d has not been updated", idx, equipmentIdx)
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, resultMsg)
}
