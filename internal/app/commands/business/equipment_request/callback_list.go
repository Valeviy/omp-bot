package equipment_request

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipment_request/pagination"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
)

const callbackListLogTag = "CallbackList"

func (c *equipmentRequestCommander) CallbackList(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var parsedData pagination.CallbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: error reading json data for type CallbackListData from input string", callbackListLogTag),
			"err", err,
			"callbackPath.CallbackData", callbackPath.CallbackData,
		)

		c.sendMessage(ctx, callback.Message.Chat.ID, "Unable to get list of equipment requests for selected page")
		return
	}

	listPagination := pagination.NewListPagination(
		c.equipmentRequestService,
		c.cfg.PerPage,
		parsedData)

	msg, buttons := listPagination.GetMessageWithButtons(ctx)
	c.sendMessageWithButtons(ctx, callback.Message.Chat.ID, msg, buttons)
}
