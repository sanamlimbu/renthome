package api

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"renthome/boiler"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

const GoogleOAuthTokenURL = "https://oauth2.googleapis.com/token"
const GoogleOAuthUserURL = "https://www.googleapis.com/oauth2/v1/userinfo"
const GoogleOAuthRedirectURI = "http://localhost:8000/api/auth/google"

const ErrDecodeJSONPayload = "Unable to decode JSON payload."
const ErrBeginTransaction = "Unable to begin transaction."
const ErrCommitTransaction = "Unable to commit transaction."
const ErrEncodeJSONPayload = "Unable to encode JSON payload."
const ErrJWTAccessToken = "Unable to generate JWT access token."
const ErrPasswordHash = "Unable to generate password hash."
const ErrSomethingWentWrong = "Something went wrong, please try again."
const ErrMissAuthToken = "Missing authorization token."
const ErrInvalidToken = "Invalid token."
const ErrBadRequest = "Bad request, please try again."
const ErrUnauthorised = "You are not authorised to perform this action."

type Role string

const (
	Member  Role = "MEMBER"
	Manager Role = "MANAGER"
	Admin   Role = "ADMIN"
)

// Auther to handle JWT authentication
type Auther struct {
	TokenExpiryDays  int
	JWTSecretByte    []byte
	CookieSecure     bool
	FacebookClientID string
	AppleClientID    string
	GoogleClientID   string
}

func NewAuther(tokenExpiryDays int, jwtSecret string, cookieSecure bool, googleClientID string, facebookClientID string, appleClientID string) *Auther {
	result := &Auther{
		TokenExpiryDays:  tokenExpiryDays,
		JWTSecretByte:    []byte(jwtSecret),
		CookieSecure:     cookieSecure,
		FacebookClientID: facebookClientID,
		AppleClientID:    appleClientID,
		GoogleClientID:   googleClientID,
	}
	return result
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailLoginResponse struct {
	User  boiler.User `json:"user"`
	Token string      `json:"token"`
}

type EmailSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailSignUpResponse struct {
	User  boiler.User `json:"user"`
	Token string      `json:"token"`
}

type GoogleUser struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleAuthResponse struct {
	User  boiler.User `json:"user"`
	Token string      `json:"token"`
}

type FacebookUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type FacebookAuthResponse struct {
	User  boiler.User `json:"user"`
	Token string      `json:"token"`
}

func (api *APIController) EmailLoginHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &EmailLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	if req.Email == "" {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("email is required"), "Email is required.")
	}

	if req.Password == "" {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("password is required"), "Password is required.")
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	user, err := boiler.Users(boiler.UserWhere.Email.EQ(null.StringFrom(strings.ToLower(req.Email)))).One(tx)
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("user not available"), "User not available, please sign up.")
	}

	if user == nil {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("user not available"), "User not available, please sign up.")
	}

	passwordHash, err := boiler.FindPasswordHash(tx, user.ID)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Wrong password, please try again.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash.PasswordHash), []byte(req.Password))
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Wrong password, please try again.")
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	resp := &EmailLoginResponse{
		User:  *user,
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusOK, nil

}

func (api *APIController) EmailSignUpHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &EmailSignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	if req.Email == "" {
		return http.StatusBadRequest, terror.Error(err, "Email is required.")
	}

	if req.Password == "" {
		return http.StatusBadRequest, terror.Error(err, "Password is required.")
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	// user with email already exists
	user, err := boiler.Users(
		boiler.UserWhere.Email.EQ(null.StringFrom(strings.ToLower(req.Email))),
	).One(tx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	if user != nil {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("user already exists"), "User already exists, please login.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrPasswordHash)
	}

	user = &boiler.User{
		Email: null.StringFrom(req.Email),
		Role:  "MEMBER",
	}

	err = user.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to create user.")
	}

	passwordHash := &boiler.PasswordHash{
		PasswordHash: string(hash),
		UserID:       user.ID,
	}
	err = passwordHash.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to create password hash.")
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	// insert default notification settings
	notifications, err := boiler.Notifications().All(tx)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	for _, notification := range notifications {
		userNotification := &boiler.UserNotification{
			UserID:         user.ID,
			NotificationID: notification.ID,
		}
		err = userNotification.Insert(tx, boil.Infer())
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}
	}

	// insert default privacies
	privacies, err := boiler.Privacies().All(tx)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	for _, privacy := range privacies {
		userPrivacy := &boiler.UserPrivacy{
			UserID:    user.ID,
			PrivacyID: privacy.ID,
		}

		err = userPrivacy.Insert(tx, boil.Infer())
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}
	}

	resp := &EmailSignUpResponse{
		User:  *user,
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusCreated, nil
}

