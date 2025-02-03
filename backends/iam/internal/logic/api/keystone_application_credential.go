package api

import (
	"context"
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

func (self *API) LoginUser(ctx context.Context, authContext *iam_auth.AuthContext, userName string) (*model.Domain, *model.User, error) {
	if authContext.DomainID == "" {
		return nil, nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "domain does not found"))
	}

	domain, err := self.db.GetDomain(ctx, &db.GetDomainsInput{
		ID: authContext.DomainID,
	})
	if err != nil {
		return nil, nil, tlog.Err(ctx, err)
	}

	dbUsers, err := self.db.GetUsers(ctx, &db.GetUsersInput{
		Name: userName,
	})
	if err != nil {
		return nil, nil, tlog.Err(ctx, err)
	}
	if len(dbUsers) > 1 {
		return nil, nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "user is duplicated"))
	}

	var dbUser *model.User
	if len(dbUsers) == 0 {
		dbUser, err = self.db.CreateUser(ctx, &db.CreateUserInput{
			Name:        userName,
			DomainID:    authContext.DomainID,
			LastLoginAt: time.Now(),
		})
		if err != nil {
			return nil, nil, tlog.Err(ctx, err)
		}
	} else {
		err = self.db.UpdateUserByID(ctx, authContext.UserID, &db.UpdateUserByIDInput{
			LastLoginAt: time.Now(),
		})
		dbUser = &dbUsers[0]
	}

	return domain, dbUser, nil
}

type GetProjectRolesInput struct {
	UserID      string
	ProjectID   *string
	ProjectName *string
}

func (self *API) GetProjectRoles(ctx context.Context, input *GetProjectRolesInput) (*model.Project, map[string]bool, error) {
	dbProject, err := self.db.GetProject(ctx, &db.GetProjectsInput{
		ID:   input.ProjectID,
		Name: input.ProjectName,
	})
	if err != nil {
		return nil, nil, tlog.Err(ctx, err)
	}

	if input.UserID == "" {
		return nil, nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusInternalServerError, "user_id is required"))
	}

	userProjectRoles, err := self.db.GetUserProjectRoles(ctx, &db.GetUserProjectRolesInput{
		UserID:    input.UserID,
		ProjectID: dbProject.ID,
	})
	if err != nil {
		return nil, nil, tlog.Err(ctx, err)
	}

	roleSet := map[string]bool{}
	for _, role := range userProjectRoles {
		if role.RoleID == model.RoleIDGroup {
			roleSet[role.TeamRoleID] = true
		} else {
			roleSet[role.RoleID] = true
		}
	}

	fmt.Println("GetProjectRoles roleSet", roleSet)

	return dbProject, roleSet, nil
}

func (self *API) CreateKeystoneApplicationCredential(
	ctx context.Context, userID string, input *oapi.CreateKeystoneApplicationCredentialInput) (*oapi.KeystoneApplicationCredential, error) {
	authContext, err := iam_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	dbDomain, dbUser, err := self.LoginUser(ctx, authContext, input.ApplicationCredential.Name)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	dbProject, roleSet, err := self.GetProjectRoles(ctx, &GetProjectRolesInput{
		UserID:    authContext.UserID,
		ProjectID: &authContext.ProjectID,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	expiresAt := time.Now().Add(time.Hour * 24)
	if input.ApplicationCredential.ExpiresAt != nil {
		fmt.Println("input.ApplicationCredential.ExpiresAt", *input.ApplicationCredential.ExpiresAt)
		expiresAt, err = time.Parse("2006-01-02T15:04:05", *input.ApplicationCredential.ExpiresAt)
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}
		fmt.Println("input.ApplicationCredential.ExpiresAt", expiresAt)
	}

	authData := iam_auth.AuthData{
		DomainID:  dbDomain.ID,
		UserID:    dbUser.ID,
		ProjectID: dbProject.ID,
		Catalog:   self.keystoneCatalogStr,
		RoleSet:   roleSet,
		Inherit:   true,
		ExpiresAt: expiresAt,
	}

	tokenStr, err := self.iamAuth.NewToken(ctx, authData)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	roles := []oapi.KeystoneTokenRole{}
	for key := range roleSet {
		roles = append(roles, oapi.KeystoneTokenRole{
			Id:   key,
			Name: key,
		})
	}

	applicationCredential := oapi.KeystoneApplicationCredential{
		Id:        dbUser.ID,
		Name:      dbUser.Name,
		Secret:    tokenStr,
		ExpiresAt: expiresAt,
		Roles:     roles,
	}

	return &applicationCredential, nil
}
