package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) CreateKeystoneToken(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneTokenInput

	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	ectx.Response().Header().Set("x-subject-token", tokenStr)

	resp := oapi.KeystoneTokenResponse{
		Token: *token,
	}
	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) CreateKeystoneFederationAuthToken(ectx echo.Context, provider, protocol string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	input := oapi.CreateKeystoneTokenInput{
		Auth: oapi.CreateKeystoneTokenInputAuth{
			Identity: oapi.CreateKeystoneTokenInputAuthIdentity{
				Methods: []string{protocol},
			},
		},
	}

	token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	ectx.Response().Header().Set("x-subject-token", tokenStr)

	resp := oapi.KeystoneTokenResponse{
		Token: *token,
	}
	return tlog.BindEchoOK(ctx, ectx, resp)
}
