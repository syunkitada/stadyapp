package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

type IAPI interface {
	IAPIKeystoneToken
	IAPIKeystoneDomain
	IAPIKeystoneProject
	IAPIKeystoneUser
	IAPIKeystoneRole
	IAPIOther
}

type IAPIOther interface {
	GetPubkeys(ctx context.Context, input *oapi.GetPubkeysParams) ([]oapi.Pubkey, error)
}

type IAPIKeystoneToken interface {
	CreateKeystoneToken(ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error)
}

type IAPIKeystoneDomain interface {
	CreateKeystoneDomain(ctx context.Context, input *oapi.CreateKeystoneDomainInput) (*oapi.KeystoneDomain, error)
	UpdateKeystoneDomainByID(
		ctx context.Context, id string, input *oapi.UpdateKeystoneDomainInput) (*oapi.KeystoneDomain, error)
	GetKeystoneDomains(ctx context.Context, input *oapi.GetKeystoneDomainsParams) ([]oapi.KeystoneDomain, error)
	GetKeystoneDomainByID(ctx context.Context, id string) (*oapi.KeystoneDomain, error)
	DeleteKeystoneDomain(ctx context.Context, id string) error
}

type IAPIKeystoneProject interface {
	CreateKeystoneProject(ctx context.Context, input *oapi.CreateKeystoneProjectInput) (*oapi.KeystoneProject, error)
	UpdateKeystoneProjectByID(
		ctx context.Context, id string, input *oapi.UpdateKeystoneProjectInput) (*oapi.KeystoneProject, error)
	GetKeystoneProjects(ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error)
	GetKeystoneUserProjects(ctx context.Context, userID string) ([]oapi.KeystoneProject, error)
	GetKeystoneProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error)
	DeleteKeystoneProject(ctx context.Context, id string) error
}

type IAPIKeystoneUser interface {
	CreateKeystoneUser(ctx context.Context, input *oapi.CreateKeystoneUserInput) (*oapi.KeystoneUser, error)
	GetKeystoneUsers(ctx context.Context, input *oapi.GetKeystoneUsersParams) ([]oapi.KeystoneUser, error)
	GetKeystoneUserByID(ctx context.Context, id string) (*oapi.KeystoneUser, error)
	DeleteKeystoneUser(ctx context.Context, id string) error
}

type IAPIKeystoneRole interface {
	CreateKeystoneRole(ctx context.Context, input *oapi.CreateKeystoneRoleInput) (*oapi.KeystoneRole, error)
	UpdateKeystoneRoleByID(
		ctx context.Context, id string, input *oapi.UpdateKeystoneRoleInput) (*oapi.KeystoneRole, error)
	GetKeystoneRoles(ctx context.Context, input *oapi.GetKeystoneRolesParams) ([]oapi.KeystoneRole, error)
	GetKeystoneRoleByID(ctx context.Context, id string) (*oapi.KeystoneRole, error)
	DeleteKeystoneRole(ctx context.Context, id string) error
	AssignRoleToProject(ctx context.Context, projectID string, userID string, roleID string) error
	UnassignRoleFromProject(ctx context.Context, projectID string, userID string, roleID string) error
	AssignRoleToDomain(ctx context.Context, projectID string, userID string, roleID string) error
	UnassignRoleFromDomain(ctx context.Context, projectID string, userID string, roleID string) error
}
