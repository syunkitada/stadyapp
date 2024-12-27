package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

func main() {
	conf := tlog.GetDefaultConfig()
	tlog.Init(&conf)

	ctx := tlog.NewContext()
	err := HelloWorld(ctx)
	if err != nil {
		tlog.Error(ctx, err.Error())
	}
}

func HelloWorld(ctx context.Context) error {
	tlog.Info(ctx, "Hello, Info!", slog.String("tag1", "value"))
	tlog.Warn(ctx, "Hello, Warn!", slog.String("tag1", "value"))
	tlog.Error(ctx, "Hello, Error!", slog.String("tag1", "value"))

	return tlog.WrapError(ctx, fmt.Errorf("DEBUG"), "Wrap", slog.String("tag1", "value"))
}
