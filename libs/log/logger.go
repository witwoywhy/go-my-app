package log

type Logger interface {
	Info(obj any)
	Infof(format string, obj ...any)
	Debug(obj any)
	Debugf(format string, obj any)
	Warn(obj any)
	Warnf(format string, obj any)
	Error(obj any)
	Errorf(format string, obj any)
	Fatal(obj any)
	Fatalf(format string, obj any)
	Panic(obj any)
	Panicf(format string, obj any)

	LogHttpRequest(msg string, headerObj, bodyObj any)
}
