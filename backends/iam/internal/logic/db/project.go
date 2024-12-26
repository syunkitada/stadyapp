package db

import (
	"context"
	"fmt"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
)

func (self *DB) FindProjects(ctx context.Context, input *db.FindProjectsInput) (items []model.Project, err error) {
	query := self.DB.WithContext(ctx).Model(model.Project{}).
		Select("id,name").
		Where("deleted = 0")

	if input.ID != 0 {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	if err = query.Scan(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to find projects: %v", err)
	}

	return items, nil
}

func (self *DB) AddProject(ctx context.Context, item *model.Project) (*model.Project, error) {
	if err := self.DB.WithContext(ctx).Model(model.Project{}).Save(item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (self *DB) DeleteProject(ctx context.Context, id uint64) (err error) {
	if err = self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Project{}).Error; err != nil {
		return err
	}

	return nil
}
