package tlog

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

var _ logger.Interface = (*Logger)(nil)

type Logger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// LogMode implements logger.Interface.
func (log *Logger) LogMode(level logger.LogLevel) logger.Interface { //nolint:ireturn
	log.LogLevel = level

	return log
}

// Info implements logger.Interface.
func (log *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if log.LogLevel >= logger.Info {
		slog.InfoContext(ctx, fmt.Sprintf(msg, data...))
	}
}

// Warn implements logger.Interface.
func (log *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if log.LogLevel >= logger.Warn {
		slog.WarnContext(ctx, fmt.Sprintf(msg, data...))
	}
}

// Error implements logger.Interface.
func (log *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if log.LogLevel >= logger.Error {
		slog.ErrorContext(ctx, fmt.Sprintf(msg, data...))
	}
}

// Trace implements logger.Interface.
func (log *Logger) Trace(ctx context.Context, begin time.Time,
	queryFn func() (sql string, rowsAffected int64), err error) {
	if log.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := queryFn()

	switch {
	case err != nil && log.LogLevel >= logger.Error:
		slog.ErrorContext(ctx, "sql error",
			slog.String("latency", elapsed.String()),
			slog.Int64("rows", rows),
			slog.String("sql", sql),
			slog.String("error", err.Error()),
		)
	case elapsed > log.SlowThreshold && log.SlowThreshold != 0 && log.LogLevel >= logger.Warn:
		slog.WarnContext(ctx, "sql slow",
			slog.String("latency", elapsed.String()),
			slog.Int64("rows", rows),
			slog.String("sql", sql))
	case log.LogLevel == logger.Info:
		slog.InfoContext(ctx, "sql",
			slog.String("latency", elapsed.String()),
			slog.Int64("rows", rows),
			slog.String("sql", sql))
	}
}
