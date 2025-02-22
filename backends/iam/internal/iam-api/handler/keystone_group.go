package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneGroups(ectx echo.Context, input oapi.GetKeystoneGroupsParams) error {
	ctx := iam_auth.WithEchoContext(ectx)

	groups, err := self.api.GetKeystoneGroups(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneGroupsResponse{
		Groups: groups,
	}
	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneGroupByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	group, err := self.api.GetKeystoneGroupByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneGroupResponse{
		Group: *group,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}
