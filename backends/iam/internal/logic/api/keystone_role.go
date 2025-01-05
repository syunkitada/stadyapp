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

	role, err := ConvertDBRoleToAPIRole(ctx, dbRole)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return role, nil
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

	role, err := self.GetKeystoneRoleByID(ctx, id)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return role, nil
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

	roles := []oapi.KeystoneRole{}

	for i := range dbRoles {
		role, err := ConvertDBRoleToAPIRole(ctx, &dbRoles[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		roles = append(roles, *role)
	}

	return roles, nil
}

func (self *API) GetKeystoneRoleByID(ctx context.Context, id string) (*oapi.KeystoneRole, error) {
	dbRoles, err := self.db.GetRoles(ctx, &db.GetRolesInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbRoles) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "role not found"))
	}

	if len(dbRoles) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "role is duplicated"))
	}

	role, err := ConvertDBRoleToAPIRole(ctx, &dbRoles[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return role, nil
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

func (self *API) AssignKeystoneRoleToUserProject(ctx context.Context, roleID, userID, projectID string) error {
	splitedID := strings.Split(projectID, ProjectTagSeparator)
	if len(splitedID) == 2 {
		if splitedID[0] == ProjectTagTeam {
			fmt.Println("AssignRoleToProject Project")
			if err := self.db.AssignRoleToUserTeam(ctx, roleID, userID, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		} else if splitedID[0] == ProjectTagOrganization {
			if err := self.db.AssignRoleToUserOrganization(ctx, roleID, userID, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		}
	}

	if err := self.db.AssignRoleToUserProject(ctx, roleID, userID, projectID); err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *API) UnassignKeystoneRoleFromUserProject(ctx context.Context, roleID, userID, projectID string) error {
	return nil
}

func (self *API) AssignKeystoneRoleToGroupProject(ctx context.Context, roleID, groupID, projectID string) error {
	splitedGroupID := strings.Split(groupID, ProjectTagSeparator)
	if splitedGroupID[0] != ProjectTagTeam {
		return tlog.Err(ctx, echo.NewHTTPError(http.StatusBadRequest, "group id is invalid"))
	}
	teamID := splitedGroupID[1]

	splitedID := strings.Split(projectID, ProjectTagSeparator)
	if len(splitedID) == 2 {
		if splitedID[0] == ProjectTagOrganization {
			if err := self.db.AssignRoleToTeamOrganization(ctx, roleID, teamID, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		}
	}

	if err := self.db.AssignRoleToTeamProject(ctx, roleID, teamID, projectID); err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *API) UnassignKeystoneRoleFromGroupProject(ctx context.Context, roleID, groupID, projectID string) error {
	return nil
}

func (self *API) AssignKeystoneRoleToUserDomain(ctx context.Context, roleID, userID, domainID string) error {
	return nil
}

func (self *API) UnassignKeystoneRoleFromUserDomain(ctx context.Context, roleID, userID, domainID string) error {
	return nil
}

func (self *API) GetKeystoneRoleAssignments(
	ctx context.Context, input *oapi.GetKeystoneRoleAssignmentsParams) ([]oapi.KeystoneRoleAssignment, error) {
	roleAssignments := []oapi.KeystoneRoleAssignment{}

	teamRoleAssignments, err := self.db.GetTeamRoleAssignments(ctx, &db.GetTeamRoleAssignmentsInput{})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	for _, roleAssignment := range teamRoleAssignments {
		roleAssignments = append(roleAssignments, oapi.KeystoneRoleAssignment{
			Scope: &oapi.KeystoneRoleAssignmentScope{
				Project: &oapi.KeystoneRoleAssignmentProject{
					Name: roleAssignment.TeamName,
					Id:   ProjectTagTeam + ProjectTagSeparator + roleAssignment.TeamID,
					Domain: oapi.KeystoneRoleAssignmentDomain{
						Name: roleAssignment.DomainName,
						Id:   roleAssignment.DomainID,
					},
				},
			},
			Role: &oapi.KeystoneRoleAssignmentRole{
				Id:   roleAssignment.RoleID,
				Name: roleAssignment.RoleName,
			},
			User: &oapi.KeystoneRoleAssignmentUser{
				Id:   roleAssignment.UserID,
				Name: roleAssignment.UserName,
				Domain: oapi.KeystoneRoleAssignmentDomain{
					Id:   roleAssignment.DomainID,
					Name: roleAssignment.DomainName,
				},
			},
		})
	}

	projectRoleAssignments, err := self.db.GetProjectRoleAssignments(ctx, &db.GetProjectRoleAssignmentsInput{})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	for _, roleAssignment := range projectRoleAssignments {
		keystoneRoleAssignment := oapi.KeystoneRoleAssignment{
			Scope: &oapi.KeystoneRoleAssignmentScope{
				Project: &oapi.KeystoneRoleAssignmentProject{
					Name: roleAssignment.ProjectName,
					Id:   roleAssignment.ProjectID,
					Domain: oapi.KeystoneRoleAssignmentDomain{
						Name: roleAssignment.DomainName,
						Id:   roleAssignment.DomainID,
					},
				},
			},
			Role: &oapi.KeystoneRoleAssignmentRole{
				Id:   roleAssignment.RoleID,
				Name: roleAssignment.RoleName,
			},
		}

		if roleAssignment.UserID != nil {
			keystoneRoleAssignment.User = &oapi.KeystoneRoleAssignmentUser{
				Id:   *roleAssignment.UserID,
				Name: roleAssignment.UserName,
				Domain: oapi.KeystoneRoleAssignmentDomain{
					Id:   roleAssignment.DomainID,
					Name: roleAssignment.DomainName,
				},
			}
		}

		if roleAssignment.TeamID != nil {
			keystoneRoleAssignment.Group = &oapi.KeystoneRoleAssignmentGroup{
				Id:   ProjectTagTeam + ProjectTagSeparator + *roleAssignment.TeamID,
				Name: roleAssignment.TeamName,
				Domain: oapi.KeystoneRoleAssignmentDomain{
					Id:   roleAssignment.DomainID,
					Name: roleAssignment.DomainName,
				},
			}
		}

		roleAssignments = append(roleAssignments, keystoneRoleAssignment)
	}

	return roleAssignments, nil
}
