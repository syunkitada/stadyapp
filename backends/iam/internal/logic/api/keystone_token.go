package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
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

	domain, dbUser, err := self.LoginUser(ctx, authContext, authContext.UserID)
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	var dbProject *model.Project

	roleSet := map[string]bool{}
	roles := []oapi.KeystoneTokenRole{}

	if authContext.Inherit {
		dbProject, err = self.db.GetProject(ctx, &db.GetProjectsInput{
			ID: &authContext.ProjectID,
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}

		for _, roleID := range authContext.Roles {
			roleSet[roleID] = true
			roles = append(roles, oapi.KeystoneTokenRole{
				Id:   roleID,
				Name: roleID,
			})
		}

	} else if input.Auth.Scope != nil {
		dbProject, roleSet, err = self.GetProjectRoles(ctx, &GetProjectRolesInput{
			ProjectID:   input.Auth.Scope.Project.Id,
			ProjectName: input.Auth.Scope.Project.Name,
			UserID:      dbUser.ID,
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}

		for roleID := range roleSet {
			roles = append(roles, oapi.KeystoneTokenRole{
				Id:   roleID,
				Name: roleID,
			})
		}
	}

	var project *oapi.KeystoneTokenProject
	var projectID string

	if dbProject != nil {
		projectID = dbProject.ID

		project = &oapi.KeystoneTokenProject{
			Domain: oapi.KeystoneTokenDomain{
				Id:   domain.ID,
				Name: domain.Name,
			},
			Id:   dbProject.Name,
			Name: dbProject.Name,
		}
	}

	traceID, err := tlog.GetTraceID(ctx)
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	authData := iam_auth.AuthData{
		DomainID:  domain.ID,
		UserID:    dbUser.ID,
		ProjectID: projectID,
		Catalog:   self.keystoneCatalogStr,
		RoleSet:   roleSet,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	tokenStr, err := self.iamAuth.NewToken(ctx, authData)

	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	token := oapi.KeystoneToken{
		AuditIds:  []string{traceID},
		Methods:   input.Auth.Identity.Methods,
		ExpiresAt: authData.ExpiresAt,
		IssuedAt:  time.Now(),
		User: oapi.KeystoneTokenUser{
			Domain: oapi.KeystoneTokenDomain{
				Id:   domain.ID,
				Name: domain.Name,
			},
			Id:                authContext.UserID,
			Name:              authContext.UserID,
			PasswordExpiresAt: time.Now(),
		},
		Project: project,
		Roles:   roles,
		Catalog: self.keystoneCatalog,
	}

	return &token, tokenStr, nil
}
