package disksusagemetricsreceiver

import (
	"context"

	"github.com/ridzuan5757/disksusagemetrics/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

func newMetricReceiver(
	_ context.Context,
	set receiver.CreateSettings,
	rCfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	cfg, _ := rCfg.(*Config)
	metricsBuilder := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, set)
	ns := newScraper(metricsBuilder, set.Logger, cfg)

	scraper, err := scraperhelper.NewScraper(
		metadata.Type.String(),
		ns.scrape,
	)

	if err != nil {
		ns.logger.Error("Failed to create scraper helper.")
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(
		&cfg.ScraperControllerSettings,
		set,
		consumer,
		scraperhelper.AddScraper(scraper),
	)
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(newMetricReceiver, metadata.MetricsStability),
	)
}
