package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

const (
	ProjectTagTeam         = "team"
	ProjectTagOrganization = "organization"
	ProjectTagSeparator    = "@"
)

func (self *API) CreateKeystoneProject(
	ctx context.Context, input *oapi.CreateKeystoneProjectInput) (*oapi.KeystoneProject, error) {
	if input.Project.Tags != nil {
		tags := *input.Project.Tags
		if slices.Contains(tags, ProjectTagTeam) {
			dbTeam, err := self.db.CreateTeam(ctx, &db.CreateTeamInput{
				Name:        input.Project.Name,
				Description: input.Project.Description,
				Extra:       input.Project.AdditionalProperties,
				DomainID:    input.Project.DomainId,
			})
			if err != nil {
				return nil, tlog.Err(ctx, err)
			}

			project, err := ConvertDBTeamToAPIProject(ctx, dbTeam)
			if err != nil {
				return nil, tlog.Err(ctx, err)
			}
			return project, nil

		} else if slices.Contains(tags, ProjectTagOrganization) {
			dbOrganization, err := self.db.CreateOrganization(ctx, &db.CreateOrganizationInput{
				Name:        input.Project.Name,
				Description: input.Project.Description,
				Extra:       input.Project.AdditionalProperties,
				DomainID:    input.Project.DomainId,
			})
			if err != nil {
				return nil, tlog.Err(ctx, err)
			}

			project, err := ConvertDBOrganizationToAPIProject(ctx, dbOrganization)
			if err != nil {
				return nil, tlog.Err(ctx, err)
			}
			return project, nil
		}
	}

	if input.Project.OrganizationId == nil {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusBadRequest, "organization_id is required"))
	}

	organizationID := strings.Replace(*input.Project.OrganizationId, ProjectTagOrganization+ProjectTagSeparator, "", 1)

	dbProject, err := self.db.CreateProject(ctx, &db.CreateProjectInput{
		Name:           input.Project.Name,
		Description:    input.Project.Description,
		Extra:          input.Project.AdditionalProperties,
		DomainID:       input.Project.DomainId,
		OrganizationID: organizationID,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	project, err := ConvertDBProjectToAPIProject(ctx, dbProject)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return project, nil
}

func (self *API) UpdateKeystoneProjectByID(
	ctx context.Context, id string, input *oapi.UpdateKeystoneProjectInput) (*oapi.KeystoneProject, error) {
	if input.Project.Tags != nil {
		tags := *input.Project.Tags
		if slices.Contains(tags, ProjectTagTeam) {

		} else if slices.Contains(tags, ProjectTagOrganization) {

		}
	}

	err := self.db.UpdateProjectByID(ctx, id, &db.UpdateProjectByIDInput{
		Name:        input.Project.Name,
		Description: input.Project.Description,
		Extra:       input.Project.AdditionalProperties,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	project, err := self.GetKeystoneProjectByID(ctx, id)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return project, nil
}

func (self *API) GetKeystoneOrganizationProjects(
	ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error) {

	getOrganizationsInput := db.GetOrganizationsInput{}
	if input.Name != nil {
		getOrganizationsInput.Name = *input.Name
	}

	dbOrganizations, err := self.db.GetOrganizations(ctx, &getOrganizationsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}

	for i := range dbOrganizations {
		project, err := ConvertDBOrganizationToAPIProject(ctx, &dbOrganizations[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		projects = append(projects, *project)
	}

	return projects, nil
}

func (self *API) GetKeystoneTeamProjects(
	ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error) {
	getTeamsInput := db.GetTeamsInput{}
	if input.Name != nil {
		getTeamsInput.Name = *input.Name
	}

	dbTeams, err := self.db.GetTeams(ctx, &getTeamsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}

	for i := range dbTeams {
		project, err := ConvertDBTeamToAPIProject(ctx, &dbTeams[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		projects = append(projects, *project)
	}

	return projects, nil
}

func (self *API) GetKeystoneProjects(
	ctx context.Context, input *oapi.GetKeystoneProjectsParams) ([]oapi.KeystoneProject, error) {
	if input.Tags != nil {
		tags := *input.Tags
		fmt.Println("tags", tags) //nolint
		if slices.Contains(tags, ProjectTagTeam) {
			if projects, err := self.GetKeystoneTeamProjects(ctx, input); err != nil {
				return nil, tlog.Err(ctx, err)
			} else {
				return projects, nil
			}
		} else if slices.Contains(tags, ProjectTagOrganization) {
			if projects, err := self.GetKeystoneOrganizationProjects(ctx, input); err != nil {
				return nil, tlog.Err(ctx, err)
			} else {
				return projects, nil
			}
		}
	}

	getProjectsInput := db.GetProjectsInput{Name: input.Name}

	dbProjects, err := self.db.GetProjects(ctx, &getProjectsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}

	for i := range dbProjects {
		project, err := ConvertDBProjectToAPIProject(ctx, &dbProjects[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		projects = append(projects, *project)
	}

	return projects, nil
}

func (self *API) GetKeystoneUserProjects(
	ctx context.Context, userID string) ([]oapi.KeystoneProject, error) {
	authContext, err := iam_auth.GetAuthContext(ctx)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if authContext.UserID != userID {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusForbidden, "forbidden"))
	}

	fmt.Println("context", authContext.ProjectID) //nolint

	getProjectsInput := db.GetProjectsInput{}
	dbProjects, err := self.db.GetProjects(ctx, &getProjectsInput)

	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	projects := []oapi.KeystoneProject{}

	for i := range dbProjects {
		project, err := ConvertDBProjectToAPIProject(ctx, &dbProjects[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}

		projects = append(projects, *project)
	}

	return projects, nil
}

func (self *API) GetKeystoneOrganizationProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error) {
	dbOrganizations, err := self.db.GetOrganizations(ctx, &db.GetOrganizationsInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbOrganizations) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "team not found"))
	}

	if len(dbOrganizations) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "team is duplicated"))
	}

	project, err := ConvertDBOrganizationToAPIProject(ctx, &dbOrganizations[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}
	return project, nil
}

func (self *API) GetKeystoneTeamProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error) {
	dbTeams, err := self.db.GetTeams(ctx, &db.GetTeamsInput{
		ID: id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbTeams) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "team not found"))
	}

	if len(dbTeams) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "team is duplicated"))
	}

	project, err := ConvertDBTeamToAPIProject(ctx, &dbTeams[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}
	return project, nil
}

