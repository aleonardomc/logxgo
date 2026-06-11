package logxgo

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var startTime = time.Now()

type TerminalFormatter struct{}

func (f *TerminalFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var buffer bytes.Buffer

	elapsed := int(time.Since(startTime).Seconds())

	level := getLevelName(entry.Level)
	color := getLevelColor(entry.Level)

	buffer.WriteString(
		fmt.Sprintf(
			"%s%s\033[0m[%04d] %s\n",
			color,
			level,
			elapsed,
			entry.Message,
		),
	)

	return buffer.Bytes(), nil
}

func getLevelName(level logrus.Level) string {

	switch level {

	case logrus.TraceLevel:
		return "TRAC"

	case logrus.DebugLevel:
		return "DEBU"

	case logrus.InfoLevel:
		return "INFO"

	case logrus.WarnLevel:
		return "WARN"

	case logrus.ErrorLevel:
		return "ERRO"

	case logrus.FatalLevel:
		return "FATA"

	case logrus.PanicLevel:
		return "PANI"

	default:
		return "INFO"
	}
}

func getLevelColor(level logrus.Level) string {

	switch level {

	case logrus.TraceLevel:
		return "\033[37m"

	case logrus.DebugLevel:
		return "\033[30m"

	case logrus.InfoLevel:
		return "\033[36m"

	case logrus.WarnLevel:
		return "\033[33m"

	case logrus.ErrorLevel:
		return "\033[31m"

	case logrus.FatalLevel:
		return "\033[35m"

	case logrus.PanicLevel:
		return "\033[35m"

	default:
		return "\033[0m"
	}
}