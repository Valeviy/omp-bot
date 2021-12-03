package equipment_request

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Help handles /help command
func (c *equipmentRequestCommander) Help(ctx context.Context, inputMessage *tgbotapi.Message) {
	c.sendMessage(ctx, inputMessage.Chat.ID, "/help__business__equipmentRequest - help\n"+
		"/list__business__equipmentRequest - list of all equipment requests\n"+
		"/get__business__equipmentRequest {id} - get one equipment request by id\n"+
		"/new__business__equipmentRequest {json} - create a new one equipment request by json\n"+
		"/edit_status__business__equipmentRequest {id} {status} - update status of existing equipment request by id\n"+
		"/edit_equipment_id__business__equipmentRequest {id} {equipment_id} - update equipment id of existing equipment request by id\n"+
		"/remove__business__equipmentRequest {id} - remove one equipment request by id")
}
