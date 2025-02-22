package echo_server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	oapi_echo_middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	Port         int
	AllowOrigins []string
}

func New(ctx context.Context, conf *Config, swagger *openapi3.T, iamAuth *iam_auth.IAMAuth) *echo.Echo {
	// This is how you set up a basic Echo router
	echoServer := echo.New()

	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     conf.AllowOrigins,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	echoServer.Use(middleware.RequestID())

	echoServer.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			ctx := tlog.WithEchoContext(c)
			slog.InfoContext(ctx,
				"payload",
				slog.String("req_body", string(reqBody)),
				slog.String("res_body", string(resBody)),
			)
		},
	}))

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
				case "XUserIDHeader":
					if err := iamAuth.AuthUserID(ctx, input.RequestValidationInput.Request.Header); err != nil {
						return tlog.Err(ctx, err)
					}

					return nil

				case "XAuthTokenHeader":
					if err := iamAuth.AuthToken(ctx, input.RequestValidationInput.Request.Header); err != nil {
						return tlog.Err(ctx, err)
					}

					return nil
				}

				return tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized,
					"unknown security scheme: "+input.SecuritySchemeName))
			},
		},
	}

	echoServer.Use(oapi_echo_middleware.OapiRequestValidatorWithOptions(swagger, options))

	echoServer.Use(middleware.Recover())

	echoServer.Logger.SetLevel(log.INFO)

	return echoServer
}
