package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) CreateKeystoneRole(
	ctx context.Context, input *oapi.CreateKeystoneRoleInput) (*oapi.KeystoneRole, error) {
	dbRole, err := self.db.CreateRole(ctx, &db.CreateRoleInput{
		Name:        input.Role.Name,
		Description: input.Role.Description,
		Extra:       input.Role.AdditionalProperties,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domain, err := ConvertDBRoleToAPIRole(ctx, dbRole)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func (self *API) UpdateKeystoneRoleByID(
	ctx context.Context, id string, input *oapi.UpdateKeystoneRoleInput) (*oapi.KeystoneRole, error) {
	err := self.db.UpdateRoleByID(ctx, id, &db.UpdateRoleByIDInput{
		Name:        input.Role.Name,
		Description: input.Role.Description,
		Extra:       input.Role.AdditionalProperties,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domain, err := self.GetKeystoneRoleByID(ctx, id)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func (self *API) GetKeystoneRoles(
	ctx context.Context, input *oapi.GetKeystoneRolesParams) ([]oapi.KeystoneRole, error) {
	getRolesInput := db.GetRolesInput{}
	if input.Name != nil {
		getRolesInput.Name = *input.Name
	}

	dbRoles, err := self.db.GetRoles(ctx, &getRolesInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domains := []oapi.KeystoneRole{}

	for i := range dbRoles {
		domain, err := ConvertDBRoleToAPIRole(ctx, &dbRoles[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		domains = append(domains, *domain)
	}

	return domains, nil
}

func (self *API) GetKeystoneRoleByID(ctx context.Context, id string) (*oapi.KeystoneRole, error) {
	dbRoles, err := self.db.GetRoles(ctx, &db.GetRolesInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbRoles) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "domain not found"))
	}

	if len(dbRoles) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "domain is duplicated"))
	}

	domain, err := ConvertDBRoleToAPIRole(ctx, &dbRoles[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func ConvertDBRoleToAPIRole(ctx context.Context, dbRole *model.Role) (*oapi.KeystoneRole, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbRole.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneRole{
		Id:                   dbRole.ID,
		Name:                 dbRole.Name,
		Description:          dbRole.Description,
		AdditionalProperties: additionalProperties,
	}, nil
}

func (self *API) DeleteKeystoneRole(ctx context.Context, id string) error {
	err := self.db.DeleteRoleByID(ctx, id)
	if err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *API) AssignRoleToProject(ctx context.Context, roleID, userID, projectID string) error {
	splitedID := strings.Split(projectID, "@")
	fmt.Println("DEBUG AssignRole", splitedID)
	if len(splitedID) == 2 {
		if splitedID[0] == ProjectTagTeam {
			if err := self.db.AssignRoleToTeam(ctx, roleID, userID, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		} else if splitedID[0] == ProjectTagOrganization {
			if err := self.db.AssignRoleToOrganization(ctx, roleID, userID, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		}
	}

	if err := self.db.AssignRoleToProject(ctx, roleID, userID, projectID); err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *API) UnassignRoleFromProject(ctx context.Context, roleID, userID, projectID string) error {
	return nil
}

func (self *API) AssignRoleToDomain(ctx context.Context, roleID, userID, domainID string) error {
	return nil
}

func (self *API) UnassignRoleFromDomain(ctx context.Context, roleID, userID, domainID string) error {
	return nil
}
