package models

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Task ...
type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;"`
	ParentID  uuid.UUID `gorm:"type:uuid;not null;"`
	Title     string
	Tags      string
	Note      string
	Completed nulls.Time
	Streak    uint
	Parent    *Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Task) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

// IsEmpty checks if the task struct has been successfully created
func (t Task) IsEmpty() bool {
	return t.ID == uuid.Nil
}

// Validate checks if the task struct is vaild, then creates them
func (t Task) Validate() error {
	return v.ValidateStruct(&t,
		v.Field(&t.UserID, v.Required),
		v.Field(&t.Title, v.Required, v.Length(1, 64)),
	)
}
