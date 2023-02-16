package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"renthome/boiler"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/ninja-software/terror/v2"
	"github.com/oklog/run"
	"github.com/urfave/cli/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

// Build Version
const Version = "v1.0.0"

func main() {

	app := &cli.App{
		Compiled: time.Now(),
		Usage:    "Seed renthome database",
		Authors: []*cli.Author{
			{
				Name:  "Sanam Limbu",
				Email: "sudosanam@gmail.com",
			},
		},
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name: "db",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "database_user", Value: "renthome", EnvVars: []string{"RENTHOME_DATABASE_USER"}, Usage: "The database user"},
					&cli.StringFlag{Name: "database_pass", Value: "devdev", EnvVars: []string{"RENTHOME_DATABASE_PASS"}, Usage: "The database pass"},
					&cli.StringFlag{Name: "database_host", Value: "localhost", EnvVars: []string{"RENTHOME_DATABASE_HOST"}, Usage: "The database host"},
					&cli.StringFlag{Name: "database_port", Value: "5438", EnvVars: []string{"RENTHOME_DATABASE_PORT"}, Usage: "The database port"},
					&cli.StringFlag{Name: "database_name", Value: "renthome", EnvVars: []string{"RENTHOME_DATABASE_NAME"}, Usage: "The database name"},
					&cli.StringFlag{Name: "database_application_name", Value: "API Server", EnvVars: []string{"RENTHOME_DATABASE_APPLICATION_NAME"}, Usage: "Postgres database name"},
					&cli.IntFlag{Name: "database_min_conns", Value: 10, EnvVars: []string{"RENTHOME_DATABASE_MIN_CONNS"}, Usage: "The database minimum connections"},
					&cli.IntFlag{Name: "database_max_conns", Value: 80, EnvVars: []string{"RENTHOME_DATABASE_MAX_CONNS"}, Usage: "The database maximum connections"},

					&cli.StringFlag{Name: "log_level", Value: "InfoLevel", EnvVars: []string{"RENTHOME_LOG_LEVEL"}, Usage: "Set the log level for zerolog (Options: DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel, NoLevel, Disabled, TraceLevel"},
					&cli.StringFlag{Name: "environment", Value: "development", DefaultText: "development", EnvVars: []string{"RENTHOME_ENVIRONMENT"}, Usage: "This program environment (development, testing, training, staging, production)"},

					&cli.StringFlag{Name: "api_addr", Value: ":8084", EnvVars: []string{"RENTHOME_API_ADDR"}, Usage: "host:port to run the API"},
					&cli.StringFlag{Name: "rootpath", Value: "../web/dist", EnvVars: []string{"RENTHOME_ROOTPATH"}, Usage: "folder path of index.html"},
					&cli.StringFlag{Name: "jwtsecret", Value: "a35eab71-f691-4dc3-98e5-980bda774fa0", EnvVars: []string{"RENTHOME_USERAUTH_JWTSECRET"}, Usage: "JWT secret"},
					&cli.StringFlag{Name: "google_client_id", Value: "", EnvVars: []string{"RENTHOME_GOOGLE_CLIENT_ID"}, Usage: "Google Client ID for OAuth functionaility."},
					&cli.StringFlag{Name: "facebook_client_id", Value: "", EnvVars: []string{"RENTHOME_FACEBOOK_CLIENT_ID"}, Usage: "Facebook Client ID for OAuth functionaility."},
					&cli.StringFlag{Name: "apple_client_id", Value: "", EnvVars: []string{"RENTHOME_APPLE_CLIENT_ID"}, Usage: "Apple Client ID for OAuth functionaility."},
					&cli.BoolFlag{Name: "cookie_secure", Value: false, EnvVars: []string{"RENTHOME_COOKIE_SECURE"}, Usage: "Cookie Secure setting option for secure cookies."},

					&cli.StringFlag{Name: "admin_host_url", Value: "http://localhost:5001", EnvVars: []string{"RENTHOME_ADMIN_FRONTEND_HOST_URL"}, Usage: "The Admin Site URL used for links in the mailer and allowed cors urls"},
					&cli.StringFlag{Name: "public_host_url", Value: "http://localhost:5002", EnvVars: []string{"RENTHOME_PUBLIC_FRONTEND_HOST_URL"}, Usage: "The Public Site URL used for links in the mailer and allowed cors urls"},

					&cli.IntFlag{Name: "tokenexpirydays", Value: 30, EnvVars: []string{"RENTHOME_USERAUTH_TOKENEXPIRYDAYS"}, Usage: "How many days before the token expires"},
					&cli.IntFlag{Name: "blacklistrefreshhours", Value: 1, EnvVars: []string{"RENTHOME_USERAUTH_BLACKLISTREFRESHHOURS"}, Usage: "How often should the issued_tokens list be cleared of expired tokens in hours"},
				},

				Usage: "seed database",
				Action: func(c *cli.Context) error {
					ctx, cancel := context.WithCancel(c.Context)

					g := &run.Group{}
					// Listen for os.interrupt
					g.Add(run.SignalHandler(ctx, os.Interrupt))
					// start the seed
					g.Add(func() error { return SeedFunc(c, ctx) }, func(err error) {
						// if err != nil {
						// 	fmt.Println(terror.Echo(err))
						// }
						cancel()
					})

					err := g.Run()
					if errors.Is(err, run.SignalError{Signal: os.Interrupt}) {
						err = terror.Warn(err)
						return err
					}
					return nil
				},
			},
		},
	}
	terror.Echo(app.Run(os.Args))
}

