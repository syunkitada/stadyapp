package db

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) GetOrganization(ctx context.Context, input *db.GetOrganizationsInput) (*model.Organization, error) {
	dbOrganizations, err := self.GetOrganizations(ctx, input)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbOrganizations) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "organization does not found"))
	}

	if len(dbOrganizations) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "organization is duplicated"))
	}

	return &dbOrganizations[0], nil
}

func (self *DB) GetOrganizations(ctx context.Context, input *db.GetOrganizationsInput) ([]model.Organization, error) {
	query := self.DB.WithContext(ctx).Model(model.Organization{}).
		Select("id,name,description,extra")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	organaziations := []model.Organization{}
	if err := query.Scan(&organaziations).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return organaziations, nil
}

func (self *DB) CreateOrganization(ctx context.Context, input *db.CreateOrganizationInput) (*model.Organization, error) {
	bytes, err := json.Marshal(input.Extra)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	organization := model.Organization{
		Name:     input.Name,
		Extra:    string(bytes),
		DomainID: input.DomainID,
	}

	if input.ID == nil {
		organization.ID = uuid.New().String()
	} else {
		organization.ID = *input.ID
	}

	if input.Description != nil {
		organization.Description = *input.Description
	}

	if err := self.Transact(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&organization).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		roleAssignment := model.OrganizationRoleAssignment{
			RoleID:         model.RoleIDManager,
			UserID:         &input.OwnerUserID,
			OrganizationID: organization.ID,
		}
		if err := tx.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		return nil
	}); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &organization, nil
}

func (self *DB) UpdateOrganizationByID(ctx context.Context, id string, input *db.UpdateOrganizationByIDInput) error {
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
		if err := self.DB.WithContext(ctx).Model(model.Organization{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteOrganizationByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Organization{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
