package router

import (
	"context"
	"fmt"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/app/commands/business"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/demo"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	routerLogTag   = "Router"
	businessDomain = "business"
	demoDomain     = "demo"
)

//Commander is a bot commander interface
type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, callback *tgbotapi.Message, commandPath path.CommandPath)
}

//Router is a command router
type Router struct {
	// bot
	bot *tgbotapi.BotAPI
	cfg config.Bot
	// demoCommander
	demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	businessCommander Commander
	// work
	// service
	// exchange
	// estate
	// rating
	// security
	// cinema
	// logistic
	// product
	// education
}

//NewRouter returns a new command router
func NewRouter(
	bot *tgbotapi.BotAPI,
	cfg config.Bot,
	equipmentRequestAPIClient pb.BssEquipmentRequestApiServiceClient,
	equipmentRequestFacadeAPIClient facadepb.BssEquipmentRequestFacadeApiServiceClient,
) *Router {
	return &Router{
		// bot
		bot: bot,
		cfg: cfg,
		// demoCommander
		demoCommander: demo.NewDemoCommander(bot),
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		businessCommander: business.NewBusinessCommander(bot, cfg, equipmentRequestAPIClient, equipmentRequestFacadeAPIClient),
		// work
		// service
		// exchange
		// estate
		// rating
		// security
		// cinema
		// logistic
		// product
		// education
	}
}

//HandleUpdate handles update item
func (c *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of equipmen request json entity", routerLogTag),
				"panicValue", panicValue,
			)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(ctx, update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(ctx, update.Message)
	}
}

func (c *Router) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: path.ParseCallback failed", routerLogTag),
			"err", err,
			"callbackData", callback.Data,
		)
		return
	}

	switch callbackPath.Domain {
	case demoDomain:
		c.demoCommander.HandleCallback(ctx, callback, callbackPath)
	//case "user":
	//	break
	//case "access":
	//	break
	//case "buy":
	//	break
	//case "delivery":
	//	break
	//case "recommendation":
	//	break
	//case "travel":
	//	break
	//case "loyalty":
	//	break
	//case "bank":
	//	break
	//case "subscription":
	//	break
	//case "license":
	//	break
	//case "insurance":
	//	break
	//case "payment":
	//	break
	//case "storage":
	//	break
	//case "streaming":
	//	break
	case businessDomain:
		c.businessCommander.HandleCallback(ctx, callback, callbackPath)
	//case "work":
	//	break
	//case "service":
	//	break
	//case "exchange":
	//	break
	//case "estate":
	//	break
	//case "rating":
	//	break
	//case "security":
	//	break
	//case "cinema":
	//	break
	//case "logistic":
	//	break
	//case "product":
	//	break
	//case "education":
	//	break
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: path.ParseCallback failed", routerLogTag),
			"callbackPathDomain", callbackPath.Domain,
		)
	}
}

func (c *Router) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(ctx, msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: path.ParseCommand failed", routerLogTag),
			"err", err,
			"msgCommand", msg.Command(),
		)
		return
	}

	switch commandPath.Domain {
	case demoDomain:
		c.demoCommander.HandleCommand(ctx, msg, commandPath)
	//case "user":
	//	break
	//case "access":
	//	break
	//case "buy":
	//	break
	//case "delivery":
	//	break
	//case "recommendation":
	//	break
	//case "travel":
	//	break
	//case "loyalty":
	//	break
	//case "bank":
	//	break
	//case "subscription":
	//	break
	//case "license":
	//	break
	//case "insurance":
	//	break
	//case "payment":
	//	break
	//case "storage":
	//	break
	//case "streaming":
	//	break
	case businessDomain:
		c.businessCommander.HandleCommand(ctx, msg, commandPath)
	//case "work":
	//	break
	//case "service":
	//	break
	//case "exchange":
	//	break
	//case "estate":
	//	break
	//case "rating":
	//	break
	//case "security":
	//	break
	//case "cinema":
	//	break
	//case "logistic":
	//	break
	//case "product":
	//	break
	//case "education":
	//	break
	default:
		logger.ErrorKV(ctx, fmt.Sprintf("%s: path.ParseCommand failed", routerLogTag),
			"commandPathDomain", commandPath.Domain,
		)
	}
}

func (c *Router) showCommandFormat(ctx context.Context, inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: bot.Send failed send reply message to chat", routerLogTag),
			"err", err,
		)
	}
}
