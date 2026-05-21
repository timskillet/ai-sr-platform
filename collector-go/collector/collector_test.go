package collector

import (
	"errors"
	"testing"

	"github.com/timskillet/ai-reliability-platform/collector-go/client"
)

type mockSender struct {
	incidents []client.Incident
	err       error
}

func (m *mockSender) Send(i client.Incident) error {
	m.incidents = append(m.incidents, i)
	return m.err
}

func TestProcessLine_ErrorLine(t *testing.T) {
	sender := &mockSender{}
	col := New("/tmp/any.log", "payments-api", sender)

	col.processLine("ERROR database connection timeout")

	if len(sender.incidents) != 1 {
		t.Fatalf("expected 1 incident, got %d", len(sender.incidents))
	}
	got := sender.incidents[0]
	if got.Service != "payments-api" {
		t.Errorf("Service = %q, want payments-api", got.Service)
	}
	if got.Severity != "critical" {
		t.Errorf("Severity = %q, want critical", got.Severity)
	}
	if got.Message != "database connection timeout" {
		t.Errorf("Message = %q", got.Message)
	}
}

func TestProcessLine_SendErrorLogged(t *testing.T) {
	// Send returns an error - collector should not panic or propagate it
	sender := &mockSender{err: errors.New("backend down")}
	col := New("/tmp/any.log", "svc", sender)

	// Should not panic
	col.processLine("ERROR something failed")
}
