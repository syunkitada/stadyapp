package config

import (
	"github.com/syunkitada/stadyapp/backends/compute/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/echo_server"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	DB      db.Config
	Logger  tlog.Config
	IAMAuth iam_auth.Config
	Compute Compute
	Server  echo_server.Config
}

func GetDefaultConfig() Config {
	return Config{
		DB:      db.GetDefaultConfig(),
		Logger:  tlog.GetDefaultConfig(),
		IAMAuth: iam_auth.GetDefaultConfig(),
		Server: echo_server.Config{
			Port:         10082,
			AllowOrigins: []string{"http://myapp.localhost.test:5173"},
		},
		Compute: Compute{
			ProxyCatalog: ProxyCatalog{
				Glance: Proxy{
					URL:         "http://localhost:9292",
					OldBasePath: "/glance",
					NewBasePath: "/",
				},
				Nova: Proxy{
					URL:         "http://localhost:8774",
					OldBasePath: "/nova",
					NewBasePath: "/",
				},
				Neutron: Proxy{
					URL:         "http://localhost:9696",
					OldBasePath: "/neutron",
					NewBasePath: "/",
				},
				Placement: Proxy{
					URL:         "http://localhost:8778",
					OldBasePath: "/placement",
					NewBasePath: "/",
				},
			},
		},
	}
}

type Compute struct {
	ProxyCatalog ProxyCatalog
}

type ProxyCatalog struct {
	Glance    Proxy
	Nova      Proxy
	Neutron   Proxy
	Placement Proxy
}

type Proxy struct {
	URL         string
	OldBasePath string
	NewBasePath string
}
