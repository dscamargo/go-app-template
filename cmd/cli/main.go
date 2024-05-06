package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dscamargo/go_app_template/config"
	"github.com/dscamargo/go_app_template/internal/controllers"
	"github.com/dscamargo/go_app_template/internal/repositories"
	"github.com/dscamargo/go_app_template/internal/services"
	"github.com/dscamargo/go_app_template/migrations"
	"github.com/dscamargo/go_app_template/pkg/web"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"log/slog"
	"os"
	"strings"
)

func main() {
	appConfig := config.New()

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(appConfig.Database.URL)))
	db := bun.NewDB(sqlDb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		// SQL_LOGS=1 logs failed queries
		// SQL_LOGS=2 logs all queries
		bundebug.FromEnv("SQL_LOGS"),
	))

	// Mounting the app services
	exampleRepository := repositories.NewExampleRepository(db)
	exampleService := services.NewExampleService(exampleRepository)

	app := &cli.App{
		Name: "my-app",
		Commands: []*cli.Command{
			appCommands(appConfig, db, exampleService),
			dbCommands(migrate.NewMigrator(db, migrations.Migrations)),
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("app.Run", "error", err.Error())
		os.Exit(1)
	}
}

func appCommands(config *config.Config, db *bun.DB, exampleService *services.ExampleService) *cli.Command {
	return &cli.Command{
		Name:  "app",
		Usage: "app commands",
		Subcommands: []*cli.Command{{
			Name:  "run",
			Usage: "run server",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				app := fiber.New()
				app.Use(cors.New())

				getAllExamplesController := controllers.NewGetAllExampleController(exampleService)

				app.Get("/health", func(c *fiber.Ctx) error {
					_, err := db.NewRaw("SELECT 1").Exec(c.Context())
					if err != nil {
						return web.NewInternalServerError(c)
					}
					return web.NewOKResponse(c, map[string]string{"ping": "pong"})
				})

				//Check jwt
				pubKey, err := web.ReadPublicKey()
				if err != nil {
					slog.Error("ReadPublicKey", "error", err.Error())
					os.Exit(1)
				}
				app.Use(jwtware.New(jwtware.Config{
					SigningKey: jwtware.SigningKey{
						JWTAlg: jwtware.RS256,
						Key:    pubKey,
					},
					ErrorHandler: func(ctx *fiber.Ctx, err error) error {
						slog.InfoContext(ctx.Context(), "JWTMiddleware ErrorHandler", "error", err.Error())
						return web.NewUnauthorizedResponse(ctx, []string{"Unauthorized"})
					},
				}))

				app.Get("/examples", getAllExamplesController.Execute)

				slog.Info("Server Started", "env", config.Server.Env)
				return app.Listen(":" + config.Server.Port)
			},
		}},
	}
}

func dbCommands(migrator *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration file",
				Action: func(c *cli.Context) error {
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer func(migrator *migrate.Migrator, ctx context.Context) {
						err := migrator.Unlock(ctx)
						if err != nil {
							panic(err)
						}
					}(migrator, c.Context)

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						slog.Info("there are no new migrations to run (database is up to date)")
						return nil
					}
					slog.Info("migrated", "group", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer func(migrator *migrate.Migrator, ctx context.Context) {
						err := migrator.Unlock(ctx)
						if err != nil {
							panic(err)
						}
					}(migrator, c.Context)

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						slog.Error("there are no groups to roll back")
						return nil
					}
					slog.Info("rolled back", "group", group)
					return nil
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					slog.Info("created migration", "name", mf.Name, "path", mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						slog.Info("created migration", "name", mf.Name, "path", mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "create_tx_sql",
				Usage: "create up and down transactional SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateTxSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						slog.Info("created transaction migration", "name", mf.Name, "path", mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
		},
	}
}
