package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task ...
type Task struct {
	gorm.Model
	UserID    uint
	ParentID  uint
	title     string
	tags      string
	note      string
	completed time.Time
	streak    int
	Parent    *Task
}
