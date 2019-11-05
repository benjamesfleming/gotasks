package routes

import (
	h "git.benfleming.nz/benfleming/gotasks/app/handlers"
	"github.com/go-pkgz/auth"
	"github.com/labstack/echo/v4"
)

// RegisterAuthRoutes registers all the routes for auth handling
// GET,POST /auth/*
func RegisterAuthRoutes(e *echo.Echo, a *auth.Service) {

	authRoutes, avaRoutes := a.Handlers()

	e.GET("/auth/me", h.AuthMeHandler)
	e.POST("/auth/register", h.AuthRegisterHandler)

	e.Any("/auth/*", echo.WrapHandler(authRoutes))
	e.Any("/avatar/*", echo.WrapHandler(avaRoutes))

}
