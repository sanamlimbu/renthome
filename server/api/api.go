package api

import (
	"database/sql"
	"renthome"
	"renthome/email"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
}

type APIController struct {
	Addr          string
	Auther        *Auther
	Mailer        *email.Mailer
	Conn          *sql.DB
	Validator     *Validator
	ObjectStorage *renthome.ObjectStorage
}

func NewAPIController(mailer *email.Mailer, addr string, auther *Auther, conn *sql.DB, validator *Validator, objectStorage *renthome.ObjectStorage) *APIController {
	api := &APIController{
		Addr:          addr,
		Conn:          conn,
		Mailer:        mailer,
		Auther:        auther,
		Validator:     validator,
		ObjectStorage: objectStorage,
	}

	return api
}

func NewRouter(api *APIController, adminHostURL, publicHostURL, agentHostURL string) *chi.Mux {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{adminHostURL, publicHostURL, agentHostURL},
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
				r.Post("/facebook", WithError(api.FacebookAuthHandler))
				r.Post("/google", WithError(api.GoogleAuthHandler))
				r.Post("/forgot-password", WithError(api.ForgotPasswordHandler))
				r.Post("/forgot-password-confirm", WithError(api.ConfirmForgotPasswordHandler))
				r.Post("/change-password", WithError(WithUser(api, api.ChangePasswordHandler)))
				r.Post("/logout", WithError(api.LogoutHandler))
				r.Post("/change-email", WithError(WithUser(api, api.ChangeEmailHandler)))
				r.Post("/signout-all", WithError(WithUser(api, api.SignoutAllDevicesHandler)))
			})
		})
		r.Get("/properties", WithError(api.GetProperties))
		r.Get("/properties/{id}", WithError(api.GetProperty))
		r.Post("/properties", WithError(api.CreateProperty))
		r.Put("/properties/{id}", WithError(WithUser(api, api.UpdateProperty)))
		r.Delete("/properties/{id}", WithError(WithUser(api, api.DeleteProperty)))

		r.Post("/notifications", WithError(api.GetNotifications))
		r.Post("/privacies", WithError(api.GetPrivacies))
		r.Put("/notifications/update", WithError(WithUser(api, api.UpdateNotificationHandler)))
		r.Put("/privacies/update", WithError(WithUser(api, api.UpdatePrivacyHandler)))

		r.Post("/files/upload", WithError(WithUser(api, api.FileUpload)))

		r.Post("/locations", WithError(api.GetLocations))
	})

	return r
}
