package logger

import (
	"fmt"
	"os"

	"github.com/Xarasho/go-boilerplate/internal/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// LoggerService manages. New Relic integration and logger creation
type LoggerService struct {
	nrApp *newrelic.Application
}

// NewLoggerService creates a new logger service with New Relic integration
func NewLoggerService(cfg *config.ObservabilityConfig) *LoggerService {
	service := &LoggerService{}

	if cfg.NewRelic.LicenseKey == "" {
		fmt.Println("New Relic license key not provided, skipping initialization")
		return service
	}

	var configOptions []newrelic.ConfigOption
	configOptions = append(configOptions,
		newrelic.ConfigAppName(cfg.ServiceName),
		newrelic.ConfigLicense(cfg.NewRelic.LicenseKey),
		newrelic.ConfigAppLogForwardingEnabled(cfg.NewRelic.AppLogForwardingEnabled),
		newrelic.ConfigDistributedTracerEnabled(cfg.NewRelic.DistributedTracingEnabled),
	)

	// Add debug logging only if explicitly enabled
	if cfg.NewRelic.DebugLogging {
		configOptions = append(configOptions, newrelic.ConfigDebugLogger(os.Stdout))
	}

	app, err := newrelic.NewApplication(configOptions...)
	if err != nil {
		fmt.Printf("Failed to initialize New Relic: %v\n", err)
	}

	service.nrApp = app
	fmt.Printf("New Relic initialized for app: %s\n", cfg.ServiceName)
	return service
}
