package collector

import (
	"fmt"
	"log"

	"github.com/nxadm/tail"
	"github.com/timskillet/ai-reliability-platform/collector-go/client"
	"github.com/timskillet/ai-reliability-platform/collector-go/parser"
)

var severityMap = map[string]string{
	"ERROR": "critical",
	"WARN":  "warning",
}

type Sender interface {
	Send(incident client.Incident) error
}

type Collector struct {
	logFile     string
	serviceName string
	sender      Sender
}

func New(logFile, serviceName string, sender Sender) *Collector {
	return &Collector{
		logFile:     logFile,
		serviceName: serviceName,
		sender:      sender,
	}
}

func (col *Collector) Run() error {
	t, err := tail.TailFile(col.logFile, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		return fmt.Errorf("tail file: %w", err)
	}
	for line := range t.Lines {
		if line.Err != nil {
			log.Printf("tail error: %v", line.Err)
			continue
		}
		col.processLine(line.Text)
	}
	return nil
}

func (col *Collector) processLine(text string) {
	event, ok := parser.ParseLine(text)
	if !ok {
		return
	}
	incident := client.Incident{
		Service:  col.serviceName,
		Severity: severityMap[event.Level],
		Message:  event.Message,
	}
	if err := col.sender.Send(incident); err != nil {
		log.Printf("send incident: %v", err)
	}
}
