package api

import (
	"database/sql"
	"renthome/email"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type APIController struct {
	Addr   string
	Auther *Auther
	Mailer *email.Mailer
	Conn   *sql.DB
}

func NewAPIController(mailer *email.Mailer, addr string, auther *Auther, conn *sql.DB) *APIController {
	api := &APIController{
		Addr:   addr,
		Conn:   conn,
		Mailer: mailer,
		Auther: auther,
	}

	return api
}

func NewRouter(api *APIController, adminHostURL, publicHostURL string) *chi.Mux {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{adminHostURL, publicHostURL},
		AllowCredentials: true,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {

	})
	r.Route(("/auth"), func(r chi.Router) {
		r.Post("/email-login", api.EmailLoginHandler)
		r.Post("/facebook-login", api.FacebookLoginHandler)
		r.Post("/apple-login", api.AppleLoginHandler)
		r.Post("/google-login", api.GoogleLoginHandler)
		r.Post("/email-signup", api.EmailSignUpHandler)
		r.Post("/facebook-signup", api.FacebookSignUpHandler)
		r.Post("/google-signup", api.GoogleSignUpHandler)
		r.Post("/apple-signup", api.AppleSignUpHandler)
		r.Post("/forget-password", api.forgetPasswordHandler)
		r.Post("/change-password", api.changePasswordHandler)
	})

	return r
}
