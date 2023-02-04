package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"renthome/boiler"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

// Auther to handle JWT authentication
type Auther struct {
	TokenExpiryDays  int
	JWTSecret        string
	CookieSecure     bool
	GoogleClientID   string
	FacebookClientID string
	AppleClientID    string
}

func NewAuther(tokenExpiryDays int, jwtSecret string, cookieSecure bool, googleClientID string, facebookClientID string, appleClientID string) *Auther {
	result := &Auther{
		TokenExpiryDays:  tokenExpiryDays,
		JWTSecret:        jwtSecret,
		CookieSecure:     cookieSecure,
		GoogleClientID:   googleClientID,
		FacebookClientID: facebookClientID,
		AppleClientID:    appleClientID,
	}
	return result
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type FacebookLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type GoogleLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type AppleLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type EmailSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type FacebookSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type GoogleSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type AppleSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

func (api *APIController) EmailLoginHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()

	req := &EmailLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Unable to decode email login request.", http.StatusBadRequest)
	}

	if req.Email == "" {
		http.Error(w, "Email is required.", http.StatusBadRequest)
	}

	if req.Password == "" {
		http.Error(w, "Password is required.", http.StatusBadRequest)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		http.Error(w, "Something went wrong, unable to login", http.StatusInternalServerError)
	}

	fmt.Print(hash)

	// begin transaction
	//tx, err := api.Conn.BeginTx(ctx, nil)

	// user := boiler.User{
	// 	Email: req.Email,
	// }

	// passwordHash := boiler.PasswordHash{
	// 	PasswordHash: hash,
	// }

}

func (api *APIController) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &GoogleLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Unable to decode Google login request.", http.StatusBadRequest)
	}
}

func (api *APIController) FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &FacebookLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Unable to decode Facebook login request.", http.StatusBadRequest)
	}
}

func (api *APIController) AppleLoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &AppleLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Unable to decode Apple login request.", http.StatusBadRequest)
	}
}

func (api *APIController) EmailSignUpHandler(w http.ResponseWriter, r *http.Request) {
	const genericErrMsg = "Something went wrong, unable to signup."

	req := &EmailSignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Unable to decode email signup request.", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		http.Error(w, "Email is required.", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		http.Error(w, "Password is required.", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		log.Println(err)
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}

	user := &boiler.User{
		Email: null.StringFrom(req.Email),
	}
	err = user.Insert(api.Conn, boil.Infer())
	if err != nil {
		log.Println(err)
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}

	passwordHash := &boiler.PasswordHash{
		PasswordHash: string(hash),
	}
	err = passwordHash.Insert(api.Conn, boil.Infer())
	if err != nil {
		log.Println(err)
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		http.Error(w, genericErrMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *APIController) FacebookSignUpHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *APIController) GoogleSignUpHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *APIController) AppleSignUpHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *APIController) forgetPasswordHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *APIController) changePasswordHandler(w http.ResponseWriter, r *http.Request) {

}
