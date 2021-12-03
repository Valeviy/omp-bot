package demo

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/demo/subdomain"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
)

const (
	demoCommanderLogTag = "DemoCommander"
	demoSubdomain       = "subdomain"
)

//Commander is an entity which can handle commands and callbacks
type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

//DemoCommander is a commander for default domain
type demoCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

//NewDemoCommander returns a new DemoCommander
func NewDemoCommander(
	bot *tgbotapi.BotAPI,
) Commander {
	return &demoCommander{
		bot: bot,
		// subdomainCommander
		subdomainCommander: subdomain.NewDemoSubdomainCommander(bot),
	}
}

func (c *demoCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case demoSubdomain:
		c.subdomainCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: subdomainCommander.HandleCallback unknown subdomain", demoCommanderLogTag),
			"callbackPathSubdomain", callbackPath.Subdomain,
		)
	}
}

func (c *demoCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case demoSubdomain:
		c.subdomainCommander.HandleCommand(ctx, msg, commandPath)
	default:
		logger.InfoKV(ctx, fmt.Sprintf("%s: subdomainCommander.HandleCommand unknown command", demoCommanderLogTag),
			"commandPathSubdomain", commandPath.Subdomain,
		)
	}
}
