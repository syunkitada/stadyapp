package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) GetPubkeys(ctx context.Context, params *oapi.GetPubkeysParams) ([]oapi.Pubkey, error) {
	pubkeys := []oapi.Pubkey{
		{
			Key: "pubkey",
		},
	}

	return pubkeys, nil
}
