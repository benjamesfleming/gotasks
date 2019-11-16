package routes

import (
	h "git.benfleming.nz/benfleming/gotasks/app/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterAPIRoutes registers all the api routes
// GET,POST /api/*
func RegisterAPIRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {

	api := e.Group("/api", m...)

	api.GET("/tasks", h.TaskListHandler)
	api.GET("/tasks/:id", h.TaskShowHandler)
	api.POST("/tasks", h.TaskCreateHandler)
	api.POST("/tasks/:id", h.TaskUpdateHandler)
	api.POST("/tasks/:task_id/steps/:step_id", h.TaskStepUpdateHandler)

	api.GET("/users", h.UserListHandler)
	api.GET("/users/:id", h.UserShowHandler)
	api.GET("/users/:id/tasks", h.UserShowTasksHandler)
	api.POST("/users", h.UserCreateHandler)

}
