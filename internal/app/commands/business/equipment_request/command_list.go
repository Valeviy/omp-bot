package equipment_request

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipment_request/pagination"
)

//List handles /list command
func (c *equipmentRequestCommander) List(ctx context.Context, inputMessage *tgbotapi.Message) {
	listPagination := pagination.NewListPagination(
		c.equipmentRequestService,
		c.cfg.PerPage,
		pagination.CallbackListData{Page: 0})

	msg, buttons := listPagination.GetMessageWithButtons(ctx)
	c.sendMessageWithButtons(ctx, inputMessage.Chat.ID, msg, buttons)
}
