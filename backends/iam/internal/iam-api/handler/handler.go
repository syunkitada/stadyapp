package handler

import (
	"github.com/labstack/echo/v4"
	domain_api "github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
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

// sendHandlerError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendHandlerError(ctx echo.Context, code int, message string) error {
	itemErr := oapi.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, itemErr)
	return err
}
