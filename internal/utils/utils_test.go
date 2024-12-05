package utils

import (
	"os"
	"testing"
)

func TestGetFromEnv(t *testing.T) {
	t.Run("should return environment variable value when it exists", func(t *testing.T) {
		key := "TEST_KEY"
		value := "test_value"
		err := os.Setenv(key, value)

		if err != nil {
			t.Errorf("error setting environment variable: %s", key)
			return
		}

		defer func(key string) {
			err := os.Unsetenv(key)
			if err != nil {
				t.Errorf("error unsetting environment variable: %s", key)
				return
			}
		}(key)

		result := GetFromEnv(key, "default_value")
		if result != value {
			t.Errorf("expected %s, got %s", value, result)
		}
	})

	t.Run("should return default value when environment variable does not exist", func(t *testing.T) {
		key := "NON_EXISTENT_KEY"
		defaultValue := "default_value"

		result := GetFromEnv(key, defaultValue)
		if result != defaultValue {
			t.Errorf("expected %s, got %s", defaultValue, result)
		}
	})
}
