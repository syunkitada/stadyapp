package db

import (
	"os"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
)

func (self *DB) MustMigrate() {
	if err := self.Migrate(); err != nil {
		print("failed to self.Migrate", err.Error())
		os.Exit(1)
	}
}

func (self *DB) Migrate() (err error) {
	if err = self.DB.AutoMigrate(
		&model.Project{},
		&model.Role{},
	); err != nil {
		return err
	}
	return nil
}
