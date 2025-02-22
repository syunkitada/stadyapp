package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) CreateKeystoneApplicationCredential(ectx echo.Context, userID string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneApplicationCredentialInput

	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	token, err := self.api.CreateKeystoneApplicationCredential(ctx, userID, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneApplicationCredentialResponse{
		ApplicationCredential: *token,
	}
	return tlog.BindEchoOK(ctx, ectx, resp)
}
