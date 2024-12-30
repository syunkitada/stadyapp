package api

import (
	"context"
	"time"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

func (self *API) CreateKeystoneToken(ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error) {
	// dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{})

	// if err != nil {
	// 	return nil, tlog.WrapError(ctx, err, "failed to self.db.FindProjects")
	// }

	// projects := []oapi.Project{}
	// for _, dbProject := range dbProjects {
	// 	projects = append(projects, oapi.Project{
	// 		Id:   dbProject.ID,
	// 		Name: dbProject.Name,
	// 	})
	// }

	token := oapi.KeystoneToken{
		Token: oapi.KeystoneTokenData{
			AuditIds:  []string{"audit_id1", "audit_id2"},
			Methods:   []string{"password"},
			ExpiresAt: time.Now(),
			IssuedAt:  time.Now(),
			User: oapi.KeystoneTokenUser{
				Domain: oapi.KeystoneTokenDomain{
					Id:   "domain_id",
					Name: "domain_name",
				},
				Id:                "user_id",
				Name:              "user_name",
				PasswordExpiresAt: time.Now(),
			},
			Project: oapi.KeystoneTokenProject{
				Domain: oapi.KeystoneTokenDomain{
					Id:   "domain_id",
					Name: "domain_name",
				},
				Id:   "project_id",
				Name: "project_name",
			},
			Roles: []oapi.KeystoneTokenRole{
				{
					Id:   "role_id",
					Name: "role_name",
				},
			},
			Catalog: []oapi.KeystoneCatalog{
				{
					Id:   "catalog_id",
					Name: "keystone",
					Type: "identity",
					Endpoints: []oapi.KeystoneEndpoint{
						{
							Id:        "endpoint_id",
							Interface: "public",
							Region:    "region1",
							Url:       "http://localhost:10080/api/iam/v3",
						},
					},
				},
			},
		},
	}
	tokenStr := "token_id"
	return &token, tokenStr, nil
}
