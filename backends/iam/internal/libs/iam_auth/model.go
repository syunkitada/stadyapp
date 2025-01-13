package iam_auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type AuthData struct {
	DomainID  string          `json:"domain"`
	UserID    string          `json:"user"`
	ProjectID string          `json:"project"`
	Roles     string          `json:"roles"`
	Catalog   string          `json:"catalog"`
	Inherit   bool            `json:"inherit"`
	RoleSet   map[string]bool `json:"-"`
	ExpiresAt time.Time       `json:"-"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	KeyName  string   `json:"key"`
	AuthData AuthData `json:"auth"`
}

type AuthContext struct {
	UserID      string
	DomainID    string
	ProjectID   string
	RolesStr    string
	CatalogJSON string
	Roles       []string
	Inherit     bool
}

type key int

const KeyAuthContext key = iota

func WithEchoContext(ectx echo.Context) context.Context {
	ctx := tlog.WithEchoContext(ectx)
	xuser := ectx.Request().Header.Get("x-user-id")
	xdomain := ectx.Request().Header.Get("x-domain-id")
	xproject := ectx.Request().Header.Get("x-project-id")
	xroles := ectx.Request().Header.Get("x-roles")
	xcatalog := ectx.Request().Header.Get("x-catalog")
	xinherit := ectx.Request().Header.Get("x-inherit")
	inherit, _ := strconv.ParseBool(xinherit)

	AuthContext := AuthContext{
		UserID:      xuser,
		DomainID:    xdomain,
		ProjectID:   xproject,
		RolesStr:    xroles,
		CatalogJSON: xcatalog,
		Inherit:     inherit,
	}
	ctx = context.WithValue(ctx, KeyAuthContext, &AuthContext)

	return ctx
}

func GetAuthContext(ctx context.Context) (*AuthContext, error) {
	authCtx, ok := ctx.Value(KeyAuthContext).(*AuthContext)
	if !ok {
		return nil, tlog.Err(ctx,
			echo.NewHTTPError(http.StatusUnauthorized, "auth context is not found"))
	}

	roles := strings.Split(authCtx.RolesStr, ",")

	authCtx.Roles = roles

	return authCtx, nil
}
