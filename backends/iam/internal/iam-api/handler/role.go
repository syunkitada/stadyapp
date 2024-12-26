package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/tlog"
)

func (self *Handler) FindRoles(ectx echo.Context, params oapi.FindRolesParams) error {
	ctx := tlog.WithEchoContext(ectx)

	items, err := self.api.FindRoles(ctx, params)

	if err != nil {
		return fmt.Errorf("self.api.FindRoles: error=%w", err)
	}

	if err := ectx.JSON(http.StatusOK, items); err != nil {
		return fmt.Errorf("ectx.JSON: error=%w", err)
	}

	return nil
}

func (self *Handler) AddRole(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var newRole oapi.NewRole

	err := ectx.Bind(&newRole)

	if err != nil {
		return sendHandlerError(ectx, http.StatusBadRequest, "Invalid format for NewRole")
	}

	if err := self.api.AddRole(ctx, &newRole); err != nil {
		return fmt.Errorf("self.api.AddRole: error=%w", err)
	}

	return nil
}

func (self *Handler) FindRoleByID(ectx echo.Context, itemId uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	item, err := self.api.FindRoleByID(ctx, itemId)

	if err != nil {
		return fmt.Errorf("self.api.FindRoleByID: error=%w", err)
	}

	if err := ectx.JSON(http.StatusOK, item); err != nil {
		return fmt.Errorf("ectx.JSON: error=%w", err)
	}

	return nil
}

func (self *Handler) DeleteRole(ectx echo.Context, id uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	err := self.api.DeleteRole(ctx, id)

	if err != nil {
		return fmt.Errorf("self.api.DeleteRole: error=%w", err)
	}

	if err := ectx.NoContent(http.StatusNoContent); err != nil {
		return fmt.Errorf("ectx.NoContent: error=%w", err)
	}

	return nil
}
