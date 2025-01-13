package db

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) GetProject(ctx context.Context, input *db.GetProjectsInput) (*model.Project, error) {
	dbProjects, err := self.GetProjects(ctx, input)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbProjects) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "project does not found"))
	}

	if len(dbProjects) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "project is duplicated"))
	}

	return &dbProjects[0], nil
}

func (self *DB) GetProjects(ctx context.Context, input *db.GetProjectsInput) ([]model.Project, error) {
	query := self.DB.WithContext(ctx).Model(model.Project{}).
		Select("id,name,description,extra")

	if input.ID != nil {
		query.Where("id = ?", *input.ID)
	}

	if input.Name != nil {
		query.Where("name = ?", *input.Name)
	}

	projects := []model.Project{}
	if err := query.Scan(&projects).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return projects, nil
}

func (self *DB) CreateProject(ctx context.Context, input *db.CreateProjectInput) (*model.Project, error) {
	bytes, err := json.Marshal(input.Extra)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	project := model.Project{
		Name:           input.Name,
		Extra:          string(bytes),
		DomainID:       input.DomainID,
		OrganizationID: input.OrganizationID,
	}

	if input.ID == nil {
		project.ID = uuid.New().String()
	} else {
		project.ID = *input.ID
	}

	if input.Description != nil {
		project.Description = *input.Description
	}

	if err := self.Transact(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&project).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		roleAssignment := model.ProjectRoleAssignment{
			RoleID:    model.RoleIDManager,
			UserID:    &input.OwnerUserID,
			ProjectID: project.ID,
		}
		if err := tx.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		return nil
	}); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &project, nil
}

func (self *DB) UpdateProjectByID(ctx context.Context, id string, input *db.UpdateProjectByIDInput) error {
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
		if err := self.DB.WithContext(ctx).Model(model.Project{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteProjectByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Project{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
