package log

import "syscall/js"

type JSLogger struct {
	console js.Value
}

func NewJSLogger() *JSLogger {
	return &JSLogger{
		console: js.Global().Get("console"),
	}
}

func (l *JSLogger) Info(msg string, args ...any) {
	l.console.Call("log", append([]any{msg}, args...))
}

func (l *JSLogger) Error(msg string, args ...any) {
	l.console.Call("error", append([]any{msg}, args...))
}
