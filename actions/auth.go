package actions

import (
	"fmt"
	"os"

	"github.com/benjamesfleming/gotasks/models"
	"github.com/benjamesfleming/gotasks/utils/jwt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

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
		fmt.Println(err)
		return c.Render(500, r.String("server failed to query database"))
	}

	if user == nil {
		// Map the providers data to a User model
		user.UserName = data.Name
		user.Provider = data.Provider
		user.Email = data.Email
		user.Privileges = strings.Join(
			[]string{""}, "\n",
		)

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

	// Generate a new JWT token pair
	tkns, _ := jwt.GenerateTokenPair(user)

	// Redirect the user to the clientside
	// to complete the signup process
	return c.Redirect(301, fmt.Sprintf("/#/auth/complete?access_token=%s&refresh_token=%s", tkns.AccessToken, tkns.RefreshToken))
}

// RefreshToken ...
func RefreshToken(c buffalo.Context) error {
	claims, err := jwt.ClaimsFromHeader(c)
	if err != nil {
		return c.Render(400, r.String("invalid token or malformed request"))
	}

	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(500, r.String("server failed to connect to database"))
	}

	// Grab the user from the database
	user := new(models.User)
	if err := tx.Where("id = ?", claims.StandardClaims.Id).First(user); err != nil {
		fmt.Println(err)
		return c.Render(500, r.String("server failed to query database"))
	}

	// Generate a new JWT token pair
	tkns, _ := jwt.GenerateTokenPair(user)
	data := struct {
		AccessToken string `json:"access_token"`
	}{
		tkns.AccessToken,
	}

	// Return the access token to the user
	return c.Render(200, r.JSON(data))
}