func (self *API) GetKeystoneProjectByID(ctx context.Context, id string) (*oapi.KeystoneProject, error) {
	splitedID := strings.Split(id, ProjectTagSeparator)
	if len(splitedID) == 2 {
		if splitedID[0] == ProjectTagTeam {
			if project, err := self.GetKeystoneTeamProjectByID(ctx, splitedID[1]); err != nil {
				return nil, tlog.Err(ctx, err)
			} else {
				return project, nil
			}
		} else if splitedID[0] == ProjectTagOrganization {
			if project, err := self.GetKeystoneOrganizationProjectByID(ctx, splitedID[1]); err != nil {
				return nil, tlog.Err(ctx, err)
			} else {
				return project, nil
			}
		}
	}

	dbProjects, err := self.db.GetProjects(ctx, &db.GetProjectsInput{
		ID: &id,
	})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	if len(dbProjects) == 0 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusNotFound, "project not found"))
	}

	if len(dbProjects) > 1 {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusConflict, "project is duplicated"))
	}

	project, err := ConvertDBProjectToAPIProject(ctx, &dbProjects[0])
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return project, nil
}

func (self *API) DeleteKeystoneProject(ctx context.Context, id string) error {
	splitedID := strings.Split(id, ProjectTagSeparator)
	if len(splitedID) == 2 {
		if splitedID[0] == ProjectTagTeam {
			if err := self.db.DeleteTeamByID(ctx, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		} else if splitedID[0] == ProjectTagOrganization {
			if err := self.db.DeleteOrganizationByID(ctx, splitedID[1]); err != nil {
				return tlog.Err(ctx, err)
			}

			return nil
		}
	}

	err := self.db.DeleteProjectByID(ctx, id)
	if err != nil {
		return tlog.Err(ctx, err)
	}

	return nil
}

func ConvertDBProjectToAPIProject(ctx context.Context, dbProject *model.Project) (*oapi.KeystoneProject, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbProject.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneProject{
		Id:                   dbProject.ID,
		Name:                 dbProject.Name,
		Description:          dbProject.Description,
		AdditionalProperties: additionalProperties,
		DomainId:             dbProject.DomainID,
	}, nil
}

func ConvertDBOrganizationToAPIProject(ctx context.Context, dbOrganization *model.Organization) (*oapi.KeystoneProject, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbOrganization.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneProject{
		Id:                   ProjectTagOrganization + ProjectTagSeparator + dbOrganization.ID,
		Name:                 dbOrganization.Name,
		Description:          dbOrganization.Description,
		AdditionalProperties: additionalProperties,
		DomainId:             dbOrganization.DomainID,
	}, nil
}

func ConvertDBTeamToAPIProject(ctx context.Context, dbTeam *model.Team) (*oapi.KeystoneProject, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbTeam.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneProject{
		Id:                   ProjectTagTeam + ProjectTagSeparator + dbTeam.ID,
		Name:                 dbTeam.Name,
		Description:          dbTeam.Description,
		AdditionalProperties: additionalProperties,
		DomainId:             dbTeam.DomainID,
	}, nil
}
