package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneProjects(ectx echo.Context, input oapi.GetKeystoneProjectsParams) error {
	ctx := iam_token_auth.WithEchoContext(ectx)

	projects, err := self.api.GetKeystoneProjects(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneProjectsResponse{
		Projects: projects,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneUserProjectsByUserID(ectx echo.Context, userID string) error {
	ctx := iam_token_auth.WithEchoContext(ectx)

	projects, err := self.api.GetKeystoneUserProjects(ctx, userID)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneProjectsResponse{
		Projects: projects,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneProjectByID(ectx echo.Context, id string) error {
	ctx := iam_token_auth.WithEchoContext(ectx)

	project, err := self.api.GetKeystoneProjectByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneProjectResponse{
		Project: *project,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) CreateKeystoneProject(ectx echo.Context) error {
	ctx := iam_token_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneProjectInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	project, err := self.api.CreateKeystoneProject(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneProjectResponse{
		Project: *project,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) DeleteKeystoneProjectByID(ectx echo.Context, id string) error {
	ctx := iam_token_auth.WithEchoContext(ectx)

	err := self.api.DeleteKeystoneProject(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
