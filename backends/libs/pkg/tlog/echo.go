package tlog

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoErrorResponse struct {
	Message string `json:"message"`
}

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

func BindEchoOK(ctx context.Context, ectx echo.Context, response interface{}) error {
	if err := ectx.JSON(http.StatusOK, response); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ectx.JSON")
	}

	return nil
}

func BindEchoNoContent(ctx context.Context, ectx echo.Context) error {
	if err := ectx.NoContent(http.StatusNoContent); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ectx.NoContent")
	}

	return nil
}

func BindEchoBadRequest(ctx context.Context, ectx echo.Context, err error) error {
	slog.ErrorContext(ctx, err.Error())

	if err := ectx.JSON(http.StatusBadRequest, &EchoErrorResponse{
		Message: err.Error(),
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ectx.JSON")
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func BindEchoError(ctx context.Context, ectx echo.Context, err error) error {
	slog.ErrorContext(ctx, err.Error())

	code := http.StatusInternalServerError

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		code = httpErr.Code
	}

	if err := ectx.JSON(code, &EchoErrorResponse{
		Message: err.Error(),
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ectx.JSON")
	}

	return err
}
