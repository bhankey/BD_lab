package logger

import (
	"io"
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// nolint: gochecknoglobals
var e *logrus.Entry // TODO Mustn't be singleton

type Logger struct {
	*logrus.Entry
}

// nolint: govet
func (l *Logger) Error(args ...interface{}) {
	l.Entry.WithFields(logrus.Fields{
		"stack": string(debug.Stack()),
	},
	).Error(args)
}

func (l *Logger) WithFields(fields logrus.Fields) *Logger {
	return &Logger{l.Entry.WithFields(fields)}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Entry.WithField(key, value)}
}

const LogLevel = logrus.DebugLevel

const loggerFileSystemRights = os.FileMode(0o755)

// Init initialize logger.
func Init() {
	log := logrus.New()

	// l.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.SetLevel(LogLevel)

	if err := os.MkdirAll("logs", loggerFileSystemRights); err != nil && !os.IsExist(err) {
		log.Fatal("can't create log directory:", err)
	}

	logFile, err := os.OpenFile("logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, loggerFileSystemRights)
	if err != nil {
		log.Fatal("can't create log file:", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(log)
}

// GetLogger return standard logger, which was initialize by func init.
func GetLogger() Logger {
	return Logger{e}
}

// GetLoggerWithField return logger with fields, which was initialize by func init.
func GetLoggerWithField(f logrus.Fields) Logger {
	return Logger{e.WithFields(f)}
}
