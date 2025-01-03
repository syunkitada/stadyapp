package api

import (
	"context"
	"fmt"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) CreateKeystoneUser(ctx context.Context, input *oapi.CreateKeystoneUserInput) (*oapi.KeystoneUser, error) {
	project := oapi.KeystoneUser{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) GetKeystoneUsers(ctx context.Context, input *oapi.GetKeystoneUsersParams) ([]oapi.KeystoneUser, error) {
	authContext, err := iam_token_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}
	fmt.Println("context", authContext.ProjectID)

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
