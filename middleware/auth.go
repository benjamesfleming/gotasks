package middleware

import (
	"fmt"

	"github.com/benjamesfleming/gotasks/models"
	"github.com/benjamesfleming/gotasks/utils/cookies"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// AuthMiddleware ...
func AuthMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return fmt.Errorf("internal server error, please contact the admin")
		}

		user := &models.User{}
		if value, err := cookies.Get(c, "user_id"); err == nil {
			tx.Find(user, value)
		}
		c.Set("user", user)

		return next(c)
	}
}
