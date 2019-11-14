package models

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Step ...
type Step struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null;"`
	Title     string
	Completed nulls.Time
	Parent    *Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Step) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

// IsEmpty checks if the task struct has been successfully created
func (s Step) IsEmpty() bool {
	return s.ID == uuid.Nil
}

// Validate checks if the task struct is vaild, then creates them
func (s Step) Validate() error {
	return v.ValidateStruct(&s,
		v.Field(&s.TaskID, v.Required),
		v.Field(&s.Title, v.Required, v.Length(1, 64)),
	)
}
