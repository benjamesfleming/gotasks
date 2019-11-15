package handlers

import (
	"errors"

	p "git.benfleming.nz/benfleming/gotasks/app/policies"
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var (
	errBadRequest     = errors.New("bad request, your request was invalid")
	errUnauthorized   = errors.New("unauthorized, you cannot access this task")
	errNotImplemented = errors.New("not implemented")
)

// TaskListHandler handles the request to list all tasks
// GET /api/tasks
func TaskListHandler(e echo.Context) error {
	if !p.NewTaskPolicy(e).CanList() {
		return e.JSON(401, errUnauthorized)
	}

	tasks := new([]models.Task)

	db := e.Get("Database").(*gorm.DB)
	db.Preload("Steps").Find(&tasks)

	return e.JSON(200, tasks)
}

// TaskShowHandler handles the requests to show a single task
// GET /api/tasks/:id
func TaskShowHandler(e echo.Context) error {
	id := e.Param("id")
	task := new(models.Task)

	db := e.Get("Database").(*gorm.DB)
	db.Preload("Steps").Where("id = ?", id).First(&task)

	if !p.NewTaskPolicy(e).CanShow(task) {
		return e.JSON(401, errUnauthorized)
	}

	return e.JSON(200, task)
}

// TaskCreateHandler handles the requests to create a new task
// POST /api/tasks
func TaskCreateHandler(e echo.Context) error {
	if !p.NewTaskPolicy(e).CanCreate() {
		return e.JSON(401, errUnauthorized)
	}

	task := new(models.Task)
	if err := e.Bind(task); err != nil {
		return e.JSON(400, err)
	}

	task.UserID = e.Get("User").(*models.User).ID
	if err := task.Validate(); err != nil {
		return e.JSON(400, err)
	}

	db := e.Get("Database").(*gorm.DB)
	if err := db.Create(&task).Error; err != nil {
		return e.JSON(400, err)
	}

	return e.JSON(200, task)
}

// TaskUpdateHandler handles the requests to update a new task
// POST /api/tasks/:id
func TaskUpdateHandler(e echo.Context) error {
	id := e.Param("id")
	task := new(models.Task)
	taskData := new(models.Task)

	user := e.Get("User").(*models.User)

	db := e.Get("Database").(*gorm.DB)
	db.Where("id = ?", id).First(&task)

	if !p.NewTaskPolicy(e).CanUpdate(task) {
		e.Logger().Errorf("[401 Unauthorized] User [%s] Failed Policy Check For Task [%s]", user.ID, task.ID)
		return e.JSON(401, errUnauthorized)
	}

	if err := e.Bind(taskData); err != nil {
		e.Logger().Error("[400 Bad Request] Failed To Bind Task Request To Task Object", err)
		return e.JSON(400, errBadRequest)
	}

	task.Title = taskData.Title
	task.Tags = taskData.Tags
	task.Note = taskData.Note
	task.Completed = taskData.Completed

	if err := task.Validate(); err != nil {
		e.Logger().Error("[400 Bad Request] Failed to validate the task", err)
		return e.JSON(400, err)
	}

	if err := db.Model(&task).Update(task).Error; err != nil {
		e.Logger().Error("[500 Internal Server Error] Failed to update the validated task in the database", err)
		return e.JSON(500, err)
	}

	return e.JSON(200, task)
}
