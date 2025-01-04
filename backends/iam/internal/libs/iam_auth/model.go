package iam_auth

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type AuthData struct {
	Domain  string `json:"1"`
	User    string `json:"2"`
	Project string `json:"3"`
	Roles   string `json:"4"`
	Catalog string `json:"5"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	KeyName  string   `json:"1"`
	AuthData AuthData `json:"2"`
}

type AuthContext struct {
	UserID    string
	ProjectID string
}

type key int

const KeyAuthContext key = iota

func WithEchoContext(ectx echo.Context) context.Context {
	ctx := tlog.WithEchoContext(ectx)
	xuser := ectx.Request().Header.Get("x-user-id")
	AuthContext := AuthContext{UserID: xuser}
	ctx = context.WithValue(ctx, KeyAuthContext, &AuthContext)

	return ctx
}

func GetAuthContext(ctx context.Context) (*AuthContext, error) {
	authCtx, ok := ctx.Value(KeyAuthContext).(*AuthContext)
	if !ok {
		return nil, tlog.Err(ctx,
			echo.NewHTTPError(http.StatusUnauthorized, "auth context is not found"))
	}

	return authCtx, nil
}
