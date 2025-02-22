package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *Handler) GetKeystoneDomains(ectx echo.Context, input oapi.GetKeystoneDomainsParams) error {
	ctx := iam_auth.WithEchoContext(ectx)

	domains, err := self.api.GetKeystoneDomains(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneDomainsResponse{
		Domains: domains,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) GetKeystoneDomainByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	domain, err := self.api.GetKeystoneDomainByID(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneDomainResponse{
		Domain: *domain,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) CreateKeystoneDomain(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.CreateKeystoneDomainInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	domain, err := self.api.CreateKeystoneDomain(ctx, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneDomainResponse{
		Domain: *domain,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) UpdateKeystoneDomainByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	var input oapi.UpdateKeystoneDomainInput
	if err := ectx.Bind(&input); err != nil {
		return tlog.BindEchoBadRequest(ctx, ectx, err)
	}

	domain, err := self.api.UpdateKeystoneDomainByID(ctx, id, &input)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	resp := oapi.KeystoneDomainResponse{
		Domain: *domain,
	}

	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) DeleteKeystoneDomainByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)

	err := self.api.DeleteKeystoneDomain(ctx, id)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	return tlog.BindEchoNoContent(ctx, ectx)
}
