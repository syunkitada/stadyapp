package main

import (
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
)

func main() {
	conf := config.GetDefaultConfig()
	db := db.New(&conf.DB)
	db.MustCreateDatabase()
	db.MustOpen()
	db.MustMigrate()
}
