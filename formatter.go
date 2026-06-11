package logxgo

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var startTime = time.Now()

type TerminalFormatter struct{}

func (f *TerminalFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var buffer bytes.Buffer

	elapsed := int(time.Since(startTime).Seconds())

	level := strings.ToUpper(entry.Level.String())

	color := getLevelColor(entry.Level)

	buffer.WriteString(
		fmt.Sprintf(
			"%s%s[%04d]\033[0m %s\n",
			color,
			level,
			elapsed,
			entry.Message,
		),
	)

	return buffer.Bytes(), nil
}

func getLevelColor(level logrus.Level) string {

	switch level {

	case logrus.TraceLevel:
		return "\033[37m" // gris

	case logrus.DebugLevel:
		return "\033[36m" // cyan

	case logrus.InfoLevel:
		return "\033[34m" // azul

	case logrus.WarnLevel:
		return "\033[33m" // amarillo

	case logrus.ErrorLevel:
		return "\033[31m" // rojo

	case logrus.PanicLevel:
		return "\033[35m" // magenta

	case logrus.FatalLevel:
		return "\033[35m" // magenta

	default:
		return "\033[0m"
	}
}