package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

/*
███╗   ███╗ ██████╗ ██████╗ ███████╗██╗
████╗ ████║██╔═══██╗██╔══██╗██╔════╝██║
██╔████╔██║██║   ██║██║  ██║█████╗  ██║
██║╚██╔╝██║██║   ██║██║  ██║██╔══╝  ██║
██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗███████╗
╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝

*/

// User ...
// The User model
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	UserName     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password"`
	Password     string    `json:"-" db:"-"`
	Provider     string    `json:"provider" db:"provider"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Users ...
// List of User models
type Users []User

/*
███╗   ███╗ ██████╗ ██████╗ ███████╗██╗         ██╗  ██╗ ██████╗  ██████╗ ██╗  ██╗███████╗
████╗ ████║██╔═══██╗██╔══██╗██╔════╝██║         ██║  ██║██╔═══██╗██╔═══██╗██║ ██╔╝██╔════╝
██╔████╔██║██║   ██║██║  ██║█████╗  ██║         ███████║██║   ██║██║   ██║█████╔╝ ███████╗
██║╚██╔╝██║██║   ██║██║  ██║██╔══╝  ██║         ██╔══██║██║   ██║██║   ██║██╔═██╗ ╚════██║
██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗███████╗    ██║  ██║╚██████╔╝╚██████╔╝██║  ██╗███████║
╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝    ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝

*/

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), err
	}
	u.Email = strings.ToLower(u.Email)
	u.PasswordHash = string(ph)
	return tx.ValidateAndCreate(u)
}

/*
     ██╗███████╗ ██████╗ ███╗   ██╗    ███████╗███╗   ██╗ ██████╗ ██████╗ ██████╗ ███████╗
     ██║██╔════╝██╔═══██╗████╗  ██║    ██╔════╝████╗  ██║██╔════╝██╔═══██╗██╔══██╗██╔════╝
     ██║███████╗██║   ██║██╔██╗ ██║    █████╗  ██╔██╗ ██║██║     ██║   ██║██║  ██║█████╗
██   ██║╚════██║██║   ██║██║╚██╗██║    ██╔══╝  ██║╚██╗██║██║     ██║   ██║██║  ██║██╔══╝
╚█████╔╝███████║╚██████╔╝██║ ╚████║    ███████╗██║ ╚████║╚██████╗╚██████╔╝██████╔╝███████╗
 ╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝    ╚══════╝╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝

*/

// String JSON encode User as string
// Encodes into a JSON object container a single User model
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// String JSON encode Users as string
// Encodes into a JSON array containing a list of User models
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

/*
██╗   ██╗ █████╗ ██╗     ██╗██████╗  █████╗ ████████╗███████╗
██║   ██║██╔══██╗██║     ██║██╔══██╗██╔══██╗╚══██╔══╝██╔════╝
██║   ██║███████║██║     ██║██║  ██║███████║   ██║   █████╗
╚██╗ ██╔╝██╔══██║██║     ██║██║  ██║██╔══██║   ██║   ██╔══╝
 ╚████╔╝ ██║  ██║███████╗██║██████╔╝██║  ██║   ██║   ███████╗
  ╚═══╝  ╚═╝  ╚═╝╚══════╝╚═╝╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝

*/

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method
// Check the user email is valid and not taken
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.UserName, Name: "UserName"},
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "%s is already taken",
			Fn: func() bool {
				r, err := tx.Where("email = ?", u.Email).Exists(u)
				if err != nil {
					return false
				}
				return !r
			},
		},
	), err
}
