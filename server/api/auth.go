package api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"renthome/boiler"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ninja-software/terror/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

const GoogleOAuthTokenRootURL = "https://oauth2.googleapis.com/token"
const GoogleOAuthUserInfoRootURL = "https://www.googleapis.com/oauth2/v1/userinfo"

// Auther to handle JWT authentication
type Auther struct {
	TokenExpiryDays  int
	JWTSecretByte    []byte
	CookieSecure     bool
	FacebookClientID string
	AppleClientID    string
	GoogleConfig     *GoogleConfig
}

type GoogleConfig struct {
	GoogleClientID          string
	GoogleClientSecret      string
	GoogleRedirectURILogin  string
	GoogleRedirectURISignup string
}

func NewAuther(tokenExpiryDays int, jwtSecret string, cookieSecure bool, googleConfig *GoogleConfig, facebookClientID string, appleClientID string) *Auther {
	result := &Auther{
		TokenExpiryDays:  tokenExpiryDays,
		JWTSecretByte:    []byte(jwtSecret),
		CookieSecure:     cookieSecure,
		FacebookClientID: facebookClientID,
		AppleClientID:    appleClientID,
		GoogleConfig:     googleConfig,
	}
	return result
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type EmailSignUpRequest struct {
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

type GoogleOAuthToken struct {
	AccessToken string
	TokenID     string
}

type GoogleUser struct {
	ID         string
	Email      string
	Verified   bool
	Name       string
	GivenName  string
	FamilyName string
	Picture    string
	Locale     string
}
type GoogleAuthRequest struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

// GoogleAuthHandler handles Google login and signup
func (api *APIController) GoogleAuthHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	// get Google authorization code
	// code := r.URL.Query().Get("code")
	// if code == "" {
	// 	http.Error(w, "Missing authorization code.", http.StatusUnauthorized)
	// 	return
	// }

	req := &GoogleAuthRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to parse JSON payload.")
	}
	fmt.Println(req.Code)
	fmt.Println(req.RedirectURI)

	accessToken, err := GetGoogleOAuthToken(
		api.Auther.GoogleConfig.GoogleClientID,
		api.Auther.GoogleConfig.GoogleClientSecret,
		req.RedirectURI,
		req.Code)

	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to retrieve Google OAuth token.")
	}

	googleUser, err := GetGoogleUser(accessToken)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to retrieve Google user info.")
	}

	// begin transaction
	ctx := context.Background()
	tx, err := api.Conn.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to begin transaction.")
	}

	// check user with Google ID
	user, err := boiler.Users(
		boiler.UserWhere.Email.EQ(null.StringFrom(googleUser.ID))).One(tx)

	if err == sql.ErrNoRows {
		user.GoogleID = null.StringFrom(googleUser.ID)
		user.Email = null.StringFrom(strings.ToLower(googleUser.Email))
		user.Role = "MEMBER"
		user.Name = googleUser.Name
		user.IsVerified = true

		err := user.Insert(tx, boil.Infer())
		if err != nil {
			return http.StatusBadRequest, terror.Error(err, "Unable to create user.")
		}
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to commit transaction.")
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to encode JSON.")
	}

	return http.StatusCreated, nil

}

func (api *APIController) FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	// req := &FacebookLoginRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	http.Error(w, "Unable to decode Facebook login request.", http.StatusBadRequest)
	// }
}

func (api *APIController) AppleLoginHandler(w http.ResponseWriter, r *http.Request) {
	// req := &AppleLoginRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	http.Error(w, "Unable to decode Apple login request.", http.StatusBadRequest)
	// }
}

func (api *APIController) EmailSignUpHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &EmailSignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to parse JSON payload.")
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
		return http.StatusInternalServerError, terror.Error(err, "Unable to begin transaction.")
	}

	// user with email already exists
	user, _ := boiler.Users(
		boiler.UserWhere.Email.EQ(null.StringFrom(strings.ToLower(req.Email))),
	).One(tx)

	if user != nil {
		return http.StatusBadRequest, terror.Error(err, "User already exists, please login.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return http.StatusBadRequest, terror.Error(err, "Unable to generate password hash.")
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

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to commit transaction.")
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return http.StatusInternalServerError, terror.Error(err, "Unable to encode JSON.")
	}

	return http.StatusCreated, nil
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

// GetGoogleOAuthToken retrieves access token, this token grants temporary, limited access to a Google API on behalf of an end-user
func GetGoogleOAuthToken(clientID string, clientSecret string, redirectURI string, code string) (*GoogleOAuthToken, error) {
	values := url.Values{}

	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	values.Add("redirect_uri", redirectURI)

	query := values.Encode()

	req, err := http.NewRequest("POST", GoogleOAuthTokenRootURL, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(res.StatusCode)

	if res.StatusCode != http.StatusOK {
		return nil, terror.Error(fmt.Errorf("Unable to retrieve Google access token"))
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	token := &GoogleOAuthToken{
		AccessToken: GoogleOauthTokenRes["access_token"].(string),
		TokenID:     GoogleOauthTokenRes["id_token"].(string),
	}

	return token, nil
}

// GetGoogleUser retireves user informatiom from Google API using the access token
func GetGoogleUser(token *GoogleOAuthToken) (*GoogleUser, error) {
	url := fmt.Sprintf("%s?alt=json&access_token=%s", GoogleOAuthUserInfoRootURL, token.AccessToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.TokenID))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleUserRes); err != nil {
		return nil, err
	}

	googleUser := &GoogleUser{
		ID:        GoogleUserRes["id"].(string),
		Email:     GoogleUserRes["email"].(string),
		Verified:  GoogleUserRes["verified_email"].(bool),
		Name:      GoogleUserRes["name"].(string),
		GivenName: GoogleUserRes["given_name"].(string),
		Picture:   GoogleUserRes["picture"].(string),
		Locale:    GoogleUserRes["locale"].(string),
	}

	return googleUser, nil
}

// func (api *APIController) GetUserByGoogleID() (*boiler.User, error) {
// 	boiler.FindUser(api.Conn,)
// }

func GenerateJWTAccessToken(user boiler.User, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": user.ID,
		},
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
