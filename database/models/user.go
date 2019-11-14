package models

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	ProviderID nulls.String
	Avatar     string
	Username   string `gorm:"size:32"`
	Email      string `gorm:"type:varchar(255);unique_index"`
	FirstName  string `gorm:"size:64"`
	LastName   string `gorm:"size:64"`
	IsAdmin    bool   `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Tasks      []Task
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

// IsEmpty checks if the user struct has been successfully created
func (u User) IsEmpty() bool {
	return u.Username == ""
}

// Validate checks if the user struct is vaild, then creates them
func (u User) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.Username, v.Required, v.Length(3, 32)),
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.FirstName, v.Required, v.Length(1, 64)),
		v.Field(&u.LastName, v.Required, v.Length(1, 64)),
	)
}
