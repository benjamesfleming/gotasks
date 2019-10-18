package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/benjamesfleming/gotasks/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

var jwtKey = []byte(os.Getenv("SESSION_SECRET"))
var jwtMethod = jwt.SigningMethodHS256

var jwtAccessTTL = 5 * time.Minute
var jwtRefreshTTL = 24 * time.Hour * 7

// Claims ...
// The claims in a given jwt token
type Claims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.StandardClaims
}

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
	)
}

// AuthCallback completes the provider based authentication flow
// and return a JWT token to the user
func AuthCallback(c buffalo.Context) error {
	// Complete the user authentication
	// and get user data
	data, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Check if the user is already signed up
	user := new(models.User)
	if err := tx.Where("email = ?", data.Email).First(user); err != nil {
		return fmt.Errorf("[ERROR]failed to query database: \n\n%s", err)
	}

	if user == nil {
		// Map the providers data to a User model
		user.UserName = data.Name
		user.Provider = data.Provider
		user.Email = data.Email

		// Validate the user data
		// if there are any unknown errors then break
		verrs, err := tx.ValidateAndCreate(user)
		if err != nil {
			return err
		}

		// Check for any validataion errors
		// if there are any return them in a 301 request to an error page
		if verrs.HasAny() {
			return c.Redirect(301, "/#/auth/error")
		}
	}

	// Generate the JWT tokens
	refreshClaims := jwt.StandardClaims{ExpiresAt: time.Now().Add(jwtRefreshTTL).Unix()}
	accessClaims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtAccessTTL).Unix(),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwtMethod, refreshClaims).SignedString(jwtKey)
	accessToken, err := jwt.NewWithClaims(jwtMethod, accessClaims).SignedString(jwtKey)

	// Check for JWT errors
	if err != nil {
		return fmt.Errorf("failed to generate tokens")
	}

	// Redirect the user to the clientside
	// to complete the signup process
	return c.Redirect(301, fmt.Sprintf("/#/auth/complete?access_token=%s&refresh_token=%s", accessToken, refreshToken))
}
