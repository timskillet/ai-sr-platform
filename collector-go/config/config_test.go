package config_test

import (
	"os"
	"testing"

	"github.com/timskillet/ai-reliability-platform/collector-go/config"
)

func TestLoad_MissingLogFile(t *testing.T) {
	os.Unsetenv("LOG_FILE")
	os.Unsetenv("SERVICE_NAME")
	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error when LOG_FILE is missing")
	}
}

func TestLoad_MissingServiceName(t *testing.T) {
	os.Setenv("LOG_FILE", "/tmp/app.log")
	os.Unsetenv("SERVICE_NAME")
	defer os.Unsetenv("LOG_FILE")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error when SERVICE_NAME is missing")
	}
}

func TestLoad_DefaultsBackendURL(t *testing.T) {
	os.Setenv("LOG_FILE", "tmp/app.log")
	os.Setenv("SERVICE_NAME", "test-svc")
	os.Unsetenv("BACKEND_URL")
	defer os.Unsetenv("LOG_FILE")
	defer os.Unsetenv("SERVICE_NAME")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.BackendURL != "http://localhost:8080" {
		t.Errorf("BackendURL = %q, want http://localhost:8080", cfg.BackendURL)
	}
}

func TestLoad_ReadsAllVars(t *testing.T) {
	os.Setenv("LOG_FILE", "/var/log/app.log")
	os.Setenv("SERVICE_NAME", "payments-api")
	os.Setenv("BACKEND_URL", "http://backend:9090")
	defer os.Unsetenv("LOG_FILE")
	defer os.Unsetenv("SERVICE_NAME")
	defer os.Unsetenv("BACKEND_URL")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.LogFile != "/var/log/app.log" {
		t.Errorf("LogFile = %q", cfg.LogFile)
	}
	if cfg.ServiceName != "payments-api" {
		t.Errorf("ServiceName = %q", cfg.ServiceName)
	}
	if cfg.BackendURL != "http://backend:9090" {
		t.Errorf("BackendURL = %q", cfg.BackendURL)
	}
}
