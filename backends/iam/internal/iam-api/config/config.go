package config

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/tlog"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
)

type Config struct {
	DB     db.Config
	Logger tlog.Config
}

func GetDefaultConfig() Config {
	return Config{
		DB:     db.GetDefaultConfig(),
		Logger: tlog.GetDefaultConfig(),
	}
}
