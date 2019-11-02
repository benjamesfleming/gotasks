package resources

import (
	"fmt"

	"github.com/benjamesfleming/gotasks/app/http"
	"github.com/benjamesfleming/gotasks/app/models"
	"github.com/benjamesfleming/gotasks/app/policies"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

var r = http.R

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

/*
██╗     ██╗███████╗████████╗
██║     ██║██╔════╝╚══██╔══╝
██║     ██║███████╗   ██║
██║     ██║╚════██║   ██║
███████╗██║███████║   ██║
╚══════╝╚═╝╚══════╝   ╚═╝

*/

// List gets all Users.
// GET /api/users
func (v UsersResource) List(c buffalo.Context) error {
	if policies.NewUsersPolicy(c).CanList() {
		// Grab the database connection from the current context
		// else return error and break
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return c.Render(500, r.String("internal server error, please contact the admin"))
		}

		// Get all the users from the database
		// paginate with `page=?`, `per_page=?` params
		users := &models.Users{}
		if tx.PaginateFromParams(c.Params()).All(users) != nil {
			return c.Render(500, r.String("internal server error, please contact the admin"))
		}

		return c.Render(200, r.JSON(users))
	}
	return c.Render(401, r.String("you are unauthorized to access this resource"))
}

/*
███████╗██╗  ██╗ ██████╗ ██╗    ██╗
██╔════╝██║  ██║██╔═══██╗██║    ██║
███████╗███████║██║   ██║██║ █╗ ██║
╚════██║██╔══██║██║   ██║██║███╗██║
███████║██║  ██║╚██████╔╝╚███╔███╔╝
╚══════╝╚═╝  ╚═╝ ╚═════╝  ╚══╝╚══╝

*/

// Show gets the data for one User.
// GET /api/users/{user_id}
func (v UsersResource) Show(c buffalo.Context) error {
	id := c.Param("user_id")
	if policies.NewUsersPolicy(c).CanShow(id) {
		// Grab the database connection from the current context
		// else return error and break
		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return c.Render(500, r.String("internal server error, please contact the admin"))
		}

		// Attempt to find the user with the given user_id
		// else return 404 not found
		user := &models.User{}
		if err := tx.Find(user, id); err != nil {
			return c.Error(404, err)
		}

		// Return the requested user
		return c.Render(200, r.JSON(user))
	}
	return c.Render(401, r.String("you are unauthorized to access this resource"))
}

/*
 ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗
██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝
██║     ██████╔╝█████╗  ███████║   ██║   █████╗
██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝
╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝

*/

// Create adds a User to the DB.
// POST /api/users
func (v UsersResource) Create(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Bind the user to the context
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	// Validate the created user
	// if there are any unknown errors then break
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	// Check for any validataion errors
	// if there are any return them in a 400 request
	if verrs.HasAny() {
		return c.Render(400, r.JSON(verrs))
	}

	// Return with the created user
	return c.Render(200, r.JSON(user))
}

/*
██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗
██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝
██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗
██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝
╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗
 ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝

*/

// Update changes a User in the DB
// PUT /api/users/{user_id}
func (v UsersResource) Update(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Attempt to find the user with the given user_id
	// else return 404 not found
	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind the user to the context
	if err := c.Bind(user); err != nil {
		return err
	}

	// Validate the updated user
	// if there are any unknown errors then break
	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return err
	}

	// Check for any validataion errors
	// if there are any return them in a 400 request
	if verrs.HasAny() {
		return c.Render(400, r.JSON(verrs))
	}

	// Return the updated user
	return c.Render(200, r.JSON(user))
}

/*
██████╗ ███████╗███████╗████████╗ ██████╗ ██████╗ ██╗   ██╗
██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║  ██║█████╗  ███████╗   ██║   ██║   ██║██████╔╝ ╚████╔╝
██║  ██║██╔══╝  ╚════██║   ██║   ██║   ██║██╔══██╗  ╚██╔╝
██████╔╝███████╗███████║   ██║   ╚██████╔╝██║  ██║   ██║
╚═════╝ ╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝

*/

// Destroy deletes a User from the DB.
// DELETE /api/users/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	// Grab the database connection from the current context
	// else return error and break
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Attempt to find the user with the given user_id
	// else return 404 not found
	user := &models.User{}
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	// Delete the given user from the database
	// else return error and break
	if err := tx.Destroy(user); err != nil {
		return err
	}

	// Return 200 success
	return c.Render(200, r.JSON(user.ID))
}
