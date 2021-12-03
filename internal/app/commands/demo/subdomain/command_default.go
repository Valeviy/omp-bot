package subdomain

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const defaultCommandLogTag = "DefaultCommand"

//Default handles unknown commands
func (c *demoSubdomainCommander) Default(ctx context.Context, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	logger.InfoKV(ctx, fmt.Sprintf("%s: default message", defaultCommandLogTag),
		"from", inputMessage.From.UserName,
		"text", inputMessage.Text,
	)

	c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("You wrote: %s", inputMessage.Text))

}
