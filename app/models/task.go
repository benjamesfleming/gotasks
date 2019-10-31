package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

/*
███╗   ███╗ ██████╗ ██████╗ ███████╗██╗
████╗ ████║██╔═══██╗██╔══██╗██╔════╝██║
██╔████╔██║██║   ██║██║  ██║█████╗  ██║
██║╚██╔╝██║██║   ██║██║  ██║██╔══╝  ██║
██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗███████╗
╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝

*/

// Task ...
// The Task model
type Task struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Tasks ...
// List of Task models
type Tasks []Task

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
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// String JSON encode Users as string
// Encodes into a JSON array containing a list of User models
func (t Tasks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
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
// This method is not required and may be deleted.
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
	), err
}
