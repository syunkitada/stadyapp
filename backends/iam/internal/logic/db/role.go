package db

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) FindRoles(ctx context.Context, input *db.FindRolesInput) ([]model.Role, error) {
	query := self.DB.WithContext(ctx).Model(model.Role{}).
		Select("id,name").
		Where("deleted = 0")

	if input.ID != 0 {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	items := []model.Role{}
	if err := query.Scan(&items).Error; err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to query.Scan")
	}

	return items, nil
}

func (self *DB) AddRole(ctx context.Context, item *model.Role) (*model.Role, error) {
	if err := self.DB.WithContext(ctx).Model(model.Role{}).Save(item).Error; err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.DB.WithContext.Save")
	}

	return item, nil
}

func (self *DB) DeleteRole(ctx context.Context, id uint64) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Role{}).Error; err != nil {
		return tlog.WrapError(ctx, err, "failed to self.DB.WithContext.Delete")
	}

	return nil
}
