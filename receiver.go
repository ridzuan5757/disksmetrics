package disksusagemetricsreceiver

import (
	"context"
	"time"

	"github.com/ridzuan5757/disksusagemetrics/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type scraper struct {
	mb     *metadata.MetricsBuilder
	logger *zap.Logger
	config *Config
}

func newScraper(
	metricsBuilder *metadata.MetricsBuilder,
	logger *zap.Logger,
	config *Config,
) *scraper {
	return &scraper{
		mb:     metricsBuilder,
		logger: logger,
		config: config,
	}
}

func (s *scraper) scrape(ctx context.Context) (pmetric.Metrics, error) {

	now := pcommon.NewTimestampFromTime(time.Now())
	diskUsage, _ := GetHddMetrics()
	s.mb.RecordDiskFreeDataPoint(now, int64(diskUsage.Free))
	s.mb.RecordDiskUsedDataPoint(now, int64(diskUsage.Used))
	s.mb.RecordDiskTotalDataPoint(now, int64(diskUsage.Total))
	s.mb.RecordDiskUtilizationDataPoint(now, diskUsage.UsedPercent)

	if s.config.Verbose {
		mountUsage, _ := GetMountMetrics()
		for _, mr := range mountUsage {
			s.mb.RecordMountFreeDataPoint(now, int64(mr.Free), mr.Mount)
			s.mb.RecordMountUsedDataPoint(now, int64(mr.Used), mr.Mount)
			s.mb.RecordMountTotalDataPoint(now, int64(mr.Total), mr.Mount)
			s.mb.RecordMountUtilizationDataPoint(now, mr.UsedPercent, mr.Mount)
		}
	}

	return s.mb.Emit(), nil
}
