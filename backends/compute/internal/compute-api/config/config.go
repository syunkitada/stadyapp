package config

import (
	"github.com/syunkitada/stadyapp/backends/compute/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/compute/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	DB      db.Config
	Logger  tlog.Config
	IAMAuth iam_auth.Config
}

func GetDefaultConfig() Config {
	return Config{
		DB:      db.GetDefaultConfig(),
		Logger:  tlog.GetDefaultConfig(),
		IAMAuth: iam_auth.GetDefaultConfig(),
	}
}
