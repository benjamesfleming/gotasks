package models

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Task ...
type Task struct {
	gorm.Model
	UserID    uint
	ParentID  uint
	Title     string
	Tags      string
	Note      string
	Completed time.Time
	Streak    uint
	Parent    *Task
}

// IsEmpty checks if the task struct has been successfully created
func (t Task) IsEmpty() bool {
	return t.ID == 0
}

// Validate checks if the task struct is vaild, then creates them
func (t Task) Validate() error {
	return v.ValidateStruct(&t,
		v.Field(&t.UserID, v.Required),
		v.Field(&t.Title, v.Required, v.Length(1, 64)),
	)
}
