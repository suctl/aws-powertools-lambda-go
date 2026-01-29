package cloudwatchemf

import (
	"encoding/json"
	"fmt"

	"github.com/suctl/aws-powertools-lambda-go/internal/utils"
	"github.com/suctl/aws-powertools-lambda-go/metric/provider/cloudwatchemf/types"
)

const (
	maxDimensionsPerMetric = 29
	maxMetricesPerEMF      = 100
	maxMetricNameLength    = 255
	minMetricNameLength    = 1
)

type CloudWatchEMFConfig struct {
	Namespace string
	Dimension string
}

type cloudWatchEMF struct {
	aws          types.CloudWatchEMFRoot
	metricFields map[string][]float64
}

func New(config CloudWatchEMFConfig) *cloudWatchEMF {
	return &cloudWatchEMF{
		aws: types.CloudWatchEMFRoot{
			CloudWatchMetrics: []types.CloudWatchMetrics{
				{
					Namespace: config.Namespace,
					Dimensions: [][]string{
						{config.Dimension},
					},
					Metrics: []types.CloudWatchEMFMetric{},
				},
			},
		},
	}
}

func (cw *cloudWatchEMF) AddMetric(name string, unit string, value float64, storageResolution int) {
	if storageResolution != 1 && storageResolution != 60 {
		panic("storage resolution must be either 1 or 60")
	}
	if len(name) > maxMetricNameLength || len(name) < minMetricNameLength {
		panic(fmt.Sprintf("metric name length should be between %d and %d characters", minMetricNameLength, maxMetricNameLength))
	}
	metric := types.CloudWatchEMFMetric{
		Name:              name,
		Unit:              unit,
		StorageResolution: storageResolution,
	}
	cw.aws.CloudWatchMetrics[0].Metrics = append(cw.aws.CloudWatchMetrics[0].Metrics, metric)
	cw.addMetricValue(metric.Name, value)

	if len(cw.aws.CloudWatchMetrics[0].Metrics) == maxMetricesPerEMF {
		fmt.Println("Max metrics per EMF reached, publishing metrics")
		publishMetric(cw.serilizedMetricOutput())
		cw.flushMetric()
	}
}

func (cw *cloudWatchEMF) addMetricValue(name string, value float64) {
	if cw.metricFields == nil {
		cw.metricFields = make(map[string][]float64)
	}
	cw.metricFields[name] = append(cw.metricFields[name], value)
}

func (cw *cloudWatchEMF) LogMetrics() {
	metricLog := cw.serilizedMetricOutput()
	publishMetric(metricLog)
	cw.flushMetric()
}

func (cw *cloudWatchEMF) flushMetric() {
	cw.aws.CloudWatchMetrics[0].Metrics = []types.CloudWatchEMFMetric{}
}

func (cw *cloudWatchEMF) serilizedMetricOutput() []byte {
	cw.aws.Timestamp = utils.GetCurrentTimestamp()

	output := map[string]any{
		"_aws": cw.aws,
	}
	for key, value := range cw.metricFields {
		output[key] = value
	}
	metricLog, error := json.Marshal(output)
	if error != nil {
		panic("failed to serialize metrics")
	}
	return metricLog
}

func publishMetric(metricLog []byte) {
	fmt.Printf("%s", metricLog)
}
