package app

import (
	h "git.benfleming.nz/benfleming/gotasks/app/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterAPIRoutes registers all the api routes
// GET,POST /api/*
func registerAPIRoutes(s *Server) {

	api := s.Router.Group(
		"/api", s.AuthMiddleware.IsAuth,
	)

	api.GET("/tasks", h.TaskListHandler)
	api.GET("/tasks/:id", h.TaskShowHandler)
	api.POST("/tasks", h.TaskCreateHandler)
	api.POST("/tasks/:id", h.TaskUpdateHandler)
	api.POST("/tasks/:task_id/steps/:step_id", h.TaskStepUpdateHandler)
	api.DELETE("/tasks/:id", h.TaskDeleteHandler)

	api.GET("/users", h.UserListHandler)
	api.GET("/users/:id", h.UserShowHandler)
	api.GET("/users/:id/tasks", h.UserShowTasksHandler)
	api.POST("/users", h.UserCreateHandler)

}

// RegisterAuthRoutes registers all the routes for auth handling
// GET,POST /auth/*
func registerAuthRoutes(s *Server) {

	authRoutes, avaRoutes := s.AuthService.Handlers()

	s.Router.GET("/auth/me", h.AuthMeHandler)
	s.Router.POST("/auth/register", h.AuthRegisterHandler)

	s.Router.Any("/auth/*", echo.WrapHandler(authRoutes))
	s.Router.Any("/avatar/*", echo.WrapHandler(avaRoutes))

}
