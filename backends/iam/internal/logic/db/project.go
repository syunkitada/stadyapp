package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) GetProjects(ctx context.Context, input *db.GetProjectsInput) ([]model.Project, error) {
	query := self.DB.WithContext(ctx).Model(model.Project{}).
		Select("id,name")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	projects := []model.Project{}
	if err := query.Scan(&projects).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return projects, nil
}

func (self *DB) CreateProject(ctx context.Context, input *db.CreateProjectInput) (*model.Project, error) {
	bytes, err := json.Marshal(input.Properties)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	project := model.Project{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Extra: string(bytes),
	}
	if err := self.DB.WithContext(ctx).Create(&project).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &project, nil
}

func (self *DB) UpdateProject(ctx context.Context, input *db.UpdateProjectInput) error {
	data := map[string]interface{}{}
	if len(data) > 0 {
		if err := self.DB.WithContext(ctx).Model(model.Project{}).Updates(data).Error; err != nil {
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
