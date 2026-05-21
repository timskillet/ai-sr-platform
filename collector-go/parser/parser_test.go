package parser_test

import (
	"testing"

	"github.com/timskillet/ai-reliability-platform/collector-go/parser"
)

func TestParseLine_ErrorLine(t *testing.T) {
	event, ok := parser.ParseLine("ERROR database connection timeout")
	if !ok {
		t.Fatal("expected ok=true for ERROR line")
	}
	if event.Level != "ERROR" {
		t.Errorf("Level = %q, want ERROR", event.Level)
	}
	if event.Message != "database connection timeout" {
		t.Errorf("Message = %q, want %q", event.Message, "database connection timeout")
	}
}

func TestParseLine_WarnLine(t *testing.T) {
	event, ok := parser.ParseLine("WARN slow query detected (2300ms)")
	if !ok {
		t.Fatal("expected ok=true for WARN line")
	}
	if event.Level != "WARN" {
		t.Errorf("Level = %q, want WARN", event.Level)
	}
	if event.Message != "slow query detected (2300ms)" {
		t.Errorf("Message = %q", event.Message)
	}
}

func TestParseLine_InfoLineSkipped(t *testing.T) {
	_, ok := parser.ParseLine("INFO application started")
	if ok {
		t.Fatal("expected ok=false for INFO line")
	}
}

func TestParseLine_EmptyLineSkipped(t *testing.T) {
	_, ok := parser.ParseLine("")
	if ok {
		t.Fatal("expected ok=false for empty line")
	}
}

func TestParseLine_LowercaseError(t *testing.T) {
	event, ok := parser.ParseLine("error something bad happened")
	if !ok {
		t.Fatal("expected ok=true for lowercase error")
	}
	if event.Level != "ERROR" {
		t.Errorf("Level = %q, want ERROR", event.Level)
	}
	if event.Message != "something bad happened" {
		t.Errorf("Message = %q", event.Message)
	}
}

func TestParseLine_LowercaseWarn(t *testing.T) {
	event, ok := parser.ParseLine("warn disk usage high")
	if !ok {
		t.Fatal("expected ok=true for lowercase warn")
	}
	if event.Level != "WARN" {
		t.Errorf("Level = %q, want WARN", event.Level)
	}
	if event.Message != "disk usage high" {
		t.Errorf("Message = %q, want %q", event.Message, "disk usage high")
	}
}

func TestParseLine_TabSeparated(t *testing.T) {
	event, ok := parser.ParseLine("ERROR\tdatabase connection timeout")
	if !ok {
		t.Fatal("expected ok=true for tab-separated ERROR line")
	}
	if event.Level != "ERROR" {
		t.Errorf("Level = %q, want ERROR", event.Level)
	}
	if event.Message != "database connection timeout" {
		t.Errorf("Message = %q, want %q", event.Message, "database connection timeout")
	}
}
