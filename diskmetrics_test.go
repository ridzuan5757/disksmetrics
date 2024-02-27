package disksusagemetricsreceiver

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHddMetrics(t *testing.T) {
	metrics, err := GetHddMetrics()
	assert.NoError(t, err, "GetHddMetrics returned an error")
	assert.NotZero(t, metrics.Total, "Total disk space is 0")
	assert.GreaterOrEqual(t, metrics.UsedPercent, 0.0, "Value has to be greater than 0")
	assert.LessOrEqual(t, metrics.UsedPercent, 1.0, "Value has to be less than 1")
}

func TestGetHddMetricsConcurrentSafety(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := GetHddMetrics()
			assert.NoError(t, err, "GetHddMetrics returned an error")
		}()
	}
	wg.Wait()
}

func BenchmarkGetHddMetrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetHddMetrics()
		if err != nil {
			b.Errorf("Error fetching HDD metrics: %v", err)
		}
	}
}

func TestGetMountMetrics(t *testing.T) {
	_, err := GetMountMetrics()
	assert.NoError(t, err, "GetHddMetrics returned an error")
}

func BenchmarkGetMountMetrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetMountMetrics()
		if err != nil {
			b.Errorf("Error fetching HDD metrics: %v", err)
		}
	}
}

func TestGetMountMetricsConcurentSafety(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := GetMountMetrics()
			assert.NoError(t, err, "GetMountMetrics returned an error")
		}()
	}
	wg.Wait()

}
