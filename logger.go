package logxgo

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log   *logrus.Logger
	entry *logrus.Entry
}

func New(options ...Option) ILogger {
	config := Config{
		FileLogging: true,
	}

	for _, option := range options {
		option(&config)
	}

	log := logrus.New()

	log.SetOutput(os.Stdout)

	if config.JSON {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		log.SetFormatter(&TerminalFormatter{})
	}

	if config.FileLogging {
		_ = os.MkdirAll("./var/log", 0755)

		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  "./var/log/info_access.log",
			logrus.ErrorLevel: "./var/log/error_access.log",
			logrus.DebugLevel: "./var/log/debug_access.log",
			logrus.WarnLevel:  "./var/log/warning_access.log",
			logrus.FatalLevel: "./var/log/fatal_access.log",
			logrus.PanicLevel: "./var/log/panic_access.log",
			logrus.TraceLevel: "./var/log/trace_access.log",
		}

		log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
			},
		))
	}

	setLogLevel(log, config.Level)

	entry := logrus.NewEntry(log)

	if config.Module != "" {
		entry = entry.WithField("module", config.Module)
	}

	return &Logger{
		log:   log,
		entry: entry,
	}
}

func NewLogger(options ...Option) ILogger {
	return New(options...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.with(fields...).Info(msg)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.with(fields...).Error(msg)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.with(fields...).Warn(msg)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.with(fields...).Debug(msg)
}

func (l *Logger) Trace(msg string, fields ...Field) {
	l.with(fields...).Trace(msg)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.with(fields...).Panic(msg)
}

func (l *Logger) WithFields(fields ...Field) ILogger {
	return &Logger{
		log:   l.log,
		entry: l.with(fields...),
	}
}

func (l *Logger) SetLevel(level string) {
	setLogLevel(l.log, level)
}

func (l *Logger) with(fields ...Field) *logrus.Entry {
	if len(fields) == 0 {
		return l.entry
	}

	logFields := logrus.Fields{}

	for _, field := range fields {
		logFields[field.Key] = field.Value
	}

	return l.entry.WithFields(logFields)
}

func setLogLevel(log *logrus.Logger, level string) {
	switch strings.ToLower(level) {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "warn", "warning":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}
