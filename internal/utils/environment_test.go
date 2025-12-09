package utils

import (
	"os"
	"testing"
)

func TestGetEnvironmentVariableWhenKeyPresent(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")
	value := GetEnvironmentVariable("TEST_KEY", "default_value")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

}

func TestGetEnvironmentVariableWhenKeyNotPresent(t *testing.T) {
	value := GetEnvironmentVariable("TEST_KEY", "DEFAULT_VALUE")
	if value != "DEFAULT_VALUE" {
		t.Errorf("Expected 'DEFAULT_VALUE', got '%s'", value)
	}
}
