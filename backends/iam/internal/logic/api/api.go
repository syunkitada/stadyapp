package api

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_token_auth"
)

type API struct {
	conf         *config.Config
	db           db.IDB
	iamTokenAuth *iam_token_auth.IAMTokenAuth
}

func New(conf *config.Config, db db.IDB, iamTokenAuth *iam_token_auth.IAMTokenAuth) api.IAPI { //nolint:ireturn
	return &API{
		conf:         conf,
		db:           db,
		iamTokenAuth: iamTokenAuth,
	}
}
