/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package log

import "github.com/Sirupsen/logrus"

// DefaultLogger is the Default Logger
var DefaultLogger Logger

func init() {
	// Create default logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	DefaultLogger = logger
}

// Print calls the same method of DefaultLogger
func Print(args ...interface{}) {
	DefaultLogger.Print(args...)
}

// Printf calls the same method of DefaultLogger
func Printf(format string, args ...interface{}) {
	DefaultLogger.Printf(format, args...)
}

// Println calls the same method of DefaultLogger
func Println(args ...interface{}) {
	DefaultLogger.Println(args...)
}

// Fatal calls the same method of DefaultLogger
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Fatalf calls the same method of DefaultLogger
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
}

// Fatalln calls the same method of DefaultLogger
func Fatalln(args ...interface{}) {
	DefaultLogger.Fatalln(args...)
}

// Panic calls the same method of DefaultLogger
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// Panicf calls the same method of DefaultLogger
func Panicf(format string, args ...interface{}) {
	DefaultLogger.Panicf(format, args...)
}

// Panicln calls the same method of DefaultLogger
func Panicln(args ...interface{}) {
	DefaultLogger.Panicln(args...)
}

// Debug calls the same method of DefaultLogger
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Debugf calls the same method of DefaultLogger
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// Debugln calls the same method of DefaultLogger
func Debugln(args ...interface{}) {
	DefaultLogger.Debugln(args...)
}

// Error calls the same method of DefaultLogger
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Errorf calls the same method of DefaultLogger
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

// Errorln calls the same method of DefaultLogger
func Errorln(args ...interface{}) {
	DefaultLogger.Errorln(args...)
}

// Info calls the same method of DefaultLogger
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Infof calls the same method of DefaultLogger
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Infoln calls the same method of DefaultLogger
func Infoln(args ...interface{}) {
	DefaultLogger.Infoln(args...)
}

// Warn calls the same method of DefaultLogger
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Warnf calls the same method of DefaultLogger
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

// Warnln calls the same method of DefaultLogger
func Warnln(args ...interface{}) {
	DefaultLogger.Warnln(args...)
}