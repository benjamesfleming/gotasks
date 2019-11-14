package policies

import (
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// TaskPolicy is the resource for the User model
type TaskPolicy struct {
	DB   *gorm.DB
	User *models.User
	Policy
}

// NewTaskPolicy returns a policy based on the given context
func NewTaskPolicy(e echo.Context) *TaskPolicy {
	db := e.Get("Database").(*gorm.DB)
	user := e.Get("User").(*models.User)

	return &TaskPolicy{
		DB:   db,
		User: user,
	}
}

// CanList validtes if the current user is allowed to
// list all the tasks.
func (p *TaskPolicy) CanList() bool {
	return p.User.IsAdmin
}

// CanShow validates if the current user is allowed to
// show the requested task
func (p *TaskPolicy) CanShow(t *models.Task) bool {
	return p.User.IsAdmin || (p.User.ID == t.UserID)
}

// CanCreate validates if the current user is allowed to
// create a new task
func (p *TaskPolicy) CanCreate() bool {
	return true
}

// CanUpdate validates if the current user is allowed to
// update the requested task
func (p *TaskPolicy) CanUpdate(t *models.Task) bool {
	return p.User.IsAdmin || (p.User.ID == t.UserID)
}
