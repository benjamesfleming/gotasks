package handlers

import (
	p "git.benfleming.nz/benfleming/gotasks/app/policies"
	"git.benfleming.nz/benfleming/gotasks/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// var (
// 	errUnauthorized   = errors.New("unauthorized, you cannot access this user")
// 	errNotImplemented = errors.New("not implemented")
// )

// UserListHandler handles the request to list all users
// GET /api/users
func UserListHandler(e echo.Context) error {
	if !p.NewUserPolicy(e).CanList() {
		return e.JSON(401, errUnauthorized)
	}

	users := new([]models.User)

	db := e.Get("Database").(*gorm.DB)
	db.Find(&users)

	return e.JSON(200, users)
}

// UserShowHandler handles the requests to show a single user
// GET /api/users/:id
func UserShowHandler(e echo.Context) error {
	id := e.Param("id")
	user := new(models.User)

	db := e.Get("Database").(*gorm.DB)
	db.Where("id = ?", id).First(&user)

	if !p.NewUserPolicy(e).CanShow(user) {
		return e.JSON(401, errUnauthorized)
	}

	return e.JSON(200, user)
}

// UserShowTasksHandler handles the requests to show a users tasks
// GET /api/users/:id/tasks
func UserShowTasksHandler(e echo.Context) error {
	id := e.Param("id")
	user := new(models.User)
	tasks := new([]models.Task)

	db := e.Get("Database").(*gorm.DB)
	db.Where("id = ?", id).First(&user).Related(&tasks)

	if !p.NewUserPolicy(e).CanShowTasks(user) {
		return e.JSON(401, errUnauthorized)
	}

	return e.JSON(200, tasks)
}

// UserCreateHandler handles the requests to create a new user
// POST /api/users
func UserCreateHandler(e echo.Context) error {
	return e.JSON(501, errNotImplemented)
}
