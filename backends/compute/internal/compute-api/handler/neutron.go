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

func (self *Handler) GetNeutronSubnets(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronSecurityGroups(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronPorts(ectx echo.Context, params oapi.GetNeutronPortsParams) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronPortByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) UpdateNeutronPortByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) DeleteNeutronPortByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronAgents(ectx echo.Context, params oapi.GetNeutronAgentsParams) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronExtensions(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) CreateNeutronPort(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) CreateNeutronNetwork(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) DeleteNeutronNetworkByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) GetNeutronSubnetByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) DeleteNeutronSubnetByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}

func (self *Handler) CreateNeutronSubnet(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Neutron)
}
