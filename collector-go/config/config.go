package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	LogFile     string
	ServiceName string
	BackendURL  string
}

func Load() (Config, error) {
	logFile := strings.TrimSpace(os.Getenv("LOG_FILE"))
	if logFile == "" {
		return Config{}, fmt.Errorf("LOG_FILE environment variable is required")
	}
	serviceName := strings.TrimSpace(os.Getenv("SERVICE_NAME"))
	if serviceName == "" {
		return Config{}, fmt.Errorf("SERVICE_NAME environment variable is required")
	}
	backendURL := strings.TrimSpace(os.Getenv("BACKEND_URL"))
	if backendURL == "" {
		backendURL = "http://localhost:8080"
	}
	return Config{
		LogFile:     logFile,
		ServiceName: serviceName,
		BackendURL:  backendURL,
	}, nil
}
