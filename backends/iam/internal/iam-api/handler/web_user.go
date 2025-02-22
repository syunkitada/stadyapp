package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetWebUser(ectx echo.Context, input oapi.GetWebUserParams) error {
	ctx := iam_auth.WithEchoContext(ectx)

	webUser, err := self.api.GetWebUser(ctx)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	createKeystoneTokenInput := oapi.CreateKeystoneTokenInput{
		Auth: oapi.CreateKeystoneTokenInputAuth{
			Identity: oapi.CreateKeystoneTokenInputAuthIdentity{
				Methods: []string{"web"},
			},
		},
	}

	if input.ProjectId != nil {
		createKeystoneTokenInput.Auth.Scope = &oapi.CreateKeystoneTokenInputAuthScope{
			Project: &oapi.CreateKeystoneTokenInputAuthScopeProject{
				Id: input.ProjectId,
			},
		}
	}

	token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &createKeystoneTokenInput)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	cookie := http.Cookie{
		Name:     "authtoken",
		Value:    tokenStr,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  token.ExpiresAt,
		Path:     "/api",
	}
	ectx.SetCookie(&cookie)

	return tlog.BindEchoOK(ctx, ectx, webUser)
}