func SeedFunc(ctxCLI *cli.Context, ctx context.Context) error {
	databaseUser := ctxCLI.String("database_user")
	databasePass := ctxCLI.String("database_pass")
	databaseHost := ctxCLI.String("database_host")
	databasePort := ctxCLI.String("database_port")
	databaseName := ctxCLI.String("database_name")
	databaseAppName := ctxCLI.String("database_application_name")
	databaseIdleMaxConns := int(ctxCLI.Int("database_idle_max_conns"))
	databaseOpenMaxConns := int(ctxCLI.Int("database_open_max_conns"))

	// db connection
	conn, err := connectDB(databaseUser, databasePass, databaseHost, databasePort, databaseName, databaseAppName, Version, databaseIdleMaxConns, databaseOpenMaxConns)
	if err != nil {
		return terror.Panic(err)
	}

	fmt.Println("Seeding users")
	err = seedUsers(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func connectDB(
	databaseUser string,
	databasePass string,
	databaseHost string,
	databasePort string,
	databaseName string,
	databaseApplicationName string,
	apiVersion string,
	maxIdle int,
	maxOpen int,
) (*sql.DB, error) {
	params := url.Values{}
	params.Add("sslmode", "disable")

	if databaseApplicationName != "" {
		params.Add("application_name", fmt.Sprintf("%s %s", databaseApplicationName, apiVersion))
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		databaseUser,
		databasePass,
		databaseHost,
		databasePort,
		databaseName,
		params.Encode(),
	)

	config, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn := stdlib.OpenDB(*config)
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxOpen)

	return conn, nil
}

func seedUsers(ctx context.Context, conn *sql.DB) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	password := "devdev"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}

	// Insert superadmin
	superAdmin := &boiler.User{
		Name:        "Super Admin",
		Email:       null.StringFrom("superadmin@example.com"),
		Title:       null.StringFrom("Developer"),
		Description: null.StringFrom("Developer of renthome.com"),
		IsVerified:  true,
		Role:        "ADMIN",
	}

	err = superAdmin.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	superAdminPasswordHash := boiler.PasswordHash{
		PasswordHash: string(hash),
		UserID:       superAdmin.ID,
	}

	err = superAdminPasswordHash.Insert(tx, boil.Infer())
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
