package db

import (
	"context"
	"os"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) MustMigrate(ctx context.Context) {
	if err := self.Migrate(ctx); err != nil {
		print("failed to self.Migrate", err.Error())
		os.Exit(1)
	}
}

func (self *DB) Migrate(ctx context.Context) error {
	if err := self.DB.AutoMigrate(
		&model.Project{},
		&model.Role{},
	); err != nil {
		return tlog.WrapError(ctx, err, "failed to self.DB.AutoMigrate")
	}

	return nil
}
