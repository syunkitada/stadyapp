package handler

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneVersion(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	var updated = time.Date(2020, 4, 7, 0, 0, 0, 0, time.Local)

	version := oapi.KeystoneVersion{
		Version: oapi.KeystoneVersionData{
			Id:      "v3.14",
			Status:  "stable",
			Updated: updated,
			Links: []oapi.KeystoneVersionLink{
				{
					Rel:  "self",
					Href: "http://localhost:5000/v3/",
				},
			},
			MediaTypes: []oapi.KeystoneVersionMediaType{
				{
					Base: "application/json",
					Type: "application/vnd.openstack.identity-v3+json",
				},
			},
		},
	}

	return tlog.BindEchoOK(ctx, ectx, version)
}
