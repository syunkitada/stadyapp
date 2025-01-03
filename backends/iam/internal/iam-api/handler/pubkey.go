package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetPubkeys(ectx echo.Context, input oapi.GetPubkeysParams) error {
	ctx := tlog.WithEchoContext(ectx)

	pubkeys, err := self.api.GetPubkeys(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.Pubkeys{
		Pubkeys: pubkeys,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}
