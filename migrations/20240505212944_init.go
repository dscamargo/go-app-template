package migrations

import (
	"context"
	"github.com/dscamargo/template_go_w_cli/internal/models"
	"log/slog"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		slog.InfoContext(ctx, "[up migration]", "message", "creating table example")
		_, err := db.NewCreateTable().Model(&models.Example{}).Exec(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		slog.InfoContext(ctx, "[down migration]", "message", "dropping table example")
		_, err := db.NewDropTable().Model(&models.Example{}).Exec(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	})
}
