package db

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
)

type IDBCommon interface {
	MustOpenMock(ctx context.Context) sqlmock.Sqlmock
	MustOpen(ctx context.Context)
	MustClose(ctx context.Context)
	MustCreateDatabase(ctx context.Context)
	MustRecreateDatabase(ctx context.Context)
	MustMigrate(ctx context.Context)
}

type IDB interface {
	IDBCommon
	IDBProject
	IDBRole
}

type FindProjectsInput struct {
	ID   uint64
	Name string
}

type IDBProject interface {
	FindProjects(ctx context.Context, input *FindProjectsInput) ([]model.Project, error)
	AddProject(ctx context.Context, item *model.Project) (*model.Project, error)
	DeleteProject(ctx context.Context, id uint64) error
}

type FindRolesInput struct {
	ID   uint64
	Name string
}

type IDBRole interface {
	FindRoles(ctx context.Context, input *FindRolesInput) ([]model.Role, error)
	AddRole(ctx context.Context, item *model.Role) (*model.Role, error)
	DeleteRole(ctx context.Context, id uint64) error
}
