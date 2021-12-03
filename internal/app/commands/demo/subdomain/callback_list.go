package subdomain

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
)

const callbackListLogTag = "CallbackList"

func (c *demoSubdomainCommander) CallbackList(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: error reading json data for type CallbackListData from input string", callbackListLogTag),
			"err", err,
			"callbackPathCallbackData", callbackPath.CallbackData,
		)

		return
	}

	c.sendMessage(ctx, callback.Message.Chat.ID, fmt.Sprintf("Parsed: %+v\n", parsedData))
}
