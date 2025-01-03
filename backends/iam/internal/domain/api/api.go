package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

type IAPI interface {
	CreateKeystoneToken(ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error)
	CreateKeystoneProject(ctx context.Context, input *oapi.CreateKeystoneProjectInput) (*oapi.KeystoneProject, error)
	GetKeystoneProjects(ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error)
	GetKeystoneUserProjects(ctx context.Context, userID string) ([]oapi.KeystoneProject, error)
	GetKeystoneProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error)
	DeleteKeystoneProject(ctx context.Context, id string) error
	CreateKeystoneUser(ctx context.Context, input *oapi.CreateKeystoneUserInput) (*oapi.KeystoneUser, error)
	GetKeystoneUsers(ctx context.Context, input *oapi.GetKeystoneUsersParams) ([]oapi.KeystoneUser, error)
	GetKeystoneUserByID(ctx context.Context, id string) (*oapi.KeystoneUser, error)
	DeleteKeystoneUser(ctx context.Context, id string) error
	CreateKeystoneRole(ctx context.Context, input *oapi.CreateKeystoneRoleInput) (*oapi.KeystoneRole, error)
	GetKeystoneRoles(ctx context.Context, input *oapi.GetKeystoneRolesParams) ([]oapi.KeystoneRole, error)
	GetKeystoneRoleByID(ctx context.Context, id string) (*oapi.KeystoneRole, error)
	DeleteKeystoneRole(ctx context.Context, id string) error
	GetPubkeys(ctx context.Context, input *oapi.GetPubkeysParams) ([]oapi.Pubkey, error)
}
