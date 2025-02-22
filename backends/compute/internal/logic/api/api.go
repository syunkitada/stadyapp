package api

import (
	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/config"
	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/pkg/libs/iam_auth"
)

type API struct {
	conf    *config.Config
	db      db.IDB
	iamAuth *iam_auth.IAMAuth
}

func New(conf *config.Config, db db.IDB, iamAuth *iam_auth.IAMAuth) api.IAPI { //nolint:ireturn
	return &API{
		conf:    conf,
		db:      db,
		iamAuth: iamAuth,
	}
}
