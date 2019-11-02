package policies

import (
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/gobuffalo/buffalo"
)

// UserPolicy is the resource for the User model
type UserPolicy struct {
	Context *buffalo.Context
	User    *models.User
	Policy
}

// NewUsersPolicy returns a policy based on the given context
func NewUsersPolicy(c buffalo.Context) *UserPolicy {
	user := c.Value("user").(*models.User)
	return &UserPolicy{
		Context: &c,
		User:    user,
	}
}

// CanList validtes if the current user is allowed to
// list all the users.
func (p *UserPolicy) CanList() bool {
	if v := p.User.NewValidator(); v != nil {
		v.HasPrivilege(`iam:gotasks:users:\*:list`)
		return v.Execute()
	}
	return false
}

// CanShow validates if the current user is allowed to
// access the requested user.
func (p *UserPolicy) CanShow(id string) bool {
	if p.User.ID.String() == id {
		return true
	}
	return false
}
