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
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT"},
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route(("/auth"), func(r chi.Router) {
				r.Post("/login", WithError(api.EmailLoginHandler))
				r.Post("/signup", WithError(api.EmailSignUpHandler))
				r.Post("/facebook", api.FacebookLoginHandler)
				r.Post("/google", WithError(api.GoogleAuthHandler))
				r.Post("/apple", api.AppleLoginHandler)
				r.Post("/forget-password", api.forgetPasswordHandler)
				r.Post("/change-password", WithError(WithUser(api, api.ChangePasswordHandler)))
				r.Post("/logout", WithError(api.LogoutHandler))
			})
		})
		r.Post("/notifications", WithError(api.GetNotificationsHandler))
		r.Post("/privacies", WithError(api.GetPrivaciesHandler))
		r.Put("/notifications/update", WithError(WithUser(api, api.UpdateNotificationHandler)))
		r.Put("/privacies/update", WithError(WithUser(api, api.UpdatePrivacyHandler)))
	})

	return r
}
