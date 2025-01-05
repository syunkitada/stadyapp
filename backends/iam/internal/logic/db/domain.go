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

func (self *DB) GetDomainByID(ctx context.Context, id string) (*model.Domain, error) {
	dbDomains, err := self.GetDomains(ctx, &db.GetDomainsInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbDomains) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "domain does not found"))
	}

	if len(dbDomains) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "domain is duplicated"))
	}

	return &dbDomains[0], nil
}

func (self *DB) GetDomains(ctx context.Context, input *db.GetDomainsInput) ([]model.Domain, error) {
	query := self.DB.WithContext(ctx).Model(model.Domain{}).
		Select("id,name,description,extra")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	domains := []model.Domain{}
	if err := query.Scan(&domains).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domains, nil
}

func (self *DB) CreateDomain(ctx context.Context, input *db.CreateDomainInput) (*model.Domain, error) {
	bytes, err := json.Marshal(input.Extra)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	domain := model.Domain{
		Name:  input.Name,
		Extra: string(bytes),
	}

	if input.ID == nil {
		domain.ID = uuid.New().String()
	} else {
		domain.ID = *input.ID
	}

	if input.Description != nil {
		domain.Description = *input.Description
	}

	if err := self.Transact(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&domain).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		if input.OwnerUserID != nil {
			roleAssignment := model.DomainRoleAssignment{
				RoleID:   model.RoleIDManager,
				UserID:   input.OwnerUserID,
				DomainID: domain.ID,
			}
			if err := tx.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
				return tlog.Err(ctx, err)
			}
		}

		return nil
	}); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &domain, nil
}

func (self *DB) UpdateDomainByID(ctx context.Context, id string, input *db.UpdateDomainByIDInput) error {
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
		if err := self.DB.WithContext(ctx).Model(model.Domain{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteDomainByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Domain{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
