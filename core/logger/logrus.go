package logger

import (
	"bytes"
	"fmt"
	"lightyear/core/global"

	"os"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	logger := global.Config.Logger

	// custom time format
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logger.Prefix, timestamp, levelColor, entry.Level.String(), fileVal, funcVal, entry.Message)

	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", logger.Prefix, timestamp, levelColor, entry.Level.String(), entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // ! Print the file name and line number of the log
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)

	global.Logger = logger

}

// ! Default Logger
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // ! Print the file name and line number of the log
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

}
