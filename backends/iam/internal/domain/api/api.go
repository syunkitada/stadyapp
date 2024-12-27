package api

import (
	"context"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

type IAPI interface {
	FindProjects(ctx context.Context, params oapi.FindProjectsParams) (items []oapi.Project, err error)
	FindProjectByID(ctx context.Context, id uint64) (item *oapi.Project, err error)
	AddProject(ctx context.Context, item *oapi.NewProject) error
	DeleteProject(ctx context.Context, id uint64) error
	FindRoles(ctx context.Context, params oapi.FindRolesParams) (items []oapi.Role, err error)
	FindRoleByID(ctx context.Context, id uint64) (item *oapi.Role, err error)
	AddRole(ctx context.Context, item *oapi.NewRole) error
	DeleteRole(ctx context.Context, id uint64) error
}
