package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
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

func (self *API) CreateKeystoneToken(
	ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error) {
	authContext, err := iam_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	users, err := self.db.GetUsers(ctx, &db.GetUsersInput{
		ID: authContext.UserID,
	})
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}
	if len(users) > 1 {
		return nil, "", tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "user is duplicated"))
	}

	if len(users) == 0 {
		_, err = self.db.CreateUser(ctx, &db.CreateUserInput{
			ID:          &authContext.UserID,
			Name:        authContext.UserID,
			LastLoginAt: time.Now(),
			DomainID:    "default",
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}
	} else {
		err = self.db.UpdateUserByID(ctx, authContext.UserID, &db.UpdateUserByIDInput{
			LastLoginAt: time.Now(),
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}
	}

	fmt.Println("DEBUG Scope", input.Auth.Scope.Project)

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

	authData := iam_auth.AuthData{
		Domain:  domainID,
		User:    authContext.UserID,
		Project: "project",
		Roles:   "role1,role2",
		Catalog: "{catalog}",
	}
	tokenStr, err := self.iamAuth.NewToken(ctx, authData)

	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	return &token, tokenStr, nil
}
