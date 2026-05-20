package parser

import "strings"

type LogEvent struct {
	Level   string
	Message string
}

func ParseLine(line string) (LogEvent, bool) {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) < 2 {
		return LogEvent{}, false
	}
	level := strings.ToUpper(parts[0])
	if level != "ERROR" && level != "WARN" {
		return LogEvent{}, false
	}
	return LogEvent{Level: level, Message: parts[1]}, true
}
