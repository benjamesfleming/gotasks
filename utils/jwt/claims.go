package jwt

import (
	"fmt"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims ...
// The claims in a given jwt token
type Claims struct {
	Id         string   `json:"id"`
	Privileges []string `json:"privileges"`
	jwt.StandardClaims
}

// ClaimsValidator ...
// A validator used to check a policy
type ClaimsValidator struct {
	Claims   *Claims
	Required []string
}

// NewValidator contructs a new claims validator
func (c *Claims) NewValidator() *ClaimsValidator {
	return &ClaimsValidator{
		Required: []string{},
		Claims:   c,
	}
}

// HasPrivilege checks if the current claims contain a privilege
func (c *ClaimsValidator) HasPrivilege(rexp string) {
	c.Required = append(c.Required, rexp)
}

// Execute runs the validator and checks for all required privileges
func (c *ClaimsValidator) Execute() bool {
RequiredLoop:
	for _, rexp := range c.Required {
		pattern := regexp.MustCompile(rexp)
		for _, privilege := range c.Claims.Privileges {
			fmt.Println("Checking Required Claim [" + rexp + "] Equals Current Privilege [" + privilege + "]")
			if ok := pattern.MatchString(privilege); ok {
				continue RequiredLoop
			}
		}
		return false
	}
	return true
}
