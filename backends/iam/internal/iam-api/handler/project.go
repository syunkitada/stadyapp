package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) FindProjects(ectx echo.Context, params oapi.FindProjectsParams) error {
	ctx := tlog.WithEchoContext(ectx)

	items, err := self.api.FindProjects(ctx, params)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.FindProjects")
	}

	if err := ectx.JSON(http.StatusOK, items); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.JSON")
	}

	return nil
}

func (self *Handler) AddProject(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var newProject oapi.NewProject

	err := ectx.Bind(&newProject)

	if err != nil {
		err = sendHandlerError(ectx, http.StatusBadRequest, "Invalid format for NewProject")
		return tlog.WrapError(ctx, err, "failed to ectx.Bind")
	}

	if err := self.api.AddProject(ctx, &newProject); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.AddProject")
	}

	return nil
}

func (self *Handler) FindProjectByID(ectx echo.Context, itemID uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	item, err := self.api.FindProjectByID(ctx, itemID)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.FindProjectByID")
	}

	if err := ectx.JSON(http.StatusOK, item); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.JSON")
	}

	return nil
}

func (self *Handler) DeleteProject(ectx echo.Context, id uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	err := self.api.DeleteProject(ctx, id)

	if err != nil {
		return tlog.WrapError(ctx, err, "failed to self.api.DeleteProject")
	}

	if err := ectx.NoContent(http.StatusNoContent); err != nil {
		return tlog.WrapError(ctx, err, "failed to ectx.NoContent")
	}

	return nil
}
