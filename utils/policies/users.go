package policies

import (
	"github.com/benjamesfleming/gotasks/utils/jwt"
	"github.com/gobuffalo/buffalo"
)

// UsersPolicy is the resource for the User model
type UsersPolicy struct {
	Context *buffalo.Context
	Claims  *jwt.Claims
	Policy
}

// NewUsersPolicy returns a policy based on the given context
func NewUsersPolicy(c buffalo.Context) *UsersPolicy {
	claims, _ := jwt.ClaimsFromHeader(c)
	return &UsersPolicy{
		Context: &c,
		Claims:  claims,
	}
}

// CanList validtes if the current user is allowed to list all the users
func (p *UsersPolicy) CanList() bool {
	if v := p.Claims.NewValidator(); v != nil {
		v.HasPrivilege(`iam:gotasks:users:\*:list`)
		return v.Execute()
	}
	return false
}
