package policies

import (
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// UserPolicy is the resource for the User model
type UserPolicy struct {
	DB   *gorm.DB
	User *models.User
	Policy
}

// NewUserPolicy returns a policy based on the given context
func NewUserPolicy(e echo.Context) *UserPolicy {
	db := e.Get("Database").(*gorm.DB)
	user := e.Get("User").(*models.User)

	return &UserPolicy{
		DB:   db,
		User: user,
	}
}

// CanList validtes if the current user is allowed to
// list all the user.
func (p *UserPolicy) CanList() bool {
	return p.User.IsAdmin
}

// CanShow validates if the current user is allowed to
// show the requested user
func (p *UserPolicy) CanShow(u *models.User) bool {
	return p.User.IsAdmin || (p.User.ID == u.ID)
}

// CanCreate validates if the current user is allowed to
// create a new user
func (p *UserPolicy) CanCreate() bool {
	return p.User.IsAdmin
}
