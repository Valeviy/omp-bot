package equipment_request

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/logger"
)

const defaultCommandLogTag = "DefaultCommand"

//Default handles unknown commands
func (c *equipmentRequestCommander) Default(ctx context.Context, inputMessage *tgbotapi.Message) {
	logger.InfoKV(ctx, fmt.Sprintf("%s: default message", defaultCommandLogTag),
		"from", inputMessage.From.UserName,
		"text", inputMessage.Text,
	)
	c.sendMessage(ctx, inputMessage.Chat.ID, fmt.Sprintf("You wrote: %s", inputMessage.Text))
}
