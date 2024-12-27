package db

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) FindProjects(ctx context.Context, input *db.FindProjectsInput) ([]model.Project, error) {
	query := self.DB.WithContext(ctx).Model(model.Project{}).
		Select("id,name").
		Where("deleted = 0")

	return nil, echo.NewHTTPError(http.StatusNotFound, "Not Found")

	if input.ID != 0 {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	items := []model.Project{}
	if err := query.Scan(&items).Error; err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to query.Scan")
	}

	return items, nil
}

func (self *DB) AddProject(ctx context.Context, item *model.Project) (*model.Project, error) {
	if err := self.DB.WithContext(ctx).Model(model.Project{}).Save(item).Error; err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.DB.WithContext.Save")
	}

	return item, nil
}

func (self *DB) DeleteProject(ctx context.Context, id uint64) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Project{}).Error; err != nil {
		return tlog.WrapError(ctx, err, "failed to self.DB.WithContext.Delete")
	}

	return nil
}
