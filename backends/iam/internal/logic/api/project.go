package api

import (
	"context"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) FindProjects(ctx context.Context, params oapi.FindProjectsParams) (items []oapi.Project, err error) {
	dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{})

	slog.InfoContext(ctx, "FindProjects", "dbProjects", dbProjects)

	if err != nil {
		return nil, err
	}

	for _, dbProject := range dbProjects {
		items = append(items, oapi.Project{
			Id:   dbProject.ID,
			Name: dbProject.Name,
		})
	}

	return items, nil
}

func (self *API) FindProjectByID(ctx context.Context, id uint64) (item oapi.Project, err error) {
	dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{ID: id})

	if err != nil {
		return item, err
	}

	if len(dbProjects) > 1 {
		return item, echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("multiple items found: id=%d", id))
	}

	if len(dbProjects) == 0 {
		return item, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("item not found: id=%d", id))
	}

	dbProject := dbProjects[0]
	item = oapi.Project{
		Id:   dbProject.ID,
		Name: dbProject.Name,
	}

	return item, nil
}

func (self *API) AddProject(ctx context.Context, item *oapi.NewProject) error {
	if _, err := self.db.AddProject(ctx, &model.Project{
		Name: item.Name,
	}); err != nil {
		return err
	}

	return nil
}

func (self *API) DeleteProject(ctx context.Context, id uint64) error {
	if err := self.db.DeleteProject(ctx, id); err != nil {
		return err
	}

	return nil
}
