package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) FindProjects(ectx echo.Context, params oapi.FindProjectsParams) error {
	ctx := tlog.WithEchoContext(ectx)

	items, err := self.api.FindProjects(ctx, params)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, items)
}

func (self *Handler) AddProject(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var newProject oapi.NewProject

	if err := ectx.Bind(&newProject); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	if err := self.api.AddProject(ctx, &newProject); err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return nil
}

func (self *Handler) FindProjectByID(ectx echo.Context, itemID uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	item, err := self.api.FindProjectByID(ctx, itemID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, item)
}

func (self *Handler) DeleteProject(ectx echo.Context, id uint64) error {
	ctx := tlog.WithEchoContext(ectx)

	err := self.api.DeleteProject(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
