package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/metrics"
	"github.com/ozonmp/omp-bot/internal/ompBot"
	mwclient "github.com/ozonmp/omp-bot/internal/pkg/mw"
	"github.com/ozonmp/omp-bot/internal/server"
	"github.com/ozonmp/omp-bot/internal/tracer"
	"google.golang.org/grpc"
	"os"

	"github.com/joho/godotenv"
)

const ompBotMainLogTag = "OmpBotMain"

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed init configuration", ompBotMainLogTag),
			"err", err,
		)
	}
	cfg := config.GetConfigInstance()

	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed get Token", ompBotMainLogTag),
			"err", "environment variable TOKEN not found in .env",
		)
	}

	syncLogger := logger.NewLogger(ctx, cfg)
	defer syncLogger()

	metrics.InitMetrics(cfg)

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	tracing, err := tracer.NewTracer(ctx, &cfg)

	if err != nil {
		return
	}
	defer tracing.Close()

	equipmentRequestAPIConn, err := grpc.DialContext(
		ctx,
		cfg.EquipmentRequestAPI.Address,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
			mwclient.AddAppInfoUnary,
			grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(cfg.Bot.ConnectAttempts)),
		),
	)
	if err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed to create equipmentRequestApi connection  ", ompBotMainLogTag),
			"err", err,
		)
	}

	equipmentRequestAPIClient := pb.NewBssEquipmentRequestApiServiceClient(equipmentRequestAPIConn)

	equipmentRequestFacadeAPIConn, err := grpc.DialContext(
		ctx,
		cfg.EquipmentRequestFacadeAPI.Address,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
			mwclient.AddAppInfoUnary,
			grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(cfg.Bot.ConnectAttempts)),
		),
	)
	if err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed to create equipmentRequestFacadeApi connection  ", ompBotMainLogTag),
			"err", err,
		)
	}

	equipmentRequestFacadeAPIClient := facadepb.NewBssEquipmentRequestFacadeApiServiceClient(equipmentRequestFacadeAPIConn)

	bot := ompBot.NewOmpBot(token, cfg.Bot, equipmentRequestAPIClient, equipmentRequestFacadeAPIClient)

	if err := server.NewOmpBotServer(bot).Start(ctx, &cfg); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: failed creating omp bot server", ompBotMainLogTag),
			"err", err,
		)

		return
	}
}
