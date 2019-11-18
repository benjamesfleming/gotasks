package models

import (
	"errors"
	"time"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Task ...
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;"`
	Title       string
	Tags        string
	Note        string
	Position    nulls.String
	CompletedAt nulls.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Steps       []Step
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Task) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

// AfterDelete will delete all the orphaned steps
func (t *Task) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where(&Step{TaskID: t.ID}).Delete(&Step{})
	return
}

// IsEmpty checks if the task struct has been successfully created
func (t Task) IsEmpty() bool {
	return t.ID == uuid.Nil
}

// Validate checks if the task struct is vaild, then creates them
func (t Task) Validate() error {
	for _, step := range t.Steps {
		if err := step.Validate(); err != nil {
			return errors.New("steps are invalid")
		}
	}

	return v.ValidateStruct(&t,
		v.Field(&t.UserID, v.Required),
		v.Field(&t.Title, v.Required, v.Length(1, 64)),
	)
}
