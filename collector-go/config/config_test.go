package config_test

import (
	"testing"

	"github.com/timskillet/ai-reliability-platform/collector-go/config"
)

func TestLoad_MissingLogFile(t *testing.T) {
	t.Setenv("LOG_FILE", "")
	t.Setenv("SERVICE_NAME", "")
	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error when LOG_FILE is missing")
	}
}

func TestLoad_MissingServiceName(t *testing.T) {
	t.Setenv("LOG_FILE", "/tmp/app.log")
	t.Setenv("SERVICE_NAME", "")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error when SERVICE_NAME is missing")
	}
}

func TestLoad_DefaultsBackendURL(t *testing.T) {
	t.Setenv("LOG_FILE", "/tmp/app.log")
	t.Setenv("SERVICE_NAME", "test-svc")
	t.Setenv("BACKEND_URL", "")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.BackendURL != "http://localhost:8080" {
		t.Errorf("BackendURL = %q, want http://localhost:8080", cfg.BackendURL)
	}
}

func TestLoad_ReadsAllVars(t *testing.T) {
	t.Setenv("LOG_FILE", "/var/log/app.log")
	t.Setenv("SERVICE_NAME", "payments-api")
	t.Setenv("BACKEND_URL", "http://backend:9090")

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

func TestLoad_WhitespaceOnlyLogFileRejected(t *testing.T) {
	t.Setenv("LOG_FILE", "   ")
	t.Setenv("SERVICE_NAME", "test-svc")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for whitespace-only LOG_FILE")
	}
}

func TestLoad_WhitespaceOnlyServiceNameRejected(t *testing.T) {
	t.Setenv("LOG_FILE", "/tmp/app.log")
	t.Setenv("SERVICE_NAME", "   ")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for whitespace-only SERVICE_NAME")
	}
}
