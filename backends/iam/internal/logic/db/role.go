package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) GetRoleByID(ctx context.Context, id string) (*model.Role, error) {
	dbRoles, err := self.GetRoles(ctx, &db.GetRolesInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbRoles) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "role does not found"))
	}

	if len(dbRoles) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "role is duplicated"))
	}

	return &dbRoles[0], nil
}

func (self *DB) GetRoles(ctx context.Context, input *db.GetRolesInput) ([]model.Role, error) {
	query := self.DB.WithContext(ctx).Model(model.Role{}).
		Select("id,name,description,extra")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	roles := []model.Role{}
	if err := query.Scan(&roles).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roles, nil
}

func (self *DB) CreateRole(ctx context.Context, input *db.CreateRoleInput) (*model.Role, error) {
	bytes, err := json.Marshal(input.Extra)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	role := model.Role{
		Name:  input.Name,
		Extra: string(bytes),
	}

	if input.ID == nil {
		role.ID = uuid.New().String()
	} else {
		role.ID = *input.ID
	}

	if input.Description != nil {
		role.Description = *input.Description
	}

	if err := self.DB.WithContext(ctx).Create(&role).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &role, nil
}

func (self *DB) UpdateRoleByID(ctx context.Context, id string, input *db.UpdateRoleByIDInput) error {
	data := map[string]interface{}{}
	if input.Name != nil {
		data["name"] = *input.Name
	}

	if input.Description != nil {
		data["description"] = *input.Description
	}

	if input.Extra != nil {
		bytes, err := json.Marshal(input.Extra)
		if err != nil {
			return tlog.Err(ctx, err)
		}

		data["extra"] = string(bytes)
	}

	if len(data) > 0 {
		if err := self.DB.WithContext(ctx).Model(model.Role{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteRoleByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Role{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) GetProjectRoleAssignments(ctx context.Context, input *db.GetProjectRoleAssignmentsInput) ([]model.ProjectRoleAssignmentDetail, error) {
	query := self.DB.WithContext(ctx).Table("project_role_assignments").
		Select("roles.id AS role_id, roles.name AS role_name, " +
			"users.id AS user_id, users.name AS user_name, " +
			"teams.id AS team_id, teams.name AS team_name, " +
			"projects.id AS project_id, projects.name AS project_name, " +
			"domains.id AS domain_id, domains.name AS domain_name").
		Joins("JOIN roles ON roles.id = project_role_assignments.role_id").
		Joins("LEFT JOIN users ON users.id = project_role_assignments.user_id").
		Joins("LEFT JOIN teams ON teams.id = project_role_assignments.team_id").
		Joins("JOIN projects ON projects.id = project_role_assignments.project_id").
		Joins("JOIN domains ON domains.id = projects.domain_id")

	roleAssignments := []model.ProjectRoleAssignmentDetail{}
	if err := query.Scan(&roleAssignments).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roleAssignments, nil
}

func (self *DB) GetUserProjectRoles(ctx context.Context, input *db.GetUserProjectRolesInput) ([]model.UserProjectRole, error) {
	query := self.DB.WithContext(ctx).Table("project_role_assignments").
		Select("project_role_assignments.project_id, project_role_assignments.role_id, "+
			"team_role_assignments.role_id AS team_role_id").
		Joins("LEFT JOIN team_role_assignments ON team_role_assignments.team_id = project_role_assignments.team_id").
		Where("(project_role_assignments.user_id = ? OR team_role_assignments.user_id = ?)", input.UserID, input.UserID)

	if input.ProjectID != "" {
		query = query.Where("project_role_assignments.project_id = ?", input.ProjectID)
	}

	roles := []model.UserProjectRole{}
	if err := query.Scan(&roles).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roles, nil
}

func (self *DB) AssignRoleToUserProject(ctx context.Context, roleID, userID, projectID string) error {
	roleAssignment := model.ProjectRoleAssignment{
		RoleID:    roleID,
		UserID:    &userID,
		ProjectID: projectID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromUserProject(ctx context.Context, roleID, userID, projectID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND project_id = ?", roleID, userID, projectID).
		Delete(model.ProjectRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToTeamProject(ctx context.Context, roleID, teamID, projectID string) error {
	fmt.Println("AssignRoleToProject Project", roleID, teamID, projectID)
	roleAssignment := model.ProjectRoleAssignment{
		RoleID:    roleID,
		TeamID:    &teamID,
		ProjectID: projectID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromTeamProject(ctx context.Context, roleID, teamID, projectID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND team_id = ? AND project_id = ?", roleID, teamID, projectID).
		Delete(model.ProjectRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) GetDomainRoleAssignments(ctx context.Context, input *db.GetDomainRoleAssignmentsInput) ([]model.DomainRoleAssignmentDetail, error) {
	query := self.DB.WithContext(ctx).Table("domain_role_assignments").
		Select("roles.id AS role_id, roles.name AS role_name, " +
			"users.id AS user_id, users.name AS user_name, " +
			"teams.id AS team_id, teams.name AS team_name, " +
			"domains.id AS domain_id, domains.name AS domain_name").
		Joins("JOIN roles ON roles.id = domain_role_assignments.role_id").
		Joins("LEFT JOIN users ON users.id = domain_role_assignments.user_id").
		Joins("LEFT JOIN teams ON teams.id = domain_role_assignments.team_id").
		Joins("JOIN domains ON domains.id = domain_role_assignments.domain_id")

	roleAssignments := []model.DomainRoleAssignmentDetail{}
	if err := query.Scan(&roleAssignments).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roleAssignments, nil
}

func (self *DB) AssignRoleToUserDomain(ctx context.Context, roleID, userID, domainID string) error {
	roleAssignment := model.DomainRoleAssignment{
		RoleID:   roleID,
		UserID:   &userID,
		DomainID: domainID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromUserDomain(ctx context.Context, roleID, userID, domainID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND domain_id = ?", roleID, userID, domainID).
		Delete(model.DomainRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToTeamDomain(ctx context.Context, roleID, teamID, domainID string) error {
	roleAssignment := model.DomainRoleAssignment{
		RoleID:   roleID,
		TeamID:   &teamID,
		DomainID: domainID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromTeamDomain(ctx context.Context, roleID, teamID, domainID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND team_id = ? AND domain_id = ?", roleID, teamID, domainID).
		Delete(model.DomainRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) GetTeamRoleAssignments(ctx context.Context, input *db.GetTeamRoleAssignmentsInput) ([]model.TeamRoleAssignmentDetail, error) {
	query := self.DB.WithContext(ctx).Table("team_role_assignments").
		Select("roles.id AS role_id, roles.name AS role_name, " +
			"users.id AS user_id, users.name AS user_name, " +
			"teams.id AS team_id, teams.name AS team_name, " +
			"domains.id AS domain_id, domains.name AS domain_name").
		Joins("JOIN roles ON roles.id = team_role_assignments.role_id").
		Joins("JOIN users ON users.id = team_role_assignments.user_id").
		Joins("JOIN teams ON teams.id = team_role_assignments.team_id").
		Joins("JOIN domains ON domains.id = teams.domain_id")

	roleAssignments := []model.TeamRoleAssignmentDetail{}
	if err := query.Scan(&roleAssignments).Debug().Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roleAssignments, nil
}

func (self *DB) AssignRoleToUserTeam(ctx context.Context, roleID, userID, teamID string) error {
	roleAssignment := model.TeamRoleAssignment{
		RoleID: roleID,
		UserID: userID,
		TeamID: teamID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromUserTeam(ctx context.Context, roleID, userID, teamID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND team_id = ?", roleID, userID, teamID).
		Delete(model.TeamRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) GetOrganizationRoleAssignments(
	ctx context.Context, input *db.GetOrganizationRoleAssignmentsInput) ([]model.OrganizationRoleAssignmentDetail, error) {
	query := self.DB.WithContext(ctx).Table("organization_role_assignments").
		Select("roles.id AS role_id, roles.name AS role_name, " +
			"users.id AS user_id, users.name AS user_name, " +
			"teams.id AS team_id, teams.name AS team_name, " +
			"organizations.id AS organization_id, organizations.name AS organization_name, " +
			"domains.id AS domain_id, domains.name AS domain_name").
		Joins("JOIN roles ON roles.id = organization_role_assignments.role_id").
		Joins("LEFT JOIN users ON users.id = organization_role_assignments.user_id").
		Joins("LEFT JOIN teams ON teams.id = organization_role_assignments.team_id").
		Joins("JOIN organizations ON organizations.id = organization_role_assignments.organization_id").
		Joins("JOIN domains ON domains.id = organizations.domain_id")

	roleAssignments := []model.OrganizationRoleAssignmentDetail{}
	if err := query.Scan(&roleAssignments).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return roleAssignments, nil
}

func (self *DB) AssignRoleToUserOrganization(ctx context.Context, roleID, userID, organizationID string) error {
	roleAssignment := model.OrganizationRoleAssignment{
		RoleID:         roleID,
		UserID:         &userID,
		OrganizationID: organizationID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromUserOrganization(ctx context.Context, roleID, userID, teamID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND team_id = ?", roleID, userID, teamID).
		Delete(model.OrganizationRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToTeamOrganization(ctx context.Context, roleID, teamID, organizationID string) error {
	roleAssignment := model.OrganizationRoleAssignment{
		RoleID:         roleID,
		TeamID:         &teamID,
		OrganizationID: organizationID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromTeamOrganization(ctx context.Context, roleID, teamID, organizationID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND team_id = ? AND organization_id = ?", roleID, teamID, organizationID).
		Delete(model.OrganizationRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
