package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

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

	users, err := self.db.GetUsers(ctx, &db.GetUsersInput{
		ID: authContext.UserID,
	})
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}
	if len(users) > 1 {
		return nil, "", tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "user is duplicated"))
	}

	if authContext.DomainID == "" {
		return nil, "", tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "domain does not found"))
	}

	domain, err := self.db.GetDomain(ctx, &db.GetDomainsInput{
		ID: authContext.DomainID,
	})
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	if len(users) == 0 {
		_, err = self.db.CreateUser(ctx, &db.CreateUserInput{
			ID:          &authContext.UserID,
			Name:        authContext.UserID,
			LastLoginAt: time.Now(),
			DomainID:    domain.ID,
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

	tokenRoles := []string{}
	roles := []oapi.KeystoneTokenRole{}

	projectName := ""
	var project *oapi.KeystoneTokenProject
	if input.Auth.Scope != nil {
		dbProject, err := self.db.GetProject(ctx, &db.GetProjectsInput{
			Name: input.Auth.Scope.Project.Name,
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}

		project = &oapi.KeystoneTokenProject{
			Domain: oapi.KeystoneTokenDomain{
				Id:   domain.ID,
				Name: domain.Name,
			},
			Id:   dbProject.Name,
			Name: dbProject.Name,
		}
		projectName = dbProject.Name

		userProjectRoles, err := self.db.GetUserProjectRoles(ctx, &db.GetUserProjectRolesInput{
			UserID:    authContext.UserID,
			ProjectID: dbProject.ID,
		})
		if err != nil {
			return nil, "", tlog.Err(ctx, err)
		}

		roleSet := map[string]bool{}
		for _, role := range userProjectRoles {
			if role.RoleID == model.RoleIDGroup {
				roleSet[role.TeamRoleID] = true
			} else {
				roleSet[role.RoleID] = true
			}
		}

		for roleID := range roleSet {
			roles = append(roles, oapi.KeystoneTokenRole{
				Id:   roleID,
				Name: roleID,
			})

			tokenRoles = append(tokenRoles, roleID)
		}
	}

	traceID, err := tlog.GetTraceID(ctx)
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	token := oapi.KeystoneToken{
		AuditIds:  []string{traceID},
		Methods:   input.Auth.Identity.Methods,
		ExpiresAt: time.Now(),
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

	authData := iam_auth.AuthData{
		Domain:  domain.ID,
		User:    authContext.UserID,
		Project: projectName,
		Catalog: self.keystoneCatalogStr,
	}

	tokenRolesJson, err := json.Marshal(tokenRoles)
	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}
	authData.Roles = string(tokenRolesJson)

	tokenStr, err := self.iamAuth.NewToken(ctx, authData)

	if err != nil {
		return nil, "", tlog.Err(ctx, err)
	}

	return &token, tokenStr, nil
}
