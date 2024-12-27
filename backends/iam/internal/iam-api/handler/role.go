package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) FindRoles(ectx echo.Context, params oapi.FindRolesParams) error {
	ctx := tlog.WithEchoContext(ectx)

	items, err := self.api.FindRoles(ctx, params)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.FindRoles")
	}

	if err := ectx.JSON(http.StatusOK, items); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.JSON")
	}

	return nil
}

func (self *Handler) AddRole(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var newRole oapi.NewRole

	err := ectx.Bind(&newRole)

	if err != nil {
		err = sendHandlerError(ectx, http.StatusBadRequest, "Invalid format for NewRole")
		return tlog.WrapError(ctx, err, "failed to ectx.Bind")
	}

	if err := self.api.AddRole(ctx, &newRole); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.AddRole")
	}

	return nil
}

func (self *Handler) FindRoleByID(ectx echo.Context, itemID uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	item, err := self.api.FindRoleByID(ctx, itemID)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.FindRoleByID")
	}

	if err := ectx.JSON(http.StatusOK, item); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.JSON")
	}

	return nil
}

func (self *Handler) DeleteRole(ectx echo.Context, id uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	err := self.api.DeleteRole(ctx, id)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.DeleteRole")
	}

	if err := ectx.NoContent(http.StatusNoContent); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.NoContent")
	}

	return nil
}
