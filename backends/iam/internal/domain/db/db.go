package db

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

import (
	"context"
	"time"

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
	IDBDomain
	IDBUser
	IDBProject
	IDBOrganization
	IDBTeam
	IDBRole
}

// --------------------------------------------------------------------------------
// Domain Interface
// --------------------------------------------------------------------------------
type GetDomainsInput struct {
	ID   string
	Name string
}

type UpdateDomainByIDInput struct {
	Name        *string
	Description *string
	Extra       map[string]interface{}
}

type CreateDomainInput struct {
	ID          *string
	Name        string
	Description *string
	Extra       map[string]interface{}
}

type IDBDomain interface {
	GetDomains(ctx context.Context, input *GetDomainsInput) ([]model.Domain, error)
	CreateDomain(ctx context.Context, input *CreateDomainInput) (*model.Domain, error)
	UpdateDomainByID(ctx context.Context, id string, input *UpdateDomainByIDInput) error
	DeleteDomainByID(ctx context.Context, id string) error
}

// --------------------------------------------------------------------------------
// Role Interface
// --------------------------------------------------------------------------------
type GetRolesInput struct {
	ID   string
	Name string
}

type UpdateRoleByIDInput struct {
	Name        *string
	Description *string
	Extra       map[string]interface{}
}

type CreateRoleInput struct {
	ID          *string
	Name        string
	Description *string
	Extra       map[string]interface{}
}

type IDBRole interface {
	GetRoles(ctx context.Context, input *GetRolesInput) ([]model.Role, error)
	CreateRole(ctx context.Context, input *CreateRoleInput) (*model.Role, error)
	UpdateRoleByID(ctx context.Context, id string, input *UpdateRoleByIDInput) error
	DeleteRoleByID(ctx context.Context, id string) error

	AssignRoleToProject(ctx context.Context, roleID, userID, projectID string) error
	UnassignRoleFromProject(ctx context.Context, roleID, userID, projectID string) error

	AssignRoleToDomain(ctx context.Context, roleID, userID, domainID string) error
	UnassignRoleFromDomain(ctx context.Context, roleID, userID, domainID string) error

	AssignRoleToTeam(ctx context.Context, roleID, userID, teamID string) error
	UnassignRoleFromTeam(ctx context.Context, roleID, userID, teamID string) error

	AssignRoleToOrganization(ctx context.Context, roleID, userID, teamID string) error
	UnassignRoleFromOrganization(ctx context.Context, roleID, userID, teamID string) error
}

// --------------------------------------------------------------------------------
// User Interface
// --------------------------------------------------------------------------------
type GetUsersInput struct {
	ID   string
	Name string
}

type UpdateUserByIDInput struct {
	Name        *string
	DomainID    string
	LastLoginAt time.Time
}

type CreateUserInput struct {
	ID          *string
	Name        string
	DomainID    string
	LastLoginAt time.Time
}

type IDBUser interface {
	GetUsers(ctx context.Context, input *GetUsersInput) ([]model.User, error)
	CreateUser(ctx context.Context, input *CreateUserInput) (*model.User, error)
	UpdateUserByID(ctx context.Context, id string, input *UpdateUserByIDInput) error
	DeleteUserByID(ctx context.Context, id string) error
}

// --------------------------------------------------------------------------------
// Project Interface
// --------------------------------------------------------------------------------
type GetProjectsInput struct {
	ID   string
	Name string
}

type UpdateProjectByIDInput struct {
	Name        *string
	Description *string
	Extra       map[string]interface{}
	DomainID    string
}

type CreateProjectInput struct {
	ID             *string
	Name           string
	Description    *string
	Extra          map[string]interface{}
	DomainID       string
	OrganizationID string
}

type IDBProject interface {
	GetProjects(ctx context.Context, input *GetProjectsInput) ([]model.Project, error)
	CreateProject(ctx context.Context, input *CreateProjectInput) (*model.Project, error)
	UpdateProjectByID(ctx context.Context, id string, input *UpdateProjectByIDInput) error
	DeleteProjectByID(ctx context.Context, id string) error
}

// --------------------------------------------------------------------------------
// Organization Interface
// --------------------------------------------------------------------------------
type GetOrganizationsInput struct {
	ID   string
	Name string
}

type UpdateOrganizationByIDInput struct {
	Name        *string
	Description *string
	Extra       map[string]interface{}
	DomainID    string
}

type CreateOrganizationInput struct {
	ID          *string
	Name        string
	Description *string
	Extra       map[string]interface{}
	DomainID    string
}

type IDBOrganization interface {
	GetOrganizations(ctx context.Context, input *GetOrganizationsInput) ([]model.Organization, error)
	CreateOrganization(ctx context.Context, input *CreateOrganizationInput) (*model.Organization, error)
	UpdateOrganizationByID(ctx context.Context, id string, input *UpdateOrganizationByIDInput) error
	DeleteOrganizationByID(ctx context.Context, id string) error
}

// --------------------------------------------------------------------------------
// Team Interface
// --------------------------------------------------------------------------------
type GetTeamsInput struct {
	ID   string
	Name string
}

type UpdateTeamByIDInput struct {
	Name        *string
	Description *string
	Extra       map[string]interface{}
	DomainID    string
}

type CreateTeamInput struct {
	ID          *string
	Name        string
	Description *string
	Extra       map[string]interface{}
	DomainID    string
}

type IDBTeam interface {
	GetTeams(ctx context.Context, input *GetTeamsInput) ([]model.Team, error)
	CreateTeam(ctx context.Context, input *CreateTeamInput) (*model.Team, error)
	UpdateTeamByID(ctx context.Context, id string, input *UpdateTeamByIDInput) error
	DeleteTeamByID(ctx context.Context, id string) error
}
