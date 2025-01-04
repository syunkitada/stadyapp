package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneUsers(ectx echo.Context, input oapi.GetKeystoneUsersParams) error {
	ctx := iam_auth.WithEchoContext(ectx)

	users, err := self.api.GetKeystoneUsers(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneUsers{
		Users: users,
	}
	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneUserByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	user, err := self.api.GetKeystoneUserByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, user)
}

func (self *Handler) CreateKeystoneUser(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneUserInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	user, err := self.api.CreateKeystoneUser(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, user)
}

func (self *Handler) DeleteKeystoneUserByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.DeleteKeystoneUser(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
