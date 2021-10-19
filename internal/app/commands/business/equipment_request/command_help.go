package equipment_request

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EquipmentRequestCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help__business__equipmentRequest - help\n"+
		"/list__business__equipmentRequest - list of all equipment requests\n"+
		"/get__business__equipmentRequest {id} - get one equipment request by id\n"+
		"/new__business__equipmentRequest {json} - create a new one equipment request by json\n"+
		"/edit__business__equipmentRequest {id} {json} - update existing equipment request by id\n"+
		"/remove__business__equipmentRequest {id} - remove one equipment request by id")

	c.sendMessage(msg)
}
