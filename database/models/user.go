package models

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gobuffalo/nulls"
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	ProviderID nulls.String
	Avatar     string
	Username   string `gorm:"size:32"`
	Email      string `gorm:"type:varchar(255);unique_index"`
	FirstName  string `gorm:"size:64"`
	LastName   string `gorm:"size:64"`
	IsAdmin    bool   `orm:"default:false"`
	Tasks      []Task
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
