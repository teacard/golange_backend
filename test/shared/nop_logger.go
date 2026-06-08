package shared

import "io"

// NopLogger 實作 logger.Logger interface 的空操作版本，供測試使用
type NopLogger struct{}

func (n *NopLogger) Info(args ...interface{})                  {}
func (n *NopLogger) Infof(format string, args ...interface{})  {}
func (n *NopLogger) Warn(args ...interface{})                  {}
func (n *NopLogger) Warnf(format string, args ...interface{})  {}
func (n *NopLogger) Error(args ...interface{})                 {}
func (n *NopLogger) Errorf(format string, args ...interface{}) {}
func (n *NopLogger) Fatal(args ...interface{})                 {}
func (n *NopLogger) Fatalf(format string, args ...interface{}) {}
func (n *NopLogger) SetOutput(w io.Writer)                     {}
