package handler

import (
	domain_api "github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/api"
)

type Handler struct {
	conf *config.Config
	api  domain_api.IAPI
}

func NewHandler(conf *config.Config, db db.IDB) *Handler {
	api := api.New(conf, db)

	return &Handler{
		conf: conf,
		api:  api,
	}
}
