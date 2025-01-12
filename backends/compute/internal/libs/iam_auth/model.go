package iam_auth

import (
	"context"
	"encoding/json"
	"fmt"
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
	UserID      string
	DomainID    string
	ProjectID   string
	RolesJSON   string
	CatalogJSON string
	Roles       []string
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

	AuthContext := AuthContext{
		UserID:      xuser,
		DomainID:    xdomain,
		ProjectID:   xproject,
		RolesJSON:   xroles,
		CatalogJSON: xcatalog,
	}
	ctx = context.WithValue(ctx, KeyAuthContext, &AuthContext)

	return ctx
}

const HeaderXIdentityStatus = "x-identity-status"
const HeaderXIdentityStatusConfirmed = "Confirmed"

func AddAuthHeader(req *http.Request, authContext *AuthContext) {
	req.Header.Add(HeaderXIdentityStatus, HeaderXIdentityStatusConfirmed)
	req.Header.Add("x-user-domain-id", authContext.DomainID)
	req.Header.Add("x-project-domain-id", authContext.DomainID)
	req.Header.Add("x-user-id", authContext.UserID)
	req.Header.Add("x-project-id", authContext.ProjectID)
	req.Header.Add("x-service-catalog", authContext.CatalogJSON)
	req.Header.Add("x-roles", "admin")
	req.Header.Add("x-is-admin-project", "true")
	fmt.Println("req.Header", req.Header)
}

func GetAuthContext(ctx context.Context) (*AuthContext, error) {
	authCtx, ok := ctx.Value(KeyAuthContext).(*AuthContext)
	if !ok {
		return nil, tlog.Err(ctx,
			echo.NewHTTPError(http.StatusUnauthorized, "auth context is not found"))
	}

	roles := []string{}
	if authCtx.RolesJSON != "" {
		fmt.Println("authCtx.RolesJSON", authCtx.RolesJSON)
		if err := json.Unmarshal([]byte(authCtx.RolesJSON), &roles); err != nil {
			return nil, tlog.WrapErr(ctx, err, "failed to unmarshal roles")
		}
	}

	authCtx.Roles = roles

	return authCtx, nil
}
