package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"renthome/api"
	"renthome/email"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/ninja-software/terror/v2"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// Build Version
const Version = "v1.0.0"

func main() {

	app := &cli.App{
		Compiled: time.Now(),
		Usage:    "Run the renthome server",
		Authors: []*cli.Author{
			{
				Name:  "Sanam Limbu",
				Email: "sudosanam@gmail.com",
			},
		},
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Show version",
				Action: func(c *cli.Context) error {
					fmt.Println(Version)
					return nil
				},
			},
			{
				Name:    "serve",
				Aliases: []string{"s"},
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

					&cli.StringFlag{Name: "api_addr", Value: ":8000", EnvVars: []string{"RENTHOME_API_ADDR"}, Usage: "host:port to run the API"},
					&cli.StringFlag{Name: "rootpath", Value: "../web/dist", EnvVars: []string{"RENTHOME_ROOTPATH"}, Usage: "folder path of index.html"},
					&cli.StringFlag{Name: "jwtsecret", Value: "a35eab71-f691-4dc3-98e5-980bda774fa0", EnvVars: []string{"RENTHOME_USERAUTH_JWTSECRET"}, Usage: "JWT secret"},
					&cli.StringFlag{Name: "google_client_id", Value: "", EnvVars: []string{"RENTHOME_GOOGLE_CLIENT_ID"}, Usage: "Google Client ID for OAuth functionaility."},
					&cli.StringFlag{Name: "facebook_client_id", Value: "", EnvVars: []string{"RENTHOME_FACEBOOK_CLIENT_ID"}, Usage: "Facebook Client ID for OAuth functionaility."},
					&cli.BoolFlag{Name: "cookie_secure", Value: false, EnvVars: []string{"RENTHOME_COOKIE_SECURE"}, Usage: "Cookie Secure setting option for secure cookies."},

					&cli.StringFlag{Name: "mail_host", Value: "smtp.gmail.com", EnvVars: []string{"RENTHOME_MAIL_HOST"}, Usage: "Gmail SMTP host address"},
					&cli.StringFlag{Name: "mail_port", Value: "587", EnvVars: []string{"RENTHOME_MAIL_PORT"}, Usage: "Mailtrap SMTP port"},
					&cli.StringFlag{Name: "mail_username", Value: "", EnvVars: []string{"RENTHOME_MAIL_USERNAME"}, Usage: "Gmail email address"},
					&cli.StringFlag{Name: "mail_password", Value: "", EnvVars: []string{"RENTHOME_MAIL_PASSWORD"}, Usage: "Gmail app specific password"},

					&cli.StringFlag{Name: "admin_host_url", Value: "http://localhost:3001", EnvVars: []string{"RENTHOME_ADMIN_FRONTEND_HOST_URL"}, Usage: "The Admin Site URL used for links in the mailer and allowed cors urls"},
					&cli.StringFlag{Name: "public_host_url", Value: "http://localhost:3000", EnvVars: []string{"RENTHOME_PUBLIC_FRONTEND_HOST_URL"}, Usage: "The Public Site URL used for links in the mailer and allowed cors urls"},

					&cli.IntFlag{Name: "tokenexpirydays", Value: 30, EnvVars: []string{"RENTHOME_USERAUTH_TOKENEXPIRYDAYS"}, Usage: "How many days before the token expires"},
					&cli.IntFlag{Name: "blacklistrefreshhours", Value: 1, EnvVars: []string{"RENTHOME_USERAUTH_BLACKLISTREFRESHHOURS"}, Usage: "How often should the issued_tokens list be cleared of expired tokens in hours"},
				},

				Usage: "run server",
				Action: func(c *cli.Context) error {
					ctx, cancel := context.WithCancel(c.Context)
					log.Info().Msg("zerolog initialised")

					g := &run.Group{}
					// Listen for os.interrupt
					g.Add(run.SignalHandler(ctx, os.Interrupt))
					// start the server
					g.Add(func() error { return ServeFunc(c, ctx) }, func(err error) {
						if err != nil {
							fmt.Println(terror.Echo(err))
						}
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
	terror.AppVersion = Version
	terror.Echo(app.Run(os.Args))
}

func ServeFunc(ctxCLI *cli.Context, ctx context.Context) error {
	apiAddr := ctxCLI.String("api_addr")
	tokenExpiryDays := ctxCLI.Int("tokenexpirydays")
	jwtSecret := ctxCLI.String("jwtsecret")
	cookieSecure := ctxCLI.Bool("cookie_secure")

	googleClientID := ctxCLI.String("google_client_id")
	facebookClientID := ctxCLI.String("facebook_client_id")

	databaseUser := ctxCLI.String("database_user")
	databasePass := ctxCLI.String("database_pass")
	databaseHost := ctxCLI.String("database_host")
	databasePort := ctxCLI.String("database_port")
	databaseName := ctxCLI.String("database_name")
	databaseAppName := ctxCLI.String("database_application_name")
	databaseMaxIdleConns := int(ctxCLI.Int("database_max_idle_conns"))
	databaseMaxOpenConns := int(ctxCLI.Int("database_max_open_conns"))

	adminHostURL := ctxCLI.String("admin_host_url")
	publicHostURL := ctxCLI.String("public_host_url")

	mailHost := ctxCLI.String("mail_host")
	mailPort := ctxCLI.String("mail_port")
	mailUsername := ctxCLI.String("mail_username")
	mailPassword := ctxCLI.String("mail_password")

	mailer, err := email.NewMailer(mailUsername, mailPassword, mailHost, mailPort)
	if err != nil {
		panic(err)
	}

	auther := api.NewAuther(tokenExpiryDays, jwtSecret, cookieSecure, googleClientID, facebookClientID)

	// db connection
	conn, err := connectDB(databaseUser, databasePass, databaseHost, databasePort, databaseName, databaseAppName, Version, databaseMaxIdleConns, databaseMaxOpenConns)
	if err != nil {
		return terror.Panic(err)
	}

	apiController := api.NewAPIController(mailer, apiAddr, auther, conn)

	router := api.NewRouter(apiController, adminHostURL, publicHostURL)

	server := &http.Server{
		Addr:    apiController.Addr,
		Handler: router,
	}

	return server.ListenAndServe()
}

func connectDB(
	databaseUser string,
	databasePass string,
	databaseHost string,
	databasePort string,
	databaseName string,
	databaseApplicationName string,
	apiVersion string,
	maxIdleConns int,
	maxOpenConns int,
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

	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetMaxOpenConns(maxOpenConns)

	return conn, nil
}
