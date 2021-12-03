package ompBot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	routerPkg "github.com/ozonmp/omp-bot/internal/app/router"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
)

const ompBotLogTag = "OmpBot"

// OmpBot is a public interface for equipment request facade
type OmpBot interface {
	Start(ctx context.Context)
}

type ompBot struct {
	token                           string
	cfg                             config.Bot
	equipmentRequestAPIClient       pb.BssEquipmentRequestApiServiceClient
	equipmentRequestFacadeAPIClient facadepb.BssEquipmentRequestFacadeApiServiceClient
}

// NewOmpBot used to create a new ompbot
func NewOmpBot(token string, cfg config.Bot, equipmentRequestAPIClient pb.BssEquipmentRequestApiServiceClient, equipmentRequestFacadeAPIClient facadepb.BssEquipmentRequestFacadeApiServiceClient) OmpBot {
	return &ompBot{
		token:                           token,
		cfg:                             cfg,
		equipmentRequestAPIClient:       equipmentRequestAPIClient,
		equipmentRequestFacadeAPIClient: equipmentRequestFacadeAPIClient,
	}
}

func (o *ompBot) Start(ctx context.Context) {
	bot, err := tgbotapi.NewBotAPI(o.token)
	if err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: tgbotapi.NewBotAPI failed ", ompBotLogTag),
			"err", err,
		)
	}

	bot.Debug = o.cfg.Debug

	logger.InfoKV(ctx, fmt.Sprintf("%s: authorized on account ", ompBotLogTag),
		"account", bot.Self.UserName,
	)

	u := tgbotapi.UpdateConfig{
		Timeout: o.cfg.Timeout,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: bot.GetUpdatesChan failed ", ompBotLogTag),
			"err", err,
		)
	}

	routerHandler := routerPkg.NewRouter(bot, o.cfg, o.equipmentRequestAPIClient, o.equipmentRequestFacadeAPIClient)

	for update := range updates {
		routerHandler.HandleUpdate(ctx, update)
	}
}
