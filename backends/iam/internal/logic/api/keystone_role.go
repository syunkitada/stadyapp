package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) CreateKeystoneRole(ctx context.Context, input *oapi.CreateKeystoneRoleInput) (*oapi.KeystoneRole, error) {
	project := oapi.KeystoneRole{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) GetKeystoneRoles(ctx context.Context, input *oapi.GetKeystoneRolesParams) ([]oapi.KeystoneRole, error) {
	projects := []oapi.KeystoneRole{
		{
			Id:   "project_id",
			Name: "project_name",
		},
	}
	return projects, nil
}

func (self *API) GetKeystoneRoleByID(ctx context.Context, id string) (*oapi.KeystoneRole, error) {
	project := oapi.KeystoneRole{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) DeleteKeystoneRole(ctx context.Context, id string) error {
	return nil
}
