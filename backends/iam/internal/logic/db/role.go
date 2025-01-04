package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

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

func (self *DB) AssignRoleToProject(ctx context.Context, roleID, userID, projectID string) error {
	roleAssignment := model.ProjectRoleAssignment{
		RoleID:    roleID,
		UserID:    userID,
		ProjectID: projectID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromProject(ctx context.Context, roleID, userID, projectID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND project_id = ?", roleID, userID, projectID).
		Delete(model.ProjectRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToDomain(ctx context.Context, roleID, userID, domainID string) error {
	roleAssignment := model.DomainRoleAssignment{
		RoleID:   roleID,
		UserID:   userID,
		DomainID: domainID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromDomain(ctx context.Context, roleID, userID, domainID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND domain_id = ?", roleID, userID, domainID).
		Delete(model.DomainRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToTeam(ctx context.Context, roleID, userID, teamID string) error {
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

func (self *DB) UnassignRoleFromTeam(ctx context.Context, roleID, userID, teamID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND team_id = ?", roleID, userID, teamID).
		Delete(model.TeamRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) AssignRoleToOrganization(ctx context.Context, roleID, userID, teamID string) error {
	roleAssignment := model.OrganizationRoleAssignment{
		RoleID:         roleID,
		UserID:         userID,
		OrganizationID: teamID,
	}

	if err := self.DB.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func (self *DB) UnassignRoleFromOrganization(ctx context.Context, roleID, userID, teamID string) error {
	if err := self.DB.WithContext(ctx).
		Where("role_id = ? AND user_id = ? AND team_id = ?", roleID, userID, teamID).
		Delete(model.OrganizationRoleAssignment{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
