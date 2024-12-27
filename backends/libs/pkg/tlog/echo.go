package tlog

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func WrapEchoNotFoundError(ctx context.Context, msg string, args ...any) error {
	err := echo.NewHTTPError(http.StatusNotFound, msg)
	slog.ErrorContext(ctx, msg, args...)

	return fmt.Errorf("%s: %w", msg, err)
}

func WrapEchoConflictError(ctx context.Context, msg string, args ...any) error {
	err := echo.NewHTTPError(http.StatusConflict, msg)
	slog.ErrorContext(ctx, msg, args...)

	return fmt.Errorf("%s: %w", msg, err)
}
