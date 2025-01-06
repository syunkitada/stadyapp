package api

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) GetKeystoneGroups(
	ctx context.Context, input *oapi.GetKeystoneGroupsParams) ([]oapi.KeystoneGroup, error) {

	getTeamsInput := db.GetTeamsInput{}
	if input.Name != nil {
		getTeamsInput.Name = *input.Name
	}

	dbTeams, err := self.db.GetTeams(ctx, &getTeamsInput)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	groups := []oapi.KeystoneGroup{}
	for i := range dbTeams {
		group, err := ConvertDBTeamToAPIGroup(ctx, &dbTeams[i])
		if err != nil {
			return nil, tlog.Err(ctx, err)
		}
		groups = append(groups, *group)
	}

	return groups, nil
}

func (self *API) GetKeystoneGroupByID(ctx context.Context, id string) (*oapi.KeystoneGroup, error) {
	teamID := strings.Replace(id, ProjectTagTeam+"@", "", 1)

	dbTeam, err := self.db.GetTeam(ctx, &db.GetTeamsInput{ID: teamID})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	group, err := ConvertDBTeamToAPIGroup(ctx, dbTeam)
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return group, nil
}

func ConvertDBTeamToAPIGroup(ctx context.Context, dbTeam *model.Team) (*oapi.KeystoneGroup, error) {
	var additionalProperties map[string]interface{}
	if err := json.Unmarshal([]byte(dbTeam.Extra), &additionalProperties); err != nil {
		return nil, tlog.Err(ctx, err)
	}

	return &oapi.KeystoneGroup{
		Id:       "Team@" + dbTeam.ID,
		Name:     dbTeam.Name,
		DomainId: dbTeam.DomainID,
	}, nil
}
