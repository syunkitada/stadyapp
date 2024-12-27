package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) FindProjects(ctx context.Context, params oapi.FindProjectsParams) ([]oapi.Project, error) {
	dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{})

	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.db.FindProjects")
	}

	projects := []oapi.Project{}
	for _, dbProject := range dbProjects {
		projects = append(projects, oapi.Project{
			Id:   dbProject.ID,
			Name: dbProject.Name,
		})
	}

	return projects, nil
}

func (self *API) FindProjectByID(ctx context.Context, id uint64) (*oapi.Project, error) {
	dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{ID: id})

	if err != nil {
		return nil, tlog.WrapError(ctx, err, "failed to self.db.FindProjects")
	}

	if len(dbProjects) > 1 {
		return nil, tlog.WrapEchoConflictError(ctx, "multiple projects found")
	}

	if len(dbProjects) == 0 {
		return nil, tlog.WrapEchoNotFoundError(ctx, "project does not found")
	}

	dbProject := dbProjects[0]
	project := &oapi.Project{
		Id:   dbProject.ID,
		Name: dbProject.Name,
	}

	return project, nil
}

func (self *API) AddProject(ctx context.Context, project *oapi.NewProject) error {
	if _, err := self.db.AddProject(ctx, &model.Project{
		Name: project.Name,
	}); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.db.AddProject")
	}

	return nil
}

func (self *API) DeleteProject(ctx context.Context, id uint64) error {
	if err := self.db.DeleteProject(ctx, id); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.db.DeleteProject")
	}

	return nil
}
