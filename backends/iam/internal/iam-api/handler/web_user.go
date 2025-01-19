package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetWebUser(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

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

	cookie := http.Cookie{
		Name:     "auth-token",
		Value:    "tokenhoge",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,         // TODO Configurable
		Expires:  time.Now().Add(1 * time.Hour), // TODO Configurable
	}
	ectx.SetCookie(&cookie)

	return tlog.BindEchoOK(ctx, ectx, webUser)
}
