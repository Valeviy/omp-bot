package subdomain

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Help handles /help command
func (c *demoSubdomainCommander) Help(ctx context.Context, inputMessage *tgbotapi.Message) {
	c.sendMessage(ctx, inputMessage.Chat.ID, "/help - help\n"+
		"/list - list products")
}
