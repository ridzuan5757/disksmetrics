package disksusagemetricsreceiver

import (
	"context"
	"testing"

	"github.com/ridzuan5757/disksusagemetrics/internal/metadata"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestNewScraper(t *testing.T) {
	logger := zap.NewNop()
	config := createDefaultConfig().(*Config)
	metricsBuilder := &metadata.MetricsBuilder{}

	s := newScraper(metricsBuilder, logger, config)

	assert.NotNil(t, s)
	assert.Equal(t, s.config, config)
	assert.Equal(t, s.logger, logger)
	assert.Equal(t, s.mb, metricsBuilder)
}

func TestScrape(t *testing.T) {
	logger := zaptest.NewLogger(t)
	metricsBuilder := metadata.NewMetricsBuilder(
		metadata.DefaultMetricsBuilderConfig(),
		receivertest.NewNopCreateSettings(),
	)
	config := createDefaultConfig().(*Config)
	s := newScraper(metricsBuilder, logger, config)

	ctx := context.Background()
	metrics, err := s.scrape(ctx)

	assert.NotNil(t, metrics)
	assert.Nil(t, err)

	assert.Equal(t, 4, metrics.MetricCount())
}

func TestScrapeModeVerbose(t *testing.T) {
	logger := zaptest.NewLogger(t)
	metricsBuilder := metadata.NewMetricsBuilder(
		metadata.DefaultMetricsBuilderConfig(),
		receivertest.NewNopCreateSettings(),
	)
	config := createDefaultConfig().(*Config)
	config.Verbose = true
	s := newScraper(metricsBuilder, logger, config)

	ctx := context.Background()
	metrics, err := s.scrape(ctx)

	assert.NotNil(t, metrics)
	assert.Nil(t, err)

	assert.Equal(t, 8, metrics.MetricCount())
}
