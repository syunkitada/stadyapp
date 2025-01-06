package main

import (
	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/config"
	"github.com/syunkitada/stadyapp/backends/compute/internal/logic/db"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func main() {
	conf := config.GetDefaultConfig()
	db := db.New(&conf.DB)
	ctx := tlog.NewContext()
	db.MustCreateDatabase(ctx)
	db.MustOpen(ctx)
	db.MustMigrate(ctx)
}
