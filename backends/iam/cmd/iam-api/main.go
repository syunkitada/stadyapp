package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/handler"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func main() {
	conf := config.GetDefaultConfig()
	tlog.Init(&conf.Logger)
	ctx := tlog.NewContext()

	db := db.New(&conf.DB)
	db.MustOpen(ctx)

	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := handler.NewHandler(&conf, db)

	// This is how you set up a basic Echo router
	echoServer := echo.New()
	// Log all requests
	echoServer.Use(echomiddleware.RequestID())

	echoServer.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogResponseSize: true,
		LogLatency:      true,
		LogValuesFunc: func(c echo.Context, values echomiddleware.RequestLoggerValues) error {
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

	options := &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				slog.Info(fmt.Sprintf("authentication func: %+v", input))

				switch input.SecuritySchemeName {
				case "XUser":
					fmt.Println("XUser", input)
				default:
					return fmt.Errorf("unknown security scheme: %s", input.SecuritySchemeName)
				}

				return nil
			},
		},
	}

	echoServer.Use(middleware.OapiRequestValidatorWithOptions(swagger, options))

	echoServer.Use(echomiddleware.Recover())
	echoServer.Logger.SetLevel(log.INFO)
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.

	// We now register our petStore above as the handler for the interface
	oapi.RegisterHandlers(echoServer, apiHandler)

	// And we serve HTTP until the world ends.
	echoServer.Logger.Fatal(echoServer.Start(net.JoinHostPort("0.0.0.0", *port)))
}
