package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_middleware"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) CreateToken(ectx echo.Context) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

	var newToken oapi.NewToken

	if err := ectx.Bind(&newToken); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	// if err := self.api.AddToken(ctx, &newToken); err != nil {
	// 	return tlog.BindEchoError(ctx, ectx, err)
	// }

	return nil
}
