package handler

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/echo_middleware"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

// GetKeystoneVersion returns keystone version as below.
// {"version": {"id": "v3.14", "status": "stable", "updated": "2020-04-07T00:00:00Z", "links": [{"rel": "self", "href": "http://localhost:5000/v3/"}], "media-types": [{"base": "application/json", "type": "application/vnd.openstack.identity-v3+json"}]}}
func (self *Handler) GetKeystoneVersion(ectx echo.Context) error {
	ctx := echo_middleware.WithAuthEchoContext(ectx)

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
