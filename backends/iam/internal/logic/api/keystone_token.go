package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
)

func ParseKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return parsed, nil
}

func (self *API) CreateKeystoneToken(ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error) {
	authContext, err := iam_token_auth.GetAuthContext(ctx)

	domainID := "default"

	token := oapi.KeystoneToken{
		Token: oapi.KeystoneTokenData{
			AuditIds:  []string{"audit_id1", "audit_id2"},
			Methods:   []string{"password"},
			ExpiresAt: time.Now(),
			IssuedAt:  time.Now(),
			User: oapi.KeystoneTokenUser{
				Domain: oapi.KeystoneTokenDomain{
					Id:   domainID,
					Name: "domain_name",
				},
				Id:                authContext.UserID,
				Name:              authContext.UserID,
				PasswordExpiresAt: time.Now(),
			},
			Project: oapi.KeystoneTokenProject{
				Domain: oapi.KeystoneTokenDomain{
					Id:   "domain_id",
					Name: "domain_name",
				},
				Id:   "project_id",
				Name: "project_name",
			},
			Roles: []oapi.KeystoneTokenRole{
				{
					Id:   "role_id",
					Name: "role_name",
				},
			},
			Catalog: []oapi.KeystoneCatalog{
				{
					Id:   "catalog_id",
					Name: "keystone",
					Type: "identity",
					Endpoints: []oapi.KeystoneEndpoint{
						{
							Id:        "endpoint_id",
							Interface: "public",
							Region:    "region1",
							Url:       "http://localhost:10080/api/iam/keystone/v3",
						},
					},
				},
			},
		},
	}

	authData := iam_token_auth.AuthData{
		Domain:  domainID,
		User:    authContext.UserID,
		Project: "project",
		Roles:   "roles",
		Catalog: "gatalog",
	}
	tokenStr, err := self.iamTokenAuth.NewToken(ctx, authData)
	if err != nil {
		return nil, "", err
	}

	return &token, tokenStr, nil
}
