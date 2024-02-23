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

func (l *JSLogger) Log(msg string) {
	l.console.Call("log", msg)
}

func (l *JSLogger) Error(err error, msg string) {
	l.console.Call("error", msg, err.Error())
}
