package business

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/equipment_request"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/metrics"
)

const (
	businessCommanderLogTag   = "BusinessCommander"
	equipmentRequestSubdomain = "equipmentRequest"
)

//Commander is an entity which can handle commands and callbacks
type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

//BusinessCommander is a commander for business domain
type businessCommander struct {
	bot                       *tgbotapi.BotAPI
	equipmentRequestCommander Commander
}

//NewBusinessCommander returns a new BusinessCommander
func NewBusinessCommander(
	bot *tgbotapi.BotAPI,
	cfg config.Bot,
	equipmentRequestAPIClient pb.BssEquipmentRequestApiServiceClient,
	equipmentRequestFacadeAPIClient facadepb.BssEquipmentRequestFacadeApiServiceClient,
) Commander {
	return &businessCommander{
		bot:                       bot,
		equipmentRequestCommander: equipment_request.NewEquipmentRequestCommander(bot, cfg, equipmentRequestAPIClient, equipmentRequestFacadeAPIClient),
	}
}

func (c *businessCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case equipmentRequestSubdomain:
		c.equipmentRequestCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: subdomainCommander.HandleCallback unknown subdomain", businessCommanderLogTag),
			"callbackPathSubdomain", callbackPath.Subdomain,
		)
	}
}

func (c *businessCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	metrics.IncTotalCommandCalls(commandPath.String())

	switch commandPath.Subdomain {
	case equipmentRequestSubdomain:
		c.equipmentRequestCommander.HandleCommand(ctx, msg, commandPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: subdomainCommander.HandleCommand unknown command", businessCommanderLogTag),
			"commandPathSubdomain", commandPath.Subdomain,
		)
	}
}
