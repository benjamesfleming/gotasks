package app

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"

	"github.com/benjamesfleming/gotasks/app/http"
	"github.com/benjamesfleming/gotasks/app/http/handlers"
	"github.com/benjamesfleming/gotasks/app/http/middleware"
	"github.com/benjamesfleming/gotasks/app/http/resources"
	"github.com/benjamesfleming/gotasks/app/models"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// NewApp is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func NewApp() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_gotasks_session",
		})

		handlers.Init(app)

		// Add the database to the context
		app.Use(popmw.Transaction(models.DB))

		// Add the application routes.
		// https://gobuffalo.io/en/docs/routing/
		api := app.Group("/api")
		api.Use(middleware.AuthMiddleware)
		api.Resource("/tasks", resources.TasksResource{})
		api.Resource("/users", resources.UsersResource{})

		// Auth routes
		auth := app.Group("/auth")
		auth.GET("/logout", handlers.AuthLogout)
		auth.GET("/3rd-party/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/3rd-party/{provider}/callback", handlers.AuthCallback)

		// Misc routes
		app.GET("/", handlers.HomeHandler)
		app.ServeFiles("/", http.RenderOptions.AssetsBox)

		// Error handling
		app.ErrorHandlers[404] = func(status int, err error, c buffalo.Context) error {
			return handlers.HomeHandler(c)
		}
	}

	return app
}
