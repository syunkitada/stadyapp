package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) CreateKeystoneProject(ctx context.Context, input *oapi.CreateKeystoneProjectInput) (*oapi.KeystoneProject, error) {
	dbProject, err := self.db.CreateProject(ctx, &db.CreateProjectInput{
		Name: input.Project.Name,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	project := oapi.KeystoneProject{
		Id:   dbProject.ID,
		Name: dbProject.Name,
	}

	return &project, nil
}

// func (self *API) GetKeystoneUserProjects(ctx context.Context, input *oapi.GetKeystoneUserProjectsParams) ([]oapi.KeystoneProject, error) {
// 	authContext, err := iam_token_auth.GetAuthContext(ctx)
// 	if err != nil {
// 		return nil, tlog.Err(ctx, err)
// 	}
//
// 	fmt.Println("context", authContext.ProjectID)
//
// 	getProjectsInput := db.GetProjectsInput{}
// 	if input.Name != nil {
// 		getProjectsInput.Name = *input.Name
// 	}
// 	dbProjects, err := self.db.GetProjects(ctx, &getProjectsInput)
// 	if err != nil {
// 		return nil, tlog.Err(ctx, err)
// 	}
//
// 	projects := []oapi.KeystoneProject{}
// 	for _, project := range dbProjects {
// 		projects = append(projects, oapi.KeystoneProject{
// 			Id:   project.ID,
// 			Name: project.Name,
// 		})
// 	}
// 	return projects, nil
// }

func (self *API) GetKeystoneProjects(ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error) {
	authContext, err := iam_token_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	fmt.Println("context", authContext.ProjectID)

	getProjectsInput := db.GetProjectsInput{}
	if input.Name != nil {
		getProjectsInput.Name = *input.Name
	}
	dbProjects, err := self.db.GetProjects(ctx, &getProjectsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}
	for _, project := range dbProjects {
		projects = append(projects, oapi.KeystoneProject{
			Id:   project.ID,
			Name: project.Name,
		})
	}
	return projects, nil
}

func (self *API) GetKeystoneUserProjects(ctx context.Context, userID string) ([]oapi.KeystoneProject, error) {
	authContext, err := iam_token_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if authContext.UserID != userID {
		err := echo.NewHTTPError(http.StatusForbidden, "forbidden")
		return nil, tlog.Err(ctx, err)
	}

	fmt.Println("context", authContext.ProjectID)

	getProjectsInput := db.GetProjectsInput{}
	dbProjects, err := self.db.GetProjects(ctx, &getProjectsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}
	for _, project := range dbProjects {
		projects = append(projects, oapi.KeystoneProject{
			Id:   project.ID,
			Name: project.Name,
		})
	}
	return projects, nil
}

func (self *API) GetKeystoneProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error) {
	dbProjects, err := self.db.GetProjects(ctx, &db.GetProjectsInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbProjects) == 0 {
		err := echo.NewHTTPError(http.StatusNotFound, "project not found")
		return nil, tlog.Err(ctx, err)
	}

	if len(dbProjects) > 1 {
		err := echo.NewHTTPError(http.StatusConflict, "project is duplicated")
		return nil, tlog.Err(ctx, err)
	}

	dbProject := dbProjects[0]

	project := oapi.KeystoneProject{
		Id:   dbProject.ID,
		Name: dbProject.Name,
	}
	return &project, nil
}

func (self *API) DeleteKeystoneProject(ctx context.Context, id string) error {
	err := self.db.DeleteProjectByID(ctx, id)
	if err != nil {
		return tlog.Err(ctx, err)
	}
	return nil
}
