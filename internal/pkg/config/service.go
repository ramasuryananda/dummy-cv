package config

import (
	"os"
	"strconv"
)

// Get retrieves the value of the environment variable identified by the given key.
// If the environment variable is not found, an empty string is returned.
func Get(key string) string {
	return os.Getenv(key)
}

// GetWithDefault retrieves the value of the environment variable identified by the given key.
// If the environment variable is not found or its value is an empty string, the defaultValue is returned.
func GetWIthDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) int {
	value := Get(key)
	if value == "" {
		return defaultValue
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return res
}
