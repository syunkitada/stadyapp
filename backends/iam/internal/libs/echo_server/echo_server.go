package echo_server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	oapi_echo_middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	Port int
}

func New(ctx context.Context, conf *Config, swagger *openapi3.T) *echo.Echo {
	// This is how you set up a basic Echo router
	echoServer := echo.New()

	echoServer.Use(middleware.RequestID())

	echoServer.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogResponseSize: true,
		LogLatency:      true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			ctx := tlog.WithEchoContext(c)
			slog.InfoContext(ctx,
				"request was processed",
				slog.String("method", values.Method),
				slog.String("uri", values.URI),
				slog.Int("status", values.Status),
				slog.Int64("response_size", values.ResponseSize),
				slog.String("latency", values.Latency.String()),
			)

			return nil
		},
	}))

	options := &oapi_echo_middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				switch input.SecuritySchemeName {
				case "XUserNameHeader":
					xuserName := input.RequestValidationInput.Request.Header.Get("x-user-name")
					if xuserName == "" {
						return errors.New("missing x-user-name header") //nolint:err113,wrapcheck
					}
				case "XAuthTokenHeader":
					xauthToken := input.RequestValidationInput.Request.Header.Get("x-auth-token")
					if xauthToken == "" {
						return errors.New("missing x-auth-token header") //nolint:err113,wrapcheck
					}
					fmt.Println("x-auth-token:", xauthToken)
				default:
					return fmt.Errorf("unknown security scheme: %s", input.SecuritySchemeName) //nolint:err113,wrapcheck
				}

				return nil
			},
		},
	}

	echoServer.Use(oapi_echo_middleware.OapiRequestValidatorWithOptions(swagger, options))

	echoServer.Use(middleware.Recover())

	echoServer.Logger.SetLevel(log.INFO)

	return echoServer
}

type IAMTokenAuth struct {
	APIEndpoint string
}

type key int

const KeyAuthContext key = iota

type AuthContext struct {
	User string
}

func WithAuthEchoContext(ectx echo.Context) context.Context {
	ctx := tlog.WithEchoContext(ectx)
	xuser := ectx.Request().Header.Get("x-user-name")
	AuthContext := AuthContext{User: xuser}
	ctx = context.WithValue(ctx, KeyAuthContext, AuthContext)

	return ctx
}
