package logxgo

type ILogger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Trace(msg string, fields ...Field)
	Panic(msg string, fields ...Field)

	WithFields(fields ...Field) ILogger
	SetLevel(level string)
}
