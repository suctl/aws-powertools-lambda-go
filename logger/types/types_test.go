package types

import "testing"

func TestLogConfig(t *testing.T) {
	logConfig := LogConfig{
		IncludeContext: true,
		Properties: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}
	if logConfig.IncludeContext != true {
		t.Errorf("expected IncludeContext to be true, got %v", logConfig.IncludeContext)
	}
	if len(logConfig.Properties) != 2 {
		t.Errorf("expected Properties to have 2 entries, got %d", len(logConfig.Properties))
	}
	if logConfig.Properties["key1"] != "value1" {
		t.Errorf("expected Properties[\"key1\"] to be \"value1\", got \"%s\"", logConfig.Properties["key1"])
	}
	if logConfig.Properties["key2"] != "value2" {
		t.Errorf("expected Properties[\"key2\"] to be \"value2\", got \"%s\"", logConfig.Properties["key2"])
	}
}
