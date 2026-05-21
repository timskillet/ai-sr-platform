package parser

import (
	"strings"
	"unicode"
)

type LogEvent struct {
	Level   string
	Message string
}

func ParseLine(line string) (LogEvent, bool) {
	idx := strings.IndexFunc(line, unicode.IsSpace)
	if idx < 0 {
		return LogEvent{}, false
	}
	level := strings.ToUpper(line[:idx])
	if level != "ERROR" && level != "WARN" {
		return LogEvent{}, false
	}
	message := strings.TrimLeftFunc(line[idx:], unicode.IsSpace)
	return LogEvent{Level: level, Message: message}, true
}
