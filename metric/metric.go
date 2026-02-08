package metric

import (
	"github.com/suctl/aws-powertools-lambda-go/internal/utils"
	"github.com/suctl/aws-powertools-lambda-go/metric/provider/cloudwatchemf"
)

type MetricInterface interface {
	AddMetric(name string, unit string, value float64, storageResolution int)
	LogMetrics()
}

func New() MetricInterface {
	return cloudwatchemf.New(cloudwatchemf.CloudWatchEMFConfig{})
}

func isMetricsDisabled() bool {
	return utils.GetEnvironmentVariable("POWERTOOLS_METRICS_DISABLED", "false") == "true"
}
