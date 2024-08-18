package log

import (
	"log/slog"
	"os"
)

type ILogLevel interface {
	LogLevelDebug(msg string)
	LogLevelError(msg string)
	LogLevelWarn(msg string)
	LogLevelInfo(msg string)
}

type LogLevel struct {
	level *slog.LevelVar
}

// LogLevelDebug implements ILogLevel.
func (l LogLevel) LogLevelDebug(msg string) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.level}))
	l.level.Set(slog.LevelDebug)
	log.Debug(msg)
}

// LogLevelError implements ILogLevel.
func (l LogLevel) LogLevelError(msg string) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.level}))
	l.level.Set(slog.LevelError)
	log.Error(msg)
}

// LogLevelInfo implements ILogLevel.
func (l LogLevel) LogLevelInfo(msg string) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.level}))
	l.level.Set(slog.LevelInfo)
	log.Info(msg)
}

// LogLevelWarn implements ILogLevel.
func (l LogLevel) LogLevelWarn(msg string) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.level}))
	l.level.Set(slog.LevelWarn)
	log.Warn(msg)
}


func NewLogLevel() ILogLevel {
	return LogLevel{
		level: new(slog.LevelVar),
	}

}
