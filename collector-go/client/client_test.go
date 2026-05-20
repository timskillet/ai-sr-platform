package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timskillet/ai-reliability-platform/collector-go/client"
)

func TestSend_PostsIncidentToBackend(t *testing.T) {
	var received client.Incident
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/incidents" {
			t.Errorf("path = %s, want /incidents", r.URL.Path)
		}
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type = %s, want application/json", ct)
		}
		json.NewDecoder(r.Body).Decode(&received)
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	c := client.New(server.URL)
	incident := client.Incident{
		Service:  "payments-api",
		Severity: "critical",
		Message:  "database connection timeout",
	}
	if err := c.Send(incident); err != nil {
		t.Fatalf("Send returned unexpected error: %v", err)
	}
	if received.Service != "payments-api" {
		t.Errorf("Service = %q, want payments-api", received.Service)
	}
	if received.Severity != "critical" {
		t.Errorf("Severity = %q, want critical", received.Severity)
	}
	if received.Message != "database connection timeout" {
		t.Errorf("Message = %q", received.Message)
	}
}

func TestSend_ReturnsErrorOnNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	c := client.New(server.URL)
	err := c.Send(client.Incident{Service: "svc", Severity: "critical", Message: "oops"})
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
}

func TestSend_ReturnsErrorOnNetworkFailure(t *testing.T) {
	c := client.New("http://127.0.0.1:1") //nothing listening on port 1
	err := c.Send(client.Incident{Service: "svc", Severity: "critical", Message: "oops"})
	if err == nil {
		t.Fatal("expected error for network failure, got nil")
	}
}
