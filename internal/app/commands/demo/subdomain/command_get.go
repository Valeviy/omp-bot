package subdomain

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/logger"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const getCommandLogTag = "GetCommand"

//Get handles /get command
func (c *demoSubdomainCommander) Get(ctx context.Context, inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid format of argument id", getCommandLogTag),
			"err", err,
			"idx", idx,
		)

		return
	}

	product, err := c.subdomainService.Get(idx)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: fail to get product by id", getCommandLogTag),
			"err", err,
			"idx", idx,
		)

		return
	}

	c.sendMessage(ctx, inputMessage.Chat.ID, product.Title)
}