// GoogleAuthHandler handles Google login and signup
func (api *APIController) GoogleAuthHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GoogleUser{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	userID, err := uuid.NewV4()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	user := &boiler.User{}
	// check user with Google ID
	userAlt, err := boiler.Users(boiler.UserWhere.GoogleID.EQ(null.StringFrom(req.Sub))).One(tx)
	// no user
	if errors.Is(err, sql.ErrNoRows) {
		user.ID = userID.String()
		user.GoogleID = null.StringFrom(req.Sub)
		user.Email = null.StringFrom(strings.ToLower(req.Email))
		user.Role = "MEMBER"
		user.Name = req.Name
		user.IsVerified = true

		err := user.Insert(tx, boil.Infer())
		if err != nil {
			return http.StatusBadRequest, terror.Error(err, "Unable to create user.")
		}

		// insert default notification settings
		notifications, err := boiler.Notifications().All(tx)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, notification := range notifications {
			userNotification := &boiler.UserNotification{
				UserID:         user.ID,
				NotificationID: notification.ID,
			}
			err = userNotification.Insert(tx, boil.Infer())
			if err != nil {
				return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
			}
		}

		// insert default privacies
		privacies, err := boiler.Privacies().All(tx)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, privacy := range privacies {
			userPrivacy := &boiler.UserPrivacy{
				UserID:    user.ID,
				PrivacyID: privacy.ID,
			}

			err = userPrivacy.Insert(tx, boil.Infer())
			if err != nil {
				return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
			}
		}

	}

	if userAlt != nil {
		user = userAlt
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	resp := &GoogleAuthResponse{
		User:  *user,
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusCreated, nil

}

type LogoutRequest struct {
	UserID string `json:"user_id"`
}

// handles logout operation
func (api *APIController) LogoutHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &LogoutRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	userID, err := GetUserIDFromToken(api, r)
	if err != nil || userID == "" {
		return http.StatusUnauthorized, terror.Error(fmt.Errorf("invalid token"), ErrInvalidToken)
	}

	if req.UserID != userID {
		return http.StatusUnauthorized, terror.Error(fmt.Errorf("unauthorized action"), "You are not authorized to perform this action.")
	}

	return http.StatusOK, nil
}

// FacebookAuthHandler handles Google login and signup
func (api *APIController) FacebookAuthHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &FacebookUser{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	userID, err := uuid.NewV4()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	user := &boiler.User{}
	// check user with Facebook ID
	userAlt, err := boiler.Users(boiler.UserWhere.FacebookID.EQ(null.StringFrom(req.ID))).One(tx)
	// no user
	if errors.Is(err, sql.ErrNoRows) {
		user.ID = userID.String()
		user.FacebookID = null.StringFrom(req.ID)
		user.Email = null.StringFrom(strings.ToLower(req.Email))
		user.Role = "MEMBER"
		user.Name = req.Name
		user.IsVerified = true

		err := user.Insert(tx, boil.Infer())
		if err != nil {
			return http.StatusBadRequest, terror.Error(err, "Unable to create user.")
		}

		// insert default notification settings
		notifications, err := boiler.Notifications().All(tx)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, notification := range notifications {
			userNotification := &boiler.UserNotification{
				UserID:         user.ID,
				NotificationID: notification.ID,
			}
			err = userNotification.Insert(tx, boil.Infer())
			if err != nil {
				return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
			}
		}

		// insert default privacies
		privacies, err := boiler.Privacies().All(tx)
		if err != nil {
			return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
		}

		for _, privacy := range privacies {
			userPrivacy := &boiler.UserPrivacy{
				UserID:    user.ID,
				PrivacyID: privacy.ID,
			}

			err = userPrivacy.Insert(tx, boil.Infer())
			if err != nil {
				return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
			}
		}

	}

	if userAlt != nil {
		user = userAlt
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	resp := &FacebookAuthResponse{
		User:  *user,
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusCreated, nil
}

func (api *APIController) AppleLoginHandler(w http.ResponseWriter, r *http.Request) {
	// req := &AppleLoginRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	http.Error(w, "Unable to decode Apple login request.", http.StatusBadRequest)
	// }
}

func (api *APIController) FacebookSignUpHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *APIController) AppleSignUpHandler(w http.ResponseWriter, r *http.Request) {

}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordResponse struct {
	ResetToken string `json:"reset_token"`
}

func (api *APIController) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &ForgotPasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	user, err := boiler.Users(boiler.UserWhere.Email.EQ(null.StringFrom(strings.ToLower(req.Email)))).One(api.Conn)
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("user not available"), "Could not reset password for the account, please contact support or try again.")
	}

	// Generate a random number between 0 and 999999 inclusive
	randInt, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}
	code := fmt.Sprintf("%06d", randInt)

	resetPassword := &boiler.ResetPassword{
		UserID:    user.ID,
		UpdatedAt: time.Now(),
		Code:      code,
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	err = resetPassword.Upsert(tx, true, []string{"user_id"}, boil.Whitelist("code", "updated_at"), boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	err = api.Mailer.SendAccountVerificationCode(user.Email.String, user.Name, resetPassword.Code)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to send reset code, please try again.")
	}

	resp := &ForgotPasswordResponse{
		ResetToken: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusOK, nil
}

type ConfirmForgotPasswordRequest struct {
	Code            string `json:"code"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ConfirmForgotPasswordResponse struct {
	User  *boiler.User `json:"user"`
	Token string       `json:"token"`
}

func (api *APIController) ConfirmForgotPasswordHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &ConfirmForgotPasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	user, err := GetUserFromToken(api, r)
	if err != nil {
		return http.StatusUnauthorized, terror.Error(err, ErrUnauthorised)
	}

	resetPassword, err := boiler.FindResetPassword(tx, user.ID)
	if err != nil {
		return http.StatusUnauthorized, terror.Error(err, ErrUnauthorised)
	}

	if resetPassword.Code != req.Code {
		return http.StatusUnauthorized, terror.Error(fmt.Errorf("reset code did not match"), "Code did not match, please try again.")
	}

	if req.NewPassword != req.ConfirmPassword {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("password did not match"), "Password did not match, please try again.")
	}

	passwordHash, err := boiler.FindPasswordHash(tx, user.ID)
	if err != nil {
		return http.StatusUnauthorized, terror.Error(err, ErrUnauthorised)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 8)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrPasswordHash)
	}

	passwordHash.PasswordHash = string(hash)
	passwordHash.UpdatedAt = time.Now()

	_, err = passwordHash.Update(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	token, claims, err := api.GenerateJWTAccessToken(user, api.Auther.JWTSecretByte)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrJWTAccessToken)
	}

	issueToken := &boiler.IssueToken{
		ID:        claims.StandardClaims.Id,
		UserID:    claims.StandardClaims.Subject,
		Device:    r.Header.Get("X-User-Agent"),
		CreatedAt: time.Unix(claims.StandardClaims.IssuedAt, 0),
		ExpiresAt: time.Unix(claims.StandardClaims.ExpiresAt, 0),
	}

	err = issueToken.Insert(tx, boil.Infer())
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	resp := &ConfirmForgotPasswordResponse{
		User:  user,
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrEncodeJSONPayload)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusOK, nil
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func (api *APIController) ChangePasswordHandler(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &ChangePasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	passwordHash, err := boiler.FindPasswordHash(api.Conn, user.ID)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Wrong password, please try again.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash.PasswordHash), []byte(req.CurrentPassword))
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Wrong password, please try again.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 8)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrPasswordHash)
	}

	passwordHash.PasswordHash = string(hash)
	passwordHash.UpdatedAt = time.Now()

	_, err = passwordHash.Update(api.Conn, boil.Infer())
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to update password.")
	}

	return http.StatusOK, nil
}

type ChangeEmailHandlerRequest struct {
	NewEmail string `json:"new_email"`
}

func (api *APIController) ChangeEmailHandler(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {
	req := &ChangeEmailHandlerRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, ErrDecodeJSONPayload)
	}

	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	// email already exists
	usr, err := boiler.Users(
		boiler.UserWhere.Email.EQ(null.StringFrom(strings.ToLower(req.NewEmail))),
	).One(tx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	if usr != nil {
		return http.StatusBadRequest, terror.Error(fmt.Errorf("email already exists"), "Email address already teken.")
	}

	return http.StatusOK, nil
}

type Claims struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	jwt.StandardClaims
}

// GenerateJWTAccessToken generates access token which expires in 24 hours
func (api *APIController) GenerateJWTAccessToken(user *boiler.User, jwtSecret []byte) (string, *Claims, error) {
	tokenID, err := uuid.NewV4()
	if err != nil {
		return "", nil, err
	}

	claims := &Claims{
		Email:      user.Email.String,
		Name:       user.Name,
		IsVerified: user.IsVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "go-jwt-auth",
			IssuedAt:  time.Now().Unix(),
			Id:        tokenID.String(),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims,
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return tokenString, claims, nil
}

// GenerateResetPasswordToken generates token for reset password which expires in a hour
func GenerateResetPasswordToken(tokenID string, userID string, jwtSecret []byte) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "go-jwt-auth",
			IssuedAt:  time.Now().Unix(),
			Id:        tokenID,
			Subject:   userID,
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims,
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetJWTAccessToken returns access token from auth header
func GetJWTAccessToken(authHeader string) (string, error) {

	// Split the Authorization header into its parts
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}
	// Get the token from the Authorization header
	token := parts[1]

	if token == "" {
		return "", terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}

	return token, nil
}

// VerifyJWTAccessToken verifies JWT access token
func VerifyJWTAccessToken(tokenString string, jwtSecret []byte) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, terror.Error(fmt.Errorf("invalid signing method"), "Invalid signing method.")
		}
		return jwtSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, terror.Error(fmt.Errorf("invalid token signature"), "Invalid token signature.")
		}
		return false, err
	}

	// Verify the token expiry
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return false, terror.Error(fmt.Errorf("token has expired"), "Token has expired.")
		}
		return true, nil
	} else {
		return false, terror.Error(fmt.Errorf("invalid token claims"), "Invalid token claims.")
	}

}

// GetUserIDFromToken checks validity of token (including not blacklisted) and returns user id
func GetUserIDFromToken(api *APIController, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}

	tokenString, err := GetJWTAccessToken(authHeader)
	if err != nil {
		return "", terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}

	// Parse the token and extract the claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, terror.Error(fmt.Errorf("invalid signing method"), "Invalid signing method.")
		}
		return api.Auther.JWTSecretByte, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", terror.Error(fmt.Errorf("invalid token signature"), "Invalid token signature.")
		}
		return "", err
	}

	// Verify the token expiry and return user id
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return "", terror.Error(fmt.Errorf("token has expired"), "Token has expired.")
		}
		return claims.StandardClaims.Subject, nil
	}

	return "", terror.Error(fmt.Errorf("invalid token claims"), "Invalid token claims.")

}

// GetUserFromToken checks validity of token (including not blacklisted) and returns user
func GetUserFromToken(api *APIController, r *http.Request) (*boiler.User, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}

	tokenString, err := GetJWTAccessToken(authHeader)
	if err != nil {
		return nil, terror.Error(fmt.Errorf("missing authorization token"), ErrMissAuthToken)
	}

	// Parse the token and extract the claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, terror.Error(fmt.Errorf("invalid signing method"), "Invalid signing method.")
		}
		return api.Auther.JWTSecretByte, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, terror.Error(fmt.Errorf("invalid token signature"), "Invalid token signature.")
		}
		return nil, err
	}

	// Verify the token expiry and returns user
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, terror.Error(fmt.Errorf("token has expired"), "Token has expired.")
		}

		user, err := boiler.FindUser(api.Conn, claims.StandardClaims.Subject)
		if err != nil {
			return nil, terror.Error(err, "Cannot find user.")
		}

		// token blacklisted?
		issuedToken, err := boiler.IssueTokens(boiler.IssueTokenWhere.ID.EQ(claims.StandardClaims.Id), boiler.IssueTokenWhere.UserID.EQ(user.ID)).One(api.Conn)
		if err != nil {
			return nil, terror.Error(err, "Invalid token.")
		}

		if issuedToken == nil {
			return nil, terror.Error(err, "Invalid token.")
		}

		if issuedToken.Blacklisted {
			return nil, terror.Error(err, "Invalid token.")
		}

		return user, nil
	}

	return nil, terror.Error(fmt.Errorf("invalid token claims"), "Invalid token claims.")

}

// SignoutAllDevicesHandler signs out user from all devices
func (api *APIController) SignoutAllDevicesHandler(w http.ResponseWriter, r *http.Request, user *boiler.User) (int, error) {

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrBeginTransaction)
	}

	_, err = boiler.IssueTokens(boiler.IssueTokenWhere.UserID.EQ(user.ID)).UpdateAll(tx, boiler.M{"blacklisted": true})
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrSomethingWentWrong)
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, ErrCommitTransaction)
	}

	return http.StatusOK, nil
}
