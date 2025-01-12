package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetNeutronVersion(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	resp := oapi.NeutronVersionResponse{
		Versions: []oapi.NeutronVersion{
			{
				Id:     "v2.15",
				Status: "stable",
				Links: []oapi.NeutronVersionLink{
					{
						Rel: "self",
					},
				},
			},
		},
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetNeutronNetworks(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronNetworkByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronQuotasByProjectID(ectx echo.Context, projectID string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronPorts(ectx echo.Context, params oapi.GetNeutronPortsParams) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronExtensions(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}
