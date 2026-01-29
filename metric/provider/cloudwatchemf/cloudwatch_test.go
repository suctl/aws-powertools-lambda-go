package cloudwatchemf

import (
	"testing"
)

func TestAddMetric(t *testing.T) {
	cw := New(CloudWatchEMFConfig{
		Namespace: "Test Namespace",
		Dimension: "Test Dimention",
	})
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic occurred")
		}
	}()
	cw.AddMetric("TestMetric", "Count", 1.0, 60)
}

func TestAddMetricWithInvalidMetricName(t *testing.T) {
	cw := New(CloudWatchEMFConfig{
		Namespace: "Test Namespace",
		Dimension: "Test Dimention",
	})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	cw.AddMetric("", "Count", 1.0, 60)
}

func TestPublishAndFlushMetric(t *testing.T) {
	cw := New(CloudWatchEMFConfig{
		Namespace: "Test Namespace",
		Dimension: "Test Dimention",
	})
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic occurred")
		}
	}()
	cw.AddMetric("TestMetric1", "Count", 1.0, 60)
	cw.AddMetric("TestMetric2", "Percent", 50.0, 60)
	cw.flushMetric()
	if len(cw.aws.CloudWatchMetrics[0].Metrics) != 0 {
		t.Errorf("expected metrics to be flushed, but found %d metrics", len(cw.aws.CloudWatchMetrics[0].Metrics))
	}
}

func TestPublishMetric(t *testing.T) {
	cw := New(CloudWatchEMFConfig{
		Namespace: "Test Namespace",
		Dimension: "Test Dimention",
	})
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic occurred")
		}
	}()
	cw.AddMetric("TestMetric", "Count", 1.0, 60)
	publishMetric(cw.serilizedMetricOutput())
}

func TestGenerateCloudWatchEMF(t *testing.T) {
	cw := New(CloudWatchEMFConfig{
		Namespace: "Test Namespace",
		Dimension: "Test Dimention",
	})
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic occurred")
		}
	}()
	cw.AddMetric("TestMetric", "Count", 1.0, 60)
	cw.flushMetric()
}
