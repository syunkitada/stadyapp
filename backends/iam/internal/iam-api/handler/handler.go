package handler

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
)

type Handler struct {
	conf *config.Config
	api  api.IAPI
}

func NewHandler(conf *config.Config, api api.IAPI) *Handler {
	return &Handler{
		conf: conf,
		api:  api,
	}
}
