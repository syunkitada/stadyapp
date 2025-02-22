package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/handler"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/echo_server"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func main() {
	conf := config.GetDefaultConfig()
	tlog.Init(&conf.Logger)
	ctx := tlog.NewContext()

	iamAuth := iam_auth.New(&conf.IAMAuth)

	db := db.New(&conf.DB)
	db.MustOpen(ctx)

	api := api.New(&conf, db, iamAuth)

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
	echoServer := echo_server.New(ctx, &conf.Server, swagger, iamAuth)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.

	// We now register our petStore above as the handler for the interface
	oapi.RegisterHandlers(echoServer, apiHandler)

	// And we serve HTTP until the world ends.
	echoServer.Logger.Fatal(echoServer.Start(net.JoinHostPort("0.0.0.0", strconv.Itoa(conf.Server.Port))))
}
