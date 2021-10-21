package equipment_request

import (
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type EquipmentRequestCommander struct {
	bot                     *tgbotapi.BotAPI
	equipmentRequestService equipment_request.EquipmentRequestService
}

func NewEquipmentRequestCommander(
	bot *tgbotapi.BotAPI,
) *EquipmentRequestCommander {
	equipmentRequestService := equipment_request.NewDummyEquipmentRequestService()

	return &EquipmentRequestCommander{
		bot:                     bot,
		equipmentRequestService: equipmentRequestService,
	}
}

func (c *EquipmentRequestCommander) sendMessageWithButtons(chatId int64, info string, buttons []tgbotapi.InlineKeyboardButton) {
	msg := tgbotapi.NewMessage(chatId, info)

	if buttons != nil && len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EquipmentRequestCommander.sendMessageWithButtons: error sending reply message to chat - %v", err)
	}
}

func (c *EquipmentRequestCommander) sendMessage(chatId int64, info string) {
	msg := tgbotapi.NewMessage(chatId, info)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EquipmentRequestCommander.sendMessage: error sending reply message to chat - %v", err)
	}
}

func (c *EquipmentRequestCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("EquipmentRequestCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EquipmentRequestCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "remove":
		c.Remove(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
