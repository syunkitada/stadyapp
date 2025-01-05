package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) CreateKeystoneUser(
	ctx context.Context, input *oapi.CreateKeystoneUserInput) (*oapi.KeystoneUser, error) {
	dbUser, err := self.db.CreateUser(ctx, &db.CreateUserInput{
		Name: input.User.Name,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	user, err := ConvertDBUserToAPIUser(ctx, dbUser)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return user, nil
}

func (self *API) GetKeystoneUsers(
	ctx context.Context, input *oapi.GetKeystoneUsersParams) ([]oapi.KeystoneUser, error) {
	getUsersInput := db.GetUsersInput{}
	if input.Name != nil {
		getUsersInput.Name = *input.Name
	}

	dbUsers, err := self.db.GetUsers(ctx, &getUsersInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	users := []oapi.KeystoneUser{}

	for i := range dbUsers {
		user, err := ConvertDBUserToAPIUser(ctx, &dbUsers[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		users = append(users, *user)
	}

	return users, nil
}

func (self *API) GetKeystoneUserByID(ctx context.Context, id string) (*oapi.KeystoneUser, error) {
	dbUsers, err := self.db.GetUsers(ctx, &db.GetUsersInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbUsers) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "user not found"))
	}

	if len(dbUsers) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "user is duplicated"))
	}

	user, err := ConvertDBUserToAPIUser(ctx, &dbUsers[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return user, nil
}

func ConvertDBUserToAPIUser(ctx context.Context, dbUser *model.User) (*oapi.KeystoneUser, error) {
	return &oapi.KeystoneUser{
		Id:   dbUser.ID,
		Name: dbUser.Name,
	}, nil
}

func (self *API) DeleteKeystoneUser(ctx context.Context, id string) error {
	err := self.db.DeleteUserByID(ctx, id)
	if err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
