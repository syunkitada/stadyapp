package db

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) MustMigrate(ctx context.Context) {
	if err := self.Migrate(ctx); err != nil {
		tlog.Fatal(ctx, "failed to MustMigrate", slog.String("err", err.Error()))
	}
}

func (self *DB) Migrate(ctx context.Context) error {
	if err := self.DB.AutoMigrate(
		&model.Project{},
		&model.Role{},
		&model.Organization{},
		&model.Team{},
		&model.DomainRoleAssignment{},
		&model.OrganizationRoleAssignment{},
		&model.ProjectRoleAssignment{},
		&model.TeamRoleAssignment{},
	); err != nil {
		return tlog.Err(ctx, err)
	}

	defaultID := "default"
	adminID := "admin"

	{
		domains, err := self.GetDomains(ctx, &db.GetDomainsInput{
			ID: defaultID,
		})
		if err != nil {
			return tlog.Err(ctx, err)
		}

		if len(domains) == 0 {
			_, err = self.CreateDomain(ctx, &db.CreateDomainInput{
				ID:   &defaultID,
				Name: defaultID,
			})
			if err != nil {
				return tlog.Err(ctx, err)
			}
		}
	}

	{
		organizations, err := self.GetOrganizations(ctx, &db.GetOrganizationsInput{
			ID: adminID,
		})
		if err != nil {
			return tlog.Err(ctx, err)
		}

		if len(organizations) == 0 {
			_, err = self.CreateOrganization(ctx, &db.CreateOrganizationInput{
				DomainID: defaultID,
				ID:       &adminID,
				Name:     adminID,
			})
			if err != nil {
				return tlog.Err(ctx, err)
			}
		}
	}

	{
		projects, err := self.GetProjects(ctx, &db.GetProjectsInput{
			ID: adminID,
		})
		if err != nil {
			return tlog.Err(ctx, err)
		}

		if len(projects) == 0 {
			_, err = self.CreateProject(ctx, &db.CreateProjectInput{
				DomainID:       defaultID,
				ID:             &adminID,
				Name:           adminID,
				OrganizationID: adminID,
			})
			if err != nil {
				return tlog.Err(ctx, err)
			}
		}
	}

	{
		teams, err := self.GetTeams(ctx, &db.GetTeamsInput{
			ID: adminID,
		})
		if err != nil {
			return tlog.Err(ctx, err)
		}

		if len(teams) == 0 {
			_, err = self.CreateTeam(ctx, &db.CreateTeamInput{
				DomainID: defaultID,
				ID:       &adminID,
				Name:     adminID,
			})
			if err != nil {
				return tlog.Err(ctx, err)
			}
		}
	}

	defaultRoles := []string{
		"admin",
		"service",
		"manager",
		"member",
	}

	for _, roleName := range defaultRoles {
		_, err := self.CreateRoleIfNotExists(ctx, &db.CreateRoleInput{
			ID:   &roleName,
			Name: roleName,
		})
		if err != nil {
			return tlog.Err(ctx, err)
		}
	}

	return nil
}

func (self *DB) CreateRoleIfNotExists(ctx context.Context, input *db.CreateRoleInput) (*model.Role, error) {
	if input.ID == nil {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusBadRequest, "id is required"))
	}

	roles, err := self.GetRoles(ctx, &db.GetRolesInput{
		ID: *input.ID,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(roles) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "role is duplicated"))
	}

	if len(roles) == 1 {
		return &roles[0], nil
	}

	role, err := self.CreateRole(ctx, input)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return role, nil
}
