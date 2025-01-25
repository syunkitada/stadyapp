package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetWebUser(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)

	// input := oapi.CreateKeystoneTokenInput{
	// 	Auth: oapi.CreateKeystoneTokenInputAuth{
	// 		Identity: oapi.CreateKeystoneTokenInputAuthIdentity{
	// 			Methods: []string{protocol},
	// 		},
	// 	},
	// }

	// token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &input)
	// if err != nil {
	// 	return tlog.BindEchoError(ctx, ectx, err)
	// }

	webUser, err := self.api.GetWebUser(ctx)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	input := oapi.CreateKeystoneTokenInput{
		Auth: oapi.CreateKeystoneTokenInputAuth{
			Identity: oapi.CreateKeystoneTokenInputAuthIdentity{
				Methods: []string{"web"},
			},
		},
	}

	token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &input)
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
