package db

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	dbUsers, err := self.GetUsers(ctx, &db.GetUsersInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbUsers) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "user does not found"))
	}

	if len(dbUsers) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "user is duplicated"))
	}

	return &dbUsers[0], nil
}

func (self *DB) GetUsers(ctx context.Context, input *db.GetUsersInput) ([]model.User, error) {
	query := self.DB.WithContext(ctx).Model(model.User{}).
		Select("id,name")

	if input.ID != "" {
		query.Where("id = ?", input.ID)
	}

	if input.Name != "" {
		query.Where("name = ?", input.Name)
	}

	users := []model.User{}
	if err := query.Scan(&users).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return users, nil
}

func (self *DB) CreateUser(ctx context.Context, input *db.CreateUserInput) (*model.User, error) {
	user := model.User{
		Name:        input.Name,
		DomainID:    input.DomainID,
		LastLoginAt: input.LastLoginAt,
	}

	if input.ID == nil {
		user.ID = uuid.New().String()
	} else {
		user.ID = *input.ID
	}

	if err := self.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &user, nil
}

func (self *DB) UpdateUserByID(ctx context.Context, id string, input *db.UpdateUserByIDInput) error {
	data := map[string]interface{}{}
	if input.Name != nil {
		data["name"] = *input.Name
	}

	if len(data) > 0 {
		if err := self.DB.WithContext(ctx).Model(model.User{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) DeleteUserByID(ctx context.Context, id string) error {
	if err := self.DB.WithContext(ctx).Where("id = ?", id).Delete(model.User{}).Error; err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
