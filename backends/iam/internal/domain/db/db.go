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
	OwnerUserID *string
}

type IDBDomain interface {
	GetDomain(ctx context.Context, input *GetDomainsInput) (*model.Domain, error)
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

type GetDomainRoleAssignmentsInput struct{}

type GetProjectRoleAssignmentsInput struct{}

type GetOrganizationRoleAssignmentsInput struct{}

type GetTeamRoleAssignmentsInput struct{}

type GetUserProjectRolesInput struct {
	UserID    string
	ProjectID string
}

type IDBRole interface {
	GetRoleByID(ctx context.Context, id string) (*model.Role, error)
	GetRoles(ctx context.Context, input *GetRolesInput) ([]model.Role, error)
	CreateRole(ctx context.Context, input *CreateRoleInput) (*model.Role, error)
	UpdateRoleByID(ctx context.Context, id string, input *UpdateRoleByIDInput) error
	DeleteRoleByID(ctx context.Context, id string) error

	GetDomainRoleAssignments(ctx context.Context, input *GetDomainRoleAssignmentsInput) ([]model.DomainRoleAssignmentDetail, error)
	AssignRoleToUserDomain(ctx context.Context, roleID, userID, domainID string) error
	UnassignRoleFromUserDomain(ctx context.Context, roleID, userID, domainID string) error
	AssignRoleToTeamDomain(ctx context.Context, roleID, teamID, domainID string) error
	UnassignRoleFromTeamDomain(ctx context.Context, roleID, teamID, domainID string) error

	GetProjectRoleAssignments(ctx context.Context, input *GetProjectRoleAssignmentsInput) ([]model.ProjectRoleAssignmentDetail, error)
	AssignRoleToUserProject(ctx context.Context, roleID, userID, projectID string) error
	UnassignRoleFromUserProject(ctx context.Context, roleID, userID, projectID string) error
	AssignRoleToTeamProject(ctx context.Context, roleID, teamID, projectID string) error
	UnassignRoleFromTeamProject(ctx context.Context, roleID, teamID, projectID string) error

	GetOrganizationRoleAssignments(ctx context.Context, input *GetOrganizationRoleAssignmentsInput) ([]model.OrganizationRoleAssignmentDetail, error)
	AssignRoleToUserOrganization(ctx context.Context, roleID, userID, teamID string) error
	UnassignRoleFromUserOrganization(ctx context.Context, roleID, userID, teamID string) error
	AssignRoleToTeamOrganization(ctx context.Context, roleID, teamID, organizationID string) error
	UnassignRoleFromTeamOrganization(ctx context.Context, roleID, teamID, organizationID string) error

	GetTeamRoleAssignments(ctx context.Context, input *GetTeamRoleAssignmentsInput) ([]model.TeamRoleAssignmentDetail, error)
	AssignRoleToUserTeam(ctx context.Context, roleID, userID, teamID string) error
	UnassignRoleFromUserTeam(ctx context.Context, roleID, userID, teamID string) error

	GetUserProjectRoles(ctx context.Context, input *GetUserProjectRolesInput) ([]model.UserProjectRole, error)
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
	GetUser(ctx context.Context, input *GetUsersInput) (*model.User, error)
	GetUsers(ctx context.Context, input *GetUsersInput) ([]model.User, error)
	CreateUser(ctx context.Context, input *CreateUserInput) (*model.User, error)
	UpdateUserByID(ctx context.Context, id string, input *UpdateUserByIDInput) error
	DeleteUserByID(ctx context.Context, id string) error
}

// --------------------------------------------------------------------------------
// Project Interface
// --------------------------------------------------------------------------------
type GetProjectsInput struct {
	ID   *string
	Name *string
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
	OwnerUserID    string
}

type IDBProject interface {
	GetProject(ctx context.Context, input *GetProjectsInput) (*model.Project, error)
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
	OwnerUserID string
}

type IDBOrganization interface {
	GetOrganization(ctx context.Context, input *GetOrganizationsInput) (*model.Organization, error)
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
	OwnerUserID string
}

type IDBTeam interface {
	GetTeam(ctx context.Context, input *GetTeamsInput) (*model.Team, error)
	GetTeams(ctx context.Context, input *GetTeamsInput) ([]model.Team, error)
	CreateTeam(ctx context.Context, input *CreateTeamInput) (*model.Team, error)
	UpdateTeamByID(ctx context.Context, id string, input *UpdateTeamByIDInput) error
	DeleteTeamByID(ctx context.Context, id string) error
}
