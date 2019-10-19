package jwt

import (
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/pop/slices"
	"github.com/gofrs/uuid"
)

// Claims ...
// The claims in a given jwt token
type Claims struct {
	ID    uuid.UUID     `json:"id"`
	Roles slices.String `json:"roles"`
	jwt.StandardClaims
}

// HasRole checks if the current claims contain a role
func (c *Claims) HasRole(rexp string) bool {
	pattern := regexp.MustCompile(rexp)
	for _, role := range c.Roles {
		if ok := pattern.MatchString(role); ok {
			return true
		}
	}
	return false
}
