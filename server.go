package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	h "git.benfleming.nz/benfleming/gotasks/app/handlers"
	"git.benfleming.nz/benfleming/gotasks/database/models"
	r "git.benfleming.nz/benfleming/gotasks/routes"
	rice "github.com/GeertJohan/go.rice"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/nulls"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var templatePath = "resources/views"
var assetPath = "public"

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()
	envy.Load()

	// Load the database
	db, err := gorm.Open(envy.Get("DB_TYPE", ""), envy.Get("DB_CONNECTION", ""))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})

	// Load the rice boxes
	assetsBox := rice.MustFindBox("public").HTTPBox()

	// Create auth service with providers
	service := auth.NewService(
		auth.Opts{
			SecretReader: token.SecretFunc(func() (string, error) { // secret key for JWT
				return envy.Get("SESSION_SECRET", "xxx"), nil
			}),
			TokenDuration:  time.Minute * 5, // token expires in 5 minutes
			CookieDuration: time.Hour * 24,  // cookie expires in 1 day and will enforce re-login
			Issuer:         "gotasks",
			URL:            envy.Get("HOST", "http://127.0.0.1:3000"),
			AvatarStore:    avatar.NewLocalFS("/tmp"),
			ClaimsUpd: token.ClaimsUpdFunc(func(claims token.Claims) token.Claims {
				if claims.User != nil {
					claims.User.SetAdmin(false)
					claims.User.SetBoolAttr("registered", false)

					user := new(models.User)
					if db.Where(&models.User{ProviderID: nulls.NewString(claims.User.ID)}).First(&user); !user.IsEmpty() {
						claims.User.SetAdmin(user.IsAdmin)
						claims.User.SetBoolAttr("registered", true)
					}
				}
				return claims
			}),
		},
	)

	service.AddProvider("github", envy.Get("GITHUB_ID", ""), envy.Get("GITHUB_SECRET", ""))

	// Retrieve auth middleware
	m := service.Middleware()

	isAuth := echo.WrapMiddleware(m.Auth)
	// isAdmin := echo.WrapMiddleware(m.AdminOnly)
	// isTraced := echo.WrapMiddleware(m.Trace)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echo.WrapMiddleware(m.Trace))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := new(models.User)
			tkn, err := token.GetUserInfo(c.Request())

			fmt.Print("\n\n\n")
			fmt.Println("  Timestamp          : ", time.Now())
			fmt.Println("  Request Path       : ", c.Request().URL.Path)
			fmt.Println("  Authenticated User : ", tkn.Name, err)
			fmt.Print("\n")

			if err == nil {
				db.Where(&models.User{ProviderID: nulls.NewString(tkn.ID)}).First(&user)
			}

			c.Set("Database", db)
			c.Set("User", user)

			return next(c)
		}
	})

	// Routes
	e.GET("/", h.HomeHandler)
	e.GET("/assets/*", echo.WrapHandler(http.FileServer(assetsBox)))

	r.RegisterAPIRoutes(e, isAuth)
	r.RegisterAuthRoutes(e, service)

	// Add error handling
	echo.NotFoundHandler = h.HomeHandler

	// Implement the template renderer
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob(templatePath + "/*.html")),
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + envy.Get("PORT", "3000")))
}
