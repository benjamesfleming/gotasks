package policies

import (
	"github.com/benjamesfleming/gotasks/app/models"
	"github.com/gobuffalo/buffalo"
)

// UsersPolicy is the resource for the User model
type UsersPolicy struct {
	Context *buffalo.Context
	User    *models.User
	Policy
}

// NewUsersPolicy returns a policy based on the given context
func NewUsersPolicy(c buffalo.Context) *UsersPolicy {
	user := c.Value("user").(*models.User)
	return &UsersPolicy{
		Context: &c,
		User:    user,
	}
}

// CanList validtes if the current user is allowed to
// list all the users.
func (p *UsersPolicy) CanList() bool {
	if v := p.User.NewValidator(); v != nil {
		v.HasPrivilege(`iam:gotasks:users:\*:list`)
		return v.Execute()
	}
	return false
}

// CanShow validates if the current user is allowed to
// access the requested user.
func (p *UsersPolicy) CanShow(id string) bool {
	if p.User.ID.String() == id {
		return true
	}
	return false
}
