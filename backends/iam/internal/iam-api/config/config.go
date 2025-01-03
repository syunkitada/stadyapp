package config

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type Config struct {
	DB           db.Config
	Logger       tlog.Config
	IAMTokenAuth iam_token_auth.Config
}

func GetDefaultConfig() Config {
	return Config{
		DB:           db.GetDefaultConfig(),
		Logger:       tlog.GetDefaultConfig(),
		IAMTokenAuth: iam_token_auth.GetDefaultConfig(),
	}
}
