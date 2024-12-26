package tlogger

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type key struct{}

const OutputPathStdout = "stdout"

type Config struct {
	OutputPath string
}

func WithEchoContext(ctx echo.Context) context.Context {
	rctx := ctx.Request().Context()
	traceID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	logger := newSjsonLogger(&Config{OutputPath: OutputPathStdout}).With("trace_id", traceID)
	return context.WithValue(rctx, key{}, logger)
}

func FromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(key{}).(Logger)
	if !ok {
		return newSlogLogger(&Config{OutputPath: OutputPathStdout})
	}
	return logger
}

type Logger interface {
	With(...any) Logger
	Info(string, ...any)
}

type SlogLogger struct {
	it *slog.Logger
}

func newSlogLogger(conf *Config) Logger {
	switch conf.OutputPath {
	case OutputPathStdout:
		return &SlogLogger{
			it: slog.New(os.Stdout, slog.Ljson),
		}
	default:
		file, err := os.OpenFile(conf.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			os.Stderr.WriteString("Failed to open log file")
			os.Exit(1)
		}
		return &SlogLogger{
			it: slog.New(file, slog.Ljson),
		}
	}
}

func (l *SlogLogger) With(...any) Logger {
	l.it = l.it.With(args...)
	return I
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.it.Info(msg, args...)
}
