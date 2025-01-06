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

func (self *DB) GetTeam(ctx context.Context, input *db.GetTeamsInput) (*model.Team, error) {
	dbTeams, err := self.GetTeams(ctx, input)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbTeams) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "team does not found"))
	}

	if len(dbTeams) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "team is duplicated"))
	}

	return &dbTeams[0], nil
}

func (self *DB) GetTeams(ctx context.Context, input *db.GetTeamsInput) ([]model.Team, error) {
	query := self.DB.WithContext(ctx).Model(model.Team{}).
		Select("id,name,description,extra")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	teams := []model.Team{}
	if err := query.Scan(&teams).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return teams, nil
}

func (self *DB) CreateTeam(ctx context.Context, input *db.CreateTeamInput) (*model.Team, error) {
	bytes, err := json.Marshal(input.Extra)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to json.Marshal")
	}

	team := model.Team{
		Name:     input.Name,
		Extra:    string(bytes),
		DomainID: input.DomainID,
	}

	if input.ID == nil {
		team.ID = uuid.New().String()
	} else {
		team.ID = *input.ID
	}

	if input.Description != nil {
		team.Description = *input.Description
	}

	if err := self.Transact(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&team).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		roleAssignment := model.TeamRoleAssignment{
			RoleID: model.RoleIDManager,
			UserID: input.OwnerUserID,
			TeamID: team.ID,
		}
		if err := tx.WithContext(ctx).Create(&roleAssignment).Error; err != nil {
			return tlog.Err(ctx, err)
		}

		return nil
	}); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &team, nil
}

func (self *DB) UpdateTeamByID(ctx context.Context, id string, input *db.UpdateTeamByIDInput) error {
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
		if err := self.DB.WithContext(ctx).Model(model.Team{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteTeamByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.Team{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
