package equipment_request

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipment_request/pagination"
)

func (c *EquipmentRequestCommander) List(inputMessage *tgbotapi.Message) {
	listPagination := pagination.NewListPagination(
		c.equipmentRequestService,
		pagination.ListPerPageDefault,
		pagination.CallbackListData{Page: 0})

	msg, buttons := listPagination.GetMessageWithButtons()
	c.sendMessageWithButtons(inputMessage.Chat.ID, msg, buttons)
}
