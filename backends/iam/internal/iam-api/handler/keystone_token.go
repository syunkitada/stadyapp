package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_middleware"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) CreateKeystoneToken(ectx echo.Context) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	var input oapi.CreateKeystoneTokenInput

	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	token, tokenStr, err := self.api.CreateKeystoneToken(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	ectx.Response().Header().Set("x-subject-token", tokenStr)

	return tlog.BindEchoOK(ctx, ectx, token)
}
