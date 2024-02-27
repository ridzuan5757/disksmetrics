package disksusagemetricsreceiver

import (
	"fmt"
	"time"

	"github.com/ridzuan5757/disksusagemetrics/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type Config struct {
	Verbose                                 bool   `mapstructure:"verbose"`
	Interval                                string `mapstructure:"interval"`
	metadata.MetricsBuilderConfig           `mapstructure:",squash"`
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
}

func createDefaultConfig() component.Config {
	return &Config{
		Interval:                  fmt.Sprint(1 * time.Second),
		ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(metadata.Type),
		MetricsBuilderConfig:      metadata.DefaultMetricsBuilderConfig(),
	}
}

func (cfg *Config) Validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Seconds() < 1 {
		return fmt.Errorf("when defined, the interval has to be at least 1s")
	}
	return nil

}
