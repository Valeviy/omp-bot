package equipment_request

import (
	"encoding/json"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipment_request/pagination"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *EquipmentRequestCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var parsedData pagination.CallbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("EquipmentRequestCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		c.sendMessage(callback.Message.Chat.ID, "Unable to get list of equipment requests for selected page")
		return
	}

	listPagination := pagination.NewListPagination(
		c.equipmentRequestService,
		pagination.ListPerPageDefault,
		parsedData)

	msg, buttons := listPagination.GetMessageWithButtons()
	c.sendMessageWithButtons(callback.Message.Chat.ID, msg, buttons)
}
