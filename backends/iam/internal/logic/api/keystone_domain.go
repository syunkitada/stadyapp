package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) CreateKeystoneDomain(
	ctx context.Context, input *oapi.CreateKeystoneDomainInput) (*oapi.KeystoneDomain, error) {

	authContext, err := iam_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	dbDomain, err := self.db.CreateDomain(ctx, &db.CreateDomainInput{
		Name:        input.Domain.Name,
		Description: input.Domain.Description,
		Extra:       input.Domain.AdditionalProperties,
		OwnerUserID: &authContext.UserID,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domain, err := ConvertDBDomainToAPIDomain(ctx, dbDomain)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func (self *API) UpdateKeystoneDomainByID(
	ctx context.Context, id string, input *oapi.UpdateKeystoneDomainInput) (*oapi.KeystoneDomain, error) {
	err := self.db.UpdateDomainByID(ctx, id, &db.UpdateDomainByIDInput{
		Name:        input.Domain.Name,
		Description: input.Domain.Description,
		Extra:       input.Domain.AdditionalProperties,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domain, err := self.GetKeystoneDomainByID(ctx, id)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func (self *API) GetKeystoneDomains(
	ctx context.Context, input *oapi.GetKeystoneDomainsParams) ([]oapi.KeystoneDomain, error) {
	getDomainsInput := db.GetDomainsInput{}
	if input.Name != nil {
		getDomainsInput.Name = *input.Name
	}

	dbDomains, err := self.db.GetDomains(ctx, &getDomainsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	domains := []oapi.KeystoneDomain{}

	for i := range dbDomains {
		domain, err := ConvertDBDomainToAPIDomain(ctx, &dbDomains[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		domains = append(domains, *domain)
	}

	return domains, nil
}

func (self *API) GetKeystoneDomainByID(ctx context.Context, id string) (*oapi.KeystoneDomain, error) {
	dbDomains, err := self.db.GetDomains(ctx, &db.GetDomainsInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbDomains) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "domain not found"))
	}

	if len(dbDomains) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "domain is duplicated"))
	}

	domain, err := ConvertDBDomainToAPIDomain(ctx, &dbDomains[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return domain, nil
}

func ConvertDBDomainToAPIDomain(ctx context.Context, dbDomain *model.Domain) (*oapi.KeystoneDomain, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbDomain.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneDomain{
		Id:                   dbDomain.ID,
		Name:                 dbDomain.Name,
		Description:          dbDomain.Description,
		AdditionalProperties: additionalProperties,
	}, nil
}

func (self *API) DeleteKeystoneDomain(ctx context.Context, id string) error {
	err := self.db.DeleteDomainByID(ctx, id)
	if err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}
