package log

import "github.com/charmbracelet/log"

type Logger interface {
	Log(string)
	Error(error, string)
}

type StdLogger struct{}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) Log(msg string) {
	log.Info(msg)
}

func (l *StdLogger) Error(err error, msg string) {
}
