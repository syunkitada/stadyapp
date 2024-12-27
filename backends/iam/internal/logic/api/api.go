package api

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
)

type API struct {
	conf *config.Config
	db   db.IDB
}

func New(conf *config.Config, db db.IDB) api.IAPI { //nolint:ireturn
	return &API{
		conf: conf,
		db:   db,
	}
}
