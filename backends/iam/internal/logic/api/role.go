package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) FindRoles(ctx context.Context, params oapi.FindRolesParams) (items []oapi.Role, err error) {
	dbRoles, err := self.db.FindRoles(ctx, &db.FindRolesInput{})
	if err != nil {
		return nil, err
	}

	for _, dbRole := range dbRoles {
		items = append(items, oapi.Role{
			Id:   dbRole.ID,
			Name: dbRole.Name,
		})
	}
	return items, nil
}

func (self *API) FindRoleByID(ctx context.Context, id uint64) (item oapi.Role, err error) {
	dbRoles, err := self.db.FindRoles(ctx, &db.FindRolesInput{ID: id})
	if err != nil {
		return item, err
	}
	if len(dbRoles) > 1 {
		return item, echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("multiple items found: id=%d", id))
	}
	if len(dbRoles) == 0 {
		return item, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("item not found: id=%d", id))
	}
	dbRole := dbRoles[0]
	item = oapi.Role{
		Id:   dbRole.ID,
		Name: dbRole.Name,
	}
	return item, nil
}

func (self *API) AddRole(ctx context.Context, item *oapi.NewRole) error {
	if _, err := self.db.AddRole(ctx, &model.Role{
		Name: item.Name,
	}); err != nil {
		return err
	}
	return nil
}

func (self *API) DeleteRole(ctx context.Context, id uint64) error {
	if err := self.db.DeleteRole(ctx, id); err != nil {
		return err
	}
	return nil
}
