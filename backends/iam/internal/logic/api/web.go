package api

import (
	"context"
	"fmt"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *API) GetWebUser(ctx context.Context) (*oapi.WebUser, error) {
	fmt.Println("GetWebUser")
	projects, err := self.GetKeystoneProjects(ctx, &oapi.GetKeystoneProjectsParams{})
	if err != nil {
		return nil, tlog.Err(ctx, err)
	}

	webUser := &oapi.WebUser{
		User: oapi.User{
			Name: "web_user",
		},
		Projects: projects,
	}

	return webUser, nil
}
