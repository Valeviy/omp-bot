package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/ompBot"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

const ompBotServerStartLogTag = "OmpBotServer.Start()"

// OmpBotServer is omp bot server
type OmpBotServer struct {
	bot ompBot.OmpBot
}

// NewOmpBotServer returns facade
func NewOmpBotServer(bot ompBot.OmpBot) *OmpBotServer {
	return &OmpBotServer{
		bot: bot,
	}
}

// Start method runs server
func (o *OmpBotServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	metricsServer := createMetricsServer(cfg)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("%s: metrics server is running on", ompBotServerStartLogTag),
			"address", metricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.ListenAndServe failed", ompBotServerStartLogTag),
				"err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(ctx, cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, fmt.Sprintf("%s: status server is running on", ompBotServerStartLogTag),
			"address", statusAdrr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.ListenAndServe failed", ompBotServerStartLogTag),
				"err", err)
		}
	}()

	o.bot.Start(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.Info(ctx, fmt.Sprintf("%s: the service is ready", ompBotServerStartLogTag))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, fmt.Sprintf("%s: signal.Notify", ompBotServerStartLogTag),
			"quit", v,
		)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, fmt.Sprintf("%s: ctx.Done", ompBotServerStartLogTag),
			"done", done,
		)
	}

	isReady.Store(false)

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.Shutdown failed", ompBotServerStartLogTag),
			"err", err,
		)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: statusServer shut down correctly", ompBotServerStartLogTag))
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.Shutdown failed", ompBotServerStartLogTag),
			"err", err,
		)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: metricsServer shut down correctly", ompBotServerStartLogTag))
	}

	return nil
}
