package client

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/pkg/auth"
	"os"
)

func stringPointerOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func getFirstNonEmptyConfig(configs ...config) config {
	for _, config := range configs {
		if config.Value != "" {
			return config
		}
	}
	return config{Source: "default"}
}

func getEnvVarConfig(key string) config {
	return config{
		os.Getenv(key),
		fmt.Sprintf("environment variable '%s'", key),
	}
}

func getChalkClientArgConfig(value string) config {
	return config{
		value,
		"ChalkClient argument",
	}
}

func getChalkYamlConfig(value string) config {
	var path string

	configPath, err := auth.GetConfigPath()
	if err != nil {
		path = "unknown"
	} else {
		path = *configPath
	}

	return config{
		value,
		fmt.Sprintf("config file %s", path),
	}
}
