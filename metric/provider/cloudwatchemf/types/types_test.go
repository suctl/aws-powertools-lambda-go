package types

import (
	"encoding/json"
	"testing"
)

func TestCloudWatchEMFMetric(t *testing.T) {
	cloudWatchEmfMetric := CloudWatchEMFMetric{
		Name: "TestMetric",
		Unit: "Count",
	}
	expected := `{"Name":"TestMetric","Unit":"Count"}`
	actual, error := json.Marshal(cloudWatchEmfMetric)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchEMFMetricWithStorageResolution(t *testing.T) {
	cloudWatchEmfMetric := CloudWatchEMFMetric{
		Name:              "TestMetric",
		Unit:              "Count",
		StorageResolution: 60,
	}
	expected := `{"Name":"TestMetric","Unit":"Count","StorageResolution":60}`
	actual, error := json.Marshal(cloudWatchEmfMetric)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchMetrics(t *testing.T) {
	cloudWatchMetrics := CloudWatchMetrics{
		Namespace: "TestNamespace",
		Dimensions: [][]string{
			{"Dimension1"},
		},
		Metrics: []CloudWatchEMFMetric{
			{
				Name: "TestMetric1",
				Unit: "Count",
			},
		},
	}
	expected := `{"Namespace":"TestNamespace","Dimensions":[["Dimension1"]],"Metrics":[{"Name":"TestMetric1","Unit":"Count"}]}`
	actual, error := json.Marshal(cloudWatchMetrics)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchMetricsWithMultipleMetrics(t *testing.T) {
	cloudWatchMetrics := CloudWatchMetrics{
		Namespace: "TestNamespace",
		Dimensions: [][]string{
			{"Dimension1"},
		},
		Metrics: []CloudWatchEMFMetric{
			{
				Name: "TestMetric1",
				Unit: "Count",
			},
			{
				Name: "TestMetric2",
				Unit: "Milliseconds",
			},
		},
	}
	expected := `{"Namespace":"TestNamespace","Dimensions":[["Dimension1"]],"Metrics":[{"Name":"TestMetric1","Unit":"Count"},{"Name":"TestMetric2","Unit":"Milliseconds"}]}`
	actual, error := json.Marshal(cloudWatchMetrics)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchEMFRoot(t *testing.T) {
	cloudWatchEmfRoot := CloudWatchEMFRoot{
		Timestamp: 1625247600,
		CloudWatchMetrics: []CloudWatchMetrics{
			{
				Namespace: "TestNamespace",
				Dimensions: [][]string{
					{"Dimension1"},
				},
				Metrics: []CloudWatchEMFMetric{
					{
						Name: "TestMetric",
						Unit: "Count",
					},
				},
			},
		},
	}
	expected := `{"Timestamp":1625247600,"CloudWatchMetrics":[{"Namespace":"TestNamespace","Dimensions":[["Dimension1"]],"Metrics":[{"Name":"TestMetric","Unit":"Count"}]}]}`
	actual, error := json.Marshal(cloudWatchEmfRoot)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchEMFOutput(t *testing.T) {
	cloudWatchEmfOutput := CloudWatchEMFOutput{
		Aws: CloudWatchEMFRoot{
			Timestamp: 1625247600,
			CloudWatchMetrics: []CloudWatchMetrics{
				{
					Namespace: "TestNamespace",
					Dimensions: [][]string{
						{"Dimension1"},
					},
					Metrics: []CloudWatchEMFMetric{
						{
							Name: "TestMetric",
							Unit: "Count",
						},
					},
				},
			},
		},
	}
	expected := `{"_aws":{"Timestamp":1625247600,"CloudWatchMetrics":[{"Namespace":"TestNamespace","Dimensions":[["Dimension1"]],"Metrics":[{"Name":"TestMetric","Unit":"Count"}]}]}}`
	actual, error := json.Marshal(cloudWatchEmfOutput)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}
