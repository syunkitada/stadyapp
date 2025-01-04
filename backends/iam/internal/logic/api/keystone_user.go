package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) CreateKeystoneUser(ctx context.Context, input *oapi.CreateKeystoneUserInput) (*oapi.KeystoneUser, error) {
	project := oapi.KeystoneUser{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) GetKeystoneUsers(ctx context.Context, input *oapi.GetKeystoneUsersParams) ([]oapi.KeystoneUser, error) {
	projects := []oapi.KeystoneUser{
		{
			Id:   "project_id",
			Name: "project_name",
		},
	}
	return projects, nil
}

func (self *API) GetKeystoneUserByID(ctx context.Context, id string) (*oapi.KeystoneUser, error) {
	project := oapi.KeystoneUser{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) DeleteKeystoneUser(ctx context.Context, id string) error {
	return nil
}
