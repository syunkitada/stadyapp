package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) FindRoles(ctx context.Context, params oapi.FindRolesParams) ([]oapi.Role, error) {
	dbRoles, err := self.db.FindRoles(ctx, &db.FindRolesInput{})

	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.db.FindRoles")
	}

	roles := []oapi.Role{}
	for _, dbRole := range dbRoles {
		roles = append(roles, oapi.Role{
			Id:   dbRole.ID,
			Name: dbRole.Name,
		})
	}

	return roles, nil
}

func (self *API) FindRoleByID(ctx context.Context, id uint64) (*oapi.Role, error) {
	dbRoles, err := self.db.FindRoles(ctx, &db.FindRolesInput{ID: id})

	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.db.FindRoles")
	}

	if len(dbRoles) > 1 {
		return nil, tlog.WrapEchoConflictError(ctx, "multiple roles found")
	}

	if len(dbRoles) == 0 {
		return nil, tlog.WrapEchoNotFoundError(ctx, "role does not found")
	}

	dbRole := dbRoles[0]
	role := &oapi.Role{
		Id:   dbRole.ID,
		Name: dbRole.Name,
	}

	return role, nil
}

func (self *API) AddRole(ctx context.Context, role *oapi.NewRole) error {
	if _, err := self.db.AddRole(ctx, &model.Role{
		Name: role.Name,
	}); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.db.AddRole")
	}

	return nil
}

func (self *API) DeleteRole(ctx context.Context, id uint64) error {
	if err := self.db.DeleteRole(ctx, id); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.db.DeleteRole")
	}

	return nil
}
