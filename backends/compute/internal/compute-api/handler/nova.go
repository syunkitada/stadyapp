package handler

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetNovaVersion(ectx echo.Context) error {
	ctx := tlog.WithEchoContext(ectx)

	resp := oapi.NovaVersionResponse{
		Version: oapi.NovaVersion{
			Id:         "v2.1",
			Status:     "stable",
			Version:    "2.96",
			MinVersion: "2.1",
			Updated:    "2013-07-23T11:33:21Z",
			Links: []oapi.NovaVersionLink{
				{
					Rel: "self",
				},
			},
		},
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetNovaServerByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) GetNovaServersDetail(ectx echo.Context, params oapi.GetNovaServersDetailParams) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) GetNovaFlavorsDetail(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) GetNovaFlavorByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) CreateNovaServer(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) DeleteNovaServerByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) GetNovaServices(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) CreateNovaExternalEvents(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) CreateNovaFlavor(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}

func (self *Handler) ActionNovaServer(ectx echo.Context, id string) error {
	time.Sleep(3 * time.Second)
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Nova)
}
