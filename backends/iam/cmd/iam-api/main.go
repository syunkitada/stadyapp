package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/handler"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_server"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func main() { //nolint:funlen
	conf := config.GetDefaultConfig()
	tlog.Init(&conf.Logger)
	ctx := tlog.NewContext()

	iamTokenAuth := iam_token_auth.New(&conf.IAMTokenAuth)

	db := db.New(&conf.DB)
	db.MustOpen(ctx)

	api := api.New(&conf, db, iamTokenAuth)

	port := flag.String("port", "10081", "Port for test HTTP server")
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
	apiHandler := handler.NewHandler(&conf, api)

	// This is how you set up a basic Echo router
	echoServer := echo_server.New(ctx, nil, swagger, iamTokenAuth)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.

	// We now register our petStore above as the handler for the interface
	oapi.RegisterHandlers(echoServer, apiHandler)

	// And we serve HTTP until the world ends.
	echoServer.Logger.Fatal(echoServer.Start(net.JoinHostPort("0.0.0.0", *port)))
}
