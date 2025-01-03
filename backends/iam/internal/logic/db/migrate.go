package db

import (
	"context"
	"log/slog"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func (self *DB) MustMigrate(ctx context.Context) {
	if err := self.Migrate(ctx); err != nil {
		tlog.Fatal(ctx, "failed to MustMigrate", slog.String("err", err.Error()))
	}
}

func (self *DB) Migrate(ctx context.Context) error {
	if err := self.DB.AutoMigrate(
		&model.Project{},
		&model.Role{},
		&model.UserRoleAssignment{},
	); err != nil {
		return tlog.WrapErr(ctx, err, "failed to self.DB.AutoMigrate")
	}

	return nil
}
