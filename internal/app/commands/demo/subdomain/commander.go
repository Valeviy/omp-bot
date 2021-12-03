package subdomain

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/demo/subdomain"
)

const (
	demoSubdomainCommanderLogTag = "DemoSubdomainCommander"
	listCommand                  = "list"
)

//Commander is an entity which can handle commands and callbacks
type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

//DemoSubdomainCommander is a commander for default subdomain
type demoSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *subdomain.Service
}

//NewDemoSubdomainCommander returns a new DemoSubdomainCommander
func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *demoSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &demoSubdomainCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

func (c *demoSubdomainCommander) sendMessage(ctx context.Context, chatID int64, info string) {
	msg := tgbotapi.NewMessage(chatID, info)
	_, err := c.bot.Send(msg)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: bot.Send failed send reply message to chat", demoSubdomainCommanderLogTag),
			"err", err,
		)
	}
}

func (c *demoSubdomainCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case listCommand:
		c.CallbackList(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: callbackPath.CallbackName unknown callback name", demoSubdomainCommanderLogTag),
			"callbackPathSubdomain", callbackPath.Subdomain,
		)
	}
}

func (c *demoSubdomainCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(ctx, msg)
	case listCommand:
		c.List(ctx, msg)
	case "get":
		c.Get(ctx, msg)
	default:
		c.Default(ctx, msg)
	}
}
