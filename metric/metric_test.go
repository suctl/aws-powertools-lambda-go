package metric

import "testing"

func TestAddMetric(t *testing.T) {
	metric := New()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic occurred")
		}
	}()
	metric.AddMetric("TestMetric", "Count", 1.0, 60)
}
