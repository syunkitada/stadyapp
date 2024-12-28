package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_middleware"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) FindRoles(ectx echo.Context, params oapi.FindRolesParams) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	items, err := self.api.FindRoles(ctx, params)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, items)
}

func (self *Handler) AddRole(ectx echo.Context) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	var newRole oapi.NewRole

	if err := ectx.Bind(&newRole); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	if err := self.api.AddRole(ctx, &newRole); err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return nil
}

func (self *Handler) FindRoleByID(ectx echo.Context, itemID uint64) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	item, err := self.api.FindRoleByID(ctx, itemID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, item)
}

func (self *Handler) DeleteRole(ectx echo.Context, id uint64) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	err := self.api.DeleteRole(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
