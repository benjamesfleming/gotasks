package actions

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/benjamesfleming/gotasks/models"
	"github.com/benjamesfleming/gotasks/utils/cookies"

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
	tx.Where("email = ?", data.Email).First(user)

	if binary.BigEndian.Uint64(user.ID.Bytes()) == 0 {
		// Map the providers data to a User model
		user.UserName = data.Name
		user.Provider = data.Provider
		user.Email = data.Email
		user.Privileges = []string{""}

		// Validate the user data
		// if there are any unknown errors then break
		verrs, err := tx.ValidateAndCreate(user)
		if err != nil {
			return err
		}

		// Check for any validataion errors
		// if there are any return them in a 301 request to an error page
		if verrs.HasAny() {
			return c.Redirect(302, "/#/auth/error")
		}
	}

	cookies.Set(c, "user_id", user.ID.String())

	// Redirect the user to the clientside
	// to complete the signup process
	return c.Redirect(302, fmt.Sprintf("/#/auth/complete?user_id=%s", user.ID))
}

// AuthLogout ...
func AuthLogout(c buffalo.Context) error {
	c.Cookies().Delete("user_id")
	return c.Redirect(302, "/#/")
}
