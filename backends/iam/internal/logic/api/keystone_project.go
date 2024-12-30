package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) CreateKeystoneProject(ctx context.Context, input *oapi.CreateKeystoneProjectInput) (*oapi.KeystoneProject, error) {
	project := oapi.KeystoneProject{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) GetKeystoneProjects(ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error) {
	projects := []oapi.KeystoneProject{
		{
			Id:   "project_id",
			Name: "project_name",
		},
	}
	return projects, nil
}

func (self *API) GetKeystoneProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error) {
	project := oapi.KeystoneProject{
		Id:   "project_id",
		Name: "project_name",
	}
	return &project, nil
}

func (self *API) DeleteKeystoneProject(ctx context.Context, id string) error {
	return nil
}
