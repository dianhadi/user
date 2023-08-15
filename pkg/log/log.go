package log

import (
	"fmt"
)

// Debug function
func Debug(args ...interface{}) {
	infoLogger.Debug().Timestamp().Msg(fmt.Sprint(args...))
}

// Debugln function
func Debugln(args ...interface{}) {
	infoLogger.Debug().Timestamp().Msg(fmt.Sprintln(args...))
}

// Debugf function
func Debugf(format string, v ...interface{}) {
	infoLogger.Debug().Timestamp().Msgf(format, v...)
}

// DebugWithFields function
func DebugWithFields(msg string, fields Fields) {
	infoLogger.Debug().Timestamp().Fields(fields).Msg(msg)
}

// Info function
func Info(args ...interface{}) {
	infoLogger.Info().Timestamp().Msg(fmt.Sprint(args...))
}

// Infoln function
func Infoln(args ...interface{}) {
	infoLogger.Info().Timestamp().Msg(fmt.Sprintln(args...))
}

// Infof function
func Infof(format string, v ...interface{}) {
	infoLogger.Info().Timestamp().Msgf(format, v...)
}

// InfoWithFields function
func InfoWithFields(msg string, fields Fields) {
	infoLogger.Info().Timestamp().Fields(fields).Msg(msg)
}

// Warn function
func Warn(args ...interface{}) {
	infoLogger.Warn().Timestamp().Msg(fmt.Sprint(args...))
}

// Warnln function
func Warnln(args ...interface{}) {
	infoLogger.Warn().Timestamp().Msg(fmt.Sprintln(args...))
}

// Warnf function
func Warnf(format string, v ...interface{}) {
	infoLogger.Warn().Timestamp().Msgf(format, v...)
}

// WarnWithFields function
func WarnWithFields(msg string, fields Fields) {
	infoLogger.Warn().Timestamp().Fields(fields).Msg(msg)
}

// Error function
func Error(args ...interface{}) {
	errorLogger.Error().Timestamp().Msg(fmt.Sprint(args...))
}

// Errorln function
func Errorln(args ...interface{}) {
	errorLogger.Error().Timestamp().Msg(fmt.Sprintln(args...))
}

// Errorf function
func Errorf(format string, v ...interface{}) {
	errorLogger.Error().Timestamp().Msgf(format, v...)
}

// ErrorWithFields function
func ErrorWithFields(msg string, fields Fields) {
	errorLogger.Error().Timestamp().Fields(fields).Msg(msg)
}

// Errors function to log errors package
func Errors(err error) {
	errorLogger.Error().Timestamp().Msg(err.Error())
}

// Fatal function
func Fatal(args ...interface{}) {
	errorLogger.Fatal().Timestamp().Msg(fmt.Sprint(args...))
}

// Fatalln function
func Fatalln(args ...interface{}) {
	errorLogger.Fatal().Timestamp().Msg(fmt.Sprintln(args...))
}

// Fatalf function
func Fatalf(format string, v ...interface{}) {
	errorLogger.Fatal().Timestamp().Msgf(format, v...)
}

// FatalWithFields function
func FatalWithFields(msg string, fields Fields) {
	errorLogger.Fatal().Timestamp().Fields(fields).Msg(msg)
}
