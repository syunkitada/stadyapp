package echo_middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type key int

const KeyAuthContext key = iota

func AuthenticationFunc(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	switch input.SecuritySchemeName {
	case "XUser":
		xuserName := input.RequestValidationInput.Request.Header.Get("x-user-name")
		if xuserName == "" {
			return errors.New("missing x-user-name header") //nolint:err113,wrapcheck
		}
	default:
		return fmt.Errorf("unknown security scheme: %s", input.SecuritySchemeName) //nolint:err113,wrapcheck
	}

	return nil
}

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
