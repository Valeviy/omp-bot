package metrics

import (
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalCommandCalls *prometheus.CounterVec

// InitMetrics - init facade service metrics
func InitMetrics(cfg config.Config) {
	totalCommandCalls = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: cfg.Project.ServiceName,
		Name:      "total_command_calls",
		Help:      "Total bot command calls",
	}, []string{"command"})
}

// IncTotalCommandCalls - increment amount of total bot command calls
func IncTotalCommandCalls(command string) {
	totalCommandCalls.WithLabelValues(command).Inc()
}
