package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneRoles(ectx echo.Context, input oapi.GetKeystoneRolesParams) error {
	ctx := iam_auth.WithEchoContext(ectx)

	roles, err := self.api.GetKeystoneRoles(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneRolesResponse{
		Roles: roles,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneRoleByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	role, err := self.api.GetKeystoneRoleByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneRoleResponse{
		Role: *role,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) CreateKeystoneRole(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneRoleInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	role, err := self.api.CreateKeystoneRole(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneRoleResponse{
		Role: *role,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) UpdateKeystoneRoleByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.UpdateKeystoneRoleInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	role, err := self.api.UpdateKeystoneRoleByID(ctx, id, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneRoleResponse{
		Role: *role,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) DeleteKeystoneRoleByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.DeleteKeystoneRole(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}

func (self *Handler) AssignRoleToProject(ectx echo.Context, projectID string, userID string, roleID string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.AssignRoleToProject(ctx, roleID, userID, projectID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}

func (self *Handler) UnassignRoleFromProject(ectx echo.Context, projectID string, userID string, roleID string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.UnassignRoleFromProject(ctx, roleID, userID, projectID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}

func (self *Handler) AssignRoleToDomain(ectx echo.Context, projectID string, userID string, roleID string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.AssignRoleToDomain(ctx, roleID, userID, projectID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}

func (self *Handler) UnassignRoleFromDomain(ectx echo.Context, projectID string, userID string, roleID string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.UnassignRoleFromDomain(ctx, roleID, userID, projectID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
