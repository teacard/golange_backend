package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Logger interface — 測試時可注入 mock 取代真實 logrus
type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	SetOutput(w io.Writer)
}

// logrusLogger 是 Logger interface 的 logrus 實作
type logrusLogger struct {
	log *logrus.Logger
}

// New 建立一個新的 logrus logger
func New() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})
	return &logrusLogger{log: l}
}

func (l *logrusLogger) Info(args ...interface{})                 { l.log.Info(args...) }
func (l *logrusLogger) Infof(f string, args ...interface{})      { l.log.Infof(f, args...) }
func (l *logrusLogger) Warn(args ...interface{})                 { l.log.Warn(args...) }
func (l *logrusLogger) Warnf(f string, args ...interface{})      { l.log.Warnf(f, args...) }
func (l *logrusLogger) Error(args ...interface{})                { l.log.Error(args...) }
func (l *logrusLogger) Errorf(f string, args ...interface{})     { l.log.Errorf(f, args...) }
func (l *logrusLogger) Fatal(args ...interface{})                { l.log.Fatal(args...) }
func (l *logrusLogger) Fatalf(f string, args ...interface{})     { l.log.Fatalf(f, args...) }
func (l *logrusLogger) SetOutput(w io.Writer)                    { l.log.SetOutput(w) }
