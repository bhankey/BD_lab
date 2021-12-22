package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime/debug"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func (l *Logger) Error(args ...interface{}) {
	l.Entry.WithFields(logrus.Fields{
		"stack": string(debug.Stack())},
	).Error(args)
}

func (l *Logger) WithFields(fields logrus.Fields) *Logger {
	return &Logger{l.Entry.WithFields(fields)}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Entry.WithField(key, value)}
}

const LogLevel = logrus.DebugLevel

// init initialize logger
func Init() {
	l := logrus.New()

	//l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	l.SetLevel(LogLevel)

	err := os.MkdirAll("logs", 0755)
	if err != nil || os.IsExist(err) {
		l.Fatal("can't create log directory:", err)
	}

	logFile, err := os.OpenFile("logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		l.Fatal("can't create log file:", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	l.SetOutput(mw)
	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}

// GetLogger return standard logger, which was initialize by func init
func GetLogger() Logger {
	return Logger{e}
}

//GetLoggerWithField return logger with fields, which was initialize by func init
func GetLoggerWithField(f logrus.Fields) Logger {
	return Logger{e.WithFields(f)}
}
