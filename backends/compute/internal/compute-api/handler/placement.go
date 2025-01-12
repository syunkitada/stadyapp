package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetPlacementVersion(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	resp := oapi.PlacementVersionResponse{
		Versions: []oapi.PlacementVersion{
			{
				Id:         "v1.0",
				MaxVersion: "1.39",
				MinVersion: "1.0",
				Status:     "CURRENT",
				Links: []oapi.PlacementVersionLink{
					{
						Rel: "self",
					},
				},
			},
		},
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetPlacementResourceProviders(ectx echo.Context, params oapi.GetPlacementResourceProvidersParams) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Placement)
}

func (self *Handler) GetPlacementResourceProviderAllocationsByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Placement)
}
