package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/tlog"
)

func (self *Handler) FindProjects(ectx echo.Context, params oapi.FindProjectsParams) error {
	ctx := tlog.WithEchoContext(ectx)

	items, err := self.api.FindProjects(ctx, params)

	if err != nil {
		return fmt.Errorf("self.api.FindProjects: error=%w", err)
	}

	if err := ectx.JSON(http.StatusOK, items); err != nil {
		return fmt.Errorf("ectx.JSON: error=%w", err)
	}

	return nil
}

func (self *Handler) AddProject(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var newProject oapi.NewProject

	err := ectx.Bind(&newProject)

	if err != nil {
		return sendHandlerError(ectx, http.StatusBadRequest, "Invalid format for NewProject")
	}

	if err := self.api.AddProject(ctx, &newProject); err != nil {
		return fmt.Errorf("self.api.AddProject: error=%w", err)
	}

	return nil
}

func (self *Handler) FindProjectByID(ectx echo.Context, itemId uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	item, err := self.api.FindProjectByID(ctx, itemId)

	if err != nil {
		return fmt.Errorf("self.api.FindProjectByID: error=%w", err)
	}

	if err := ectx.JSON(http.StatusOK, item); err != nil {
		return fmt.Errorf("ectx.JSON: error=%w", err)
	}

	return nil
}

func (self *Handler) DeleteProject(ectx echo.Context, id uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	err := self.api.DeleteProject(ctx, id)

	if err != nil {
		return fmt.Errorf("self.api.DeleteProject: error=%w", err)
	}

	if err := ectx.NoContent(http.StatusNoContent); err != nil {
		return fmt.Errorf("ectx.NoContent: error=%w", err)
	}

	return nil
}
