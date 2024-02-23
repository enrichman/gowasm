package log

type Logger interface {
	Log(string)
	Error(error, string)
}
