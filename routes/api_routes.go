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

}
