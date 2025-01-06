package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetGlanceVersion(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	resp := oapi.GlanceVersionResponse{
		Versions: []oapi.GlanceVersion{
			{
				Id:     "v2.15",
				Status: "stable",
				Links: []oapi.GlanceVersionLink{
					{
						Rel: "self",
					},
				},
			},
		},
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}
