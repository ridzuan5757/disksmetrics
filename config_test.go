package disksusagemetricsreceiver

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ridzuan5757/disksusagemetrics/internal/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(
		filepath.Join("testdata", "config.yaml"),
	)
	require.NoError(t, err)

	overrideMetricsBuilderConfig := metadata.DefaultMetricsBuilderConfig()

	tests := []struct {
		id       component.ID
		expected component.Config
	}{
		{
			id:       component.NewIDWithName(metadata.Type, "defaults"),
			expected: createDefaultConfig(),
		},
		{
			id: component.NewIDWithName(metadata.Type, ""),
			expected: &Config{
				ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
					CollectionInterval: time.Minute,
					InitialDelay:       time.Second,
				},
				Interval:             "1s",
				MetricsBuilderConfig: overrideMetricsBuilderConfig,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			assert.NoError(t, component.ValidateConfig(cfg))
			diff := cmp.Diff(
				tt.expected,
				cfg,
				cmpopts.IgnoreUnexported(metadata.MetricsBuilderConfig{}),
				cmpopts.IgnoreUnexported(metadata.MetricConfig{}),
			)
			if diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
