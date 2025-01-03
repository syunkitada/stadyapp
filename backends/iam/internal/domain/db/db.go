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
	// IDBRole
	// IDBUserRoleAssignment
}

type GetProjectsInput struct {
	ID   string
	Name string
}

type UpdateProjectInput struct {
	ID         string
	Properties map[string]string
}

type CreateProjectInput struct {
	Name       string
	Properties map[string]string
}

type IDBProject interface {
	GetProjects(ctx context.Context, input *GetProjectsInput) ([]model.Project, error)
	CreateProject(ctx context.Context, input *CreateProjectInput) (*model.Project, error)
	UpdateProject(ctx context.Context, input *UpdateProjectInput) error
	DeleteProjectByID(ctx context.Context, id string) error
}

type GetRolesInput struct {
	ID   string
	Name string
}

type UpdateRoleInput struct {
	ID         string
	Properties map[string]string
}

type CreateRoleInput struct {
	Name       string
	Properties map[string]string
}

type IDBRole interface {
	GetRoles(ctx context.Context, input *GetRolesInput) ([]model.Role, error)
	CreateRole(ctx context.Context, input *CreateRoleInput) (*model.Role, error)
	UpdateRole(ctx context.Context, input *UpdateRoleInput) error
	DeleteRoleByID(ctx context.Context, id string) error
}

type GetUserRoleAssignmentsInput struct {
	UserID    string
	ProjectID string
}

type DeleteUserRoleAssignmentsInput struct {
	UserID    string
	RoleID    string
	ProjectID string
}

type CreateUserRoleAssignmentInput struct {
	UserID    string
	RoleID    string
	ProjectID string
}

type IDBUserRoleAssignment interface {
	GetUserRoleAssignments(ctx context.Context, input *GetUserRoleAssignmentsInput) ([]model.UserRoleAssignment, error)
	CreateUserRoleAssignment(ctx context.Context, input *CreateUserRoleAssignmentInput) (*model.UserRoleAssignment, error)
	DeleteUserRoleAssignment(ctx context.Context, input *DeleteUserRoleAssignmentsInput) error
}
