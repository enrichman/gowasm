package log

import (
	"github.com/charmbracelet/log"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type StdLogger struct{}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) Info(msg string, args ...any) {
	log.Info(msg)
}

func (l *StdLogger) Error(msg string, args ...any) {
	log.Error(msg)
}
