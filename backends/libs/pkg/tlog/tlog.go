package tlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const OutputPathStdout = "stdout"
const TagTraceID = "trace_id"
const TagFile = "file"
const GormFile = "gorm.go"

type key int

const KeyTraceID key = iota

type Config struct {
	OutputPath string
}

func GetDefaultConfig() Config {
	return Config{
		OutputPath: OutputPathStdout,
	}
}

var _ slog.Handler = &TLogHandler{}

type TLogHandler struct {
	slog.Handler
}

func (h *TLogHandler) Handle(ctx context.Context, record slog.Record) error {
	_, file, line, _ := runtime.Caller(4) //nolint:mnd
	if filepath.Base(file) == GormFile {
		_, file, line, _ = runtime.Caller(5) //nolint:mnd
	}

	record.AddAttrs(slog.Attr{Key: TagFile, Value: slog.StringValue(fmt.Sprintf("%s:%d", file, line))})

	if v := ctx.Value(KeyTraceID); v != nil {
		record.AddAttrs(slog.Attr{Key: string(TagTraceID), Value: slog.AnyValue(v)})
	}

	return h.Handler.Handle(ctx, record) //nolint:wrapcheck
}

func Init(conf *Config) {
	opts := &slog.HandlerOptions{}

	var handler slog.Handler

	switch conf.OutputPath {
	case OutputPathStdout:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		file, err := os.OpenFile(conf.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			os.Stderr.WriteString("Failed to open log file")
			os.Exit(1)
		}

		handler = slog.NewJSONHandler(file, opts)
	}

	logger := slog.New(&TLogHandler{handler})
	slog.SetDefault(logger)
}

func WithEchoContext(ectx echo.Context) context.Context {
	ctx := ectx.Request().Context()
	traceID := ectx.Response().Header().Get(echo.HeaderXRequestID)
	ctx = context.WithValue(ctx, KeyTraceID, traceID)

	return ctx
}

func NewContext() context.Context {
	ctx := context.Background()
	traceID := uuid.New().String()
	ctx = context.WithValue(ctx, KeyTraceID, traceID)

	return ctx
}

func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}

func WrapError(ctx context.Context, err error, msg string, args ...any) error {
	slog.ErrorContext(ctx, msg, args...)

	return fmt.Errorf("%s: %w", msg, err)
}

func FatalStdout(msg string, args ...any) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(msg, args...))
	os.Exit(1)
}
