package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Incident struct {
	Service  string `json:"service"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type IncidentClient struct {
	backendURL string
	httpClient *http.Client
}

func New(backendURL string) *IncidentClient {
	return &IncidentClient{
		backendURL: backendURL,
		httpClient: &http.Client{},
	}
}

func (c *IncidentClient) Send(incident Incident) error {
	body, err := json.Marshal(incident)
	if err != nil {
		return fmt.Errorf("marshal incident: %w", err)
	}

	resp, err := c.httpClient.Post(c.backendURL+"/incidents", "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("post incident: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d from backend", resp.StatusCode)
	}

	return nil
}
