package main

import (
	"log"

	"github.com/timskillet/ai-reliability-platform/collector-go/client"
	"github.com/timskillet/ai-reliability-platform/collector-go/collector"
	"github.com/timskillet/ai-reliability-platform/collector-go/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	c := client.New(cfg.BackendURL)
	col := collector.New(cfg.LogFile, cfg.ServiceName, c)
	if err := col.Run(); err != nil {
		log.Fatalf("collector: %v", err)
	}
}
