package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_middleware"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneProjects(ectx echo.Context, input oapi.GetKeystoneProjectsParams) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	projects, err := self.api.GetKeystoneProjects(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneProjects{
		Projects: projects,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneProjectByID(ectx echo.Context, id string) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	project, err := self.api.GetKeystoneProjectByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, project)
}

func (self *Handler) CreateKeystoneProject(ectx echo.Context) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	var input oapi.CreateKeystoneProjectInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	project, err := self.api.CreateKeystoneProject(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoOK(ctx, ectx, project)
}

func (self *Handler) DeleteKeystoneProjectByID(ectx echo.Context, id string) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	err := self.api.DeleteKeystoneProject(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
