package equipment_request

import (
	"context"
	"fmt"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	equipmentRequestCommanderLogTag = "EquipmentRequestCommander"
	listCommand                     = "list"
)

//Commander is an entity which can handle commands and callbacks
type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

//EquipmentRequestCommander is a commander for equipment request subdomain
type equipmentRequestCommander struct {
	bot                     *tgbotapi.BotAPI
	cfg                     config.Bot
	equipmentRequestService equipment_request.EquipmentRequestService
}

//NewEquipmentRequestCommander returns a new EquipmentRequestCommander
func NewEquipmentRequestCommander(
	bot *tgbotapi.BotAPI,
	cfg config.Bot,
	equipmentRequestAPIClient pb.BssEquipmentRequestApiServiceClient,
	equipmentRequestFacadeAPIClient facadepb.BssEquipmentRequestFacadeApiServiceClient,
) Commander {
	equipmentRequestService := equipment_request.NewEquipmentRequestService(equipmentRequestAPIClient, equipmentRequestFacadeAPIClient)

	return &equipmentRequestCommander{
		bot:                     bot,
		cfg:                     cfg,
		equipmentRequestService: equipmentRequestService,
	}
}

func (c *equipmentRequestCommander) sendMessageWithButtons(ctx context.Context, chatID int64, info string, buttons []tgbotapi.InlineKeyboardButton) {
	msg := tgbotapi.NewMessage(chatID, info)

	if buttons != nil && len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: bot.Send failed send reply message to chat", equipmentRequestCommanderLogTag),
			"err", err,
		)
	}
}

func (c *equipmentRequestCommander) sendMessage(ctx context.Context, chatID int64, info string) {
	msg := tgbotapi.NewMessage(chatID, info)
	_, err := c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: bot.Send failed send reply message to chat", equipmentRequestCommanderLogTag),
			"err", err,
		)
	}
}

func (c *equipmentRequestCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case listCommand:
		c.CallbackList(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: callbackPath.CallbackName unknown callback name", equipmentRequestCommanderLogTag),
			"callbackPathSubdomain", callbackPath.Subdomain,
		)
	}
}

func (c *equipmentRequestCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(ctx, msg)
	case listCommand:
		c.List(ctx, msg)
	case "get":
		c.Get(ctx, msg)
	case "remove":
		c.Remove(ctx, msg)
	case "new":
		c.New(ctx, msg)
	case "edit_status":
		c.EditStatus(ctx, msg)
	case "edit_equipment_id":
		c.EditEquipmentID(ctx, msg)
	default:
		c.Default(ctx, msg)
	}
}
