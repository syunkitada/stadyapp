package config

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/echo_server"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	DB      db.Config
	Logger  tlog.Config
	IAMAuth iam_auth.Config
	IAM     IAMConfig
	Server  echo_server.Config
}

func GetDefaultConfig() Config {
	return Config{
		DB:      db.GetDefaultConfig(),
		Logger:  tlog.GetDefaultConfig(),
		IAMAuth: iam_auth.GetDefaultConfig(),
		Server: echo_server.Config{
			Port:         10081,
			AllowOrigins: []string{"https://myapp.localhost.test:11443", "http://myapp.localhost.test:5173"},
		},
		IAM: IAMConfig{
			Catalogs: []Catalog{
				{
					Type: "identity",
					Name: "keystone",
					Endpoints: []Endpoint{
						{
							Interface: "public",
							Region:    "region1",
							URL:       "http://localhost:11080/api/iam/keystone/v3",
						},
					},
				},
				{
					Type: "image",
					Name: "glance",
					Endpoints: []Endpoint{
						{
							Interface: "public",
							Region:    "region1",
							URL:       "http://localhost:11080/api/compute/glance/v2",
						},
					},
				},
				{
					Type: "compute",
					Name: "nova",
					Endpoints: []Endpoint{
						{
							Interface: "public",
							Region:    "region1",
							URL:       "http://localhost:11080/api/compute/nova/v2.1",
						},
					},
				},
				{
					Type: "network",
					Name: "neutron",
					Endpoints: []Endpoint{
						{
							Interface: "public",
							Region:    "region1",
							URL:       "http://localhost:11080/api/compute/neutron",
						},
					},
				},
				{
					Type: "placement",
					Name: "placement",
					Endpoints: []Endpoint{
						{
							Interface: "public",
							Region:    "region1",
							URL:       "http://localhost:11080/api/compute/placement",
						},
					},
				},
			},
		},
	}
}

type Endpoint struct {
	Interface string
	Region    string
	URL       string
}

type Catalog struct {
	Type      string
	Name      string
	Endpoints []Endpoint
}

type IAMConfig struct {
	Catalogs []Catalog
}
