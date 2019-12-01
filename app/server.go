package app

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"time"

	h "git.benfleming.nz/benfleming/gotasks/app/handlers"
	"git.benfleming.nz/benfleming/gotasks/app/models"
	rice "github.com/GeertJohan/go.rice"
	authservice "github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"github.com/gobuffalo/nulls"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // include mysql support
	_ "github.com/jinzhu/gorm/dialects/sqlite" // include sqlite3 support
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v2"
	"golang.org/x/oauth2"
)

// ServerOptions ...
type ServerOptions struct {

	// Server Config
	HTTPPort        string
	BindAddr        string
	Host            string
	AvatarDir       string
	UniqueKey       string
	CookiesHashKey  string
	CookiesBlockKey string

	// Database Config
	DatabaseType       string
	DatabaseConnection string

	// Auth Config
	GravatarEnabled string

	// Github OAuth2 Config
	GithubEnabled  bool
	GithubClientID string
	GithubSecret   string

	// Google OAuth2 Config
	GoogleEnabled  bool
	GoogleClientID string
	GoogleSecret   string
}

// Server ...
type Server struct {
	Config         *ServerOptions
	Router         *echo.Echo
	DB             *gorm.DB
	AssetsBox      *rice.Box
	TemplateBox    *rice.Box
	Logger         *lecho.Logger
	AuthService    *authservice.Service
	AuthMiddleware struct {
		Trace   echo.MiddlewareFunc
		IsAuth  echo.MiddlewareFunc
		IsAdmin echo.MiddlewareFunc
	}
}

// NewServer ...
func NewServer(opts *ServerOptions) *Server {
	s := &Server{Config: opts, Router: echo.New()}

	// Build The Logger
	// this is a zerolog wrapper for echo
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	s.Logger = lecho.New(
		zerolog.ConsoleWriter{Out: os.Stdout},
		lecho.WithTimestamp(),
	)
	s.Router.Logger = s.Logger

	// Get The Rice Boxes
	// load the rices boxes either from memory or embeded file system
	s.AssetsBox = rice.MustFindBox("../public")
	s.TemplateBox = rice.MustFindBox("../resources/views")

	// Load The Authentication Service
	// also load the enabled auth providers
	s.AuthService = authservice.NewService(
		authservice.Opts{

			// Return the secret defined in the config
			// should be a `64` bytes string
			SecretReader: token.SecretFunc(
				func() (string, error) { return s.Config.UniqueKey, nil },
			),

			// // Enable / Disable secure cookies
			// SecureCookies: true,

			// Set token & cookie timeouts
			// 	resonable defaults are 5 mins for the token
			// 	and 1 day for the cookie
			TokenDuration:  time.Minute * 5,
			CookieDuration: time.Hour * 24, // cookie expires in 1 day and will enforce re-login

			// JWT Issuer Name
			// identifies this program in the JWTs it creates
			Issuer: "gotasks",

			// Authentication URL
			// 	the URL which this server can be reached, and
			//	must be set for thrid party auth to work
			URL: s.Config.Host,

			// Avatar Store
			// local filesystem location that users avatars can be saved
			AvatarStore: avatar.NewLocalFS(s.Config.AvatarDir),

			// Claims Updater
			//	updates the claims sent when the `/auth/user` path is requested, this should
			//	add the registered status to the claims
			ClaimsUpd: token.ClaimsUpdFunc(func(claims token.Claims) token.Claims {
				if claims.User != nil {
					claims.User.SetAdmin(false)
					claims.User.SetBoolAttr("registered", false)

					user := new(models.User)
					if s.DB.Where(&models.User{ProviderID: nulls.NewString(claims.User.ID)}).First(&user); !user.IsEmpty() {
						claims.User.SetAdmin(user.IsAdmin)
						claims.User.SetBoolAttr("registered", true)
					}

					if s.Config.GravatarEnabled == "all" || (s.Config.GravatarEnabled == "fallback" && claims.User.Picture == "") {
						hash := md5.Sum([]byte(claims.User.Email))
						claims.User.Picture = fmt.Sprintf("https://www.gravatar.com/avatar/%x.png", hash)
					}
				}
				return claims
			}),
		},
	)

	if !s.Config.GithubEnabled && !s.Config.GoogleEnabled {
		s.Logger.Error("Authentication Config Invalid, Ensure Config Contains At Least ONE Enabled Provider")
		s.Logger.Fatal("Aborting")
		panic("no auth providers")
	}

	if s.Config.GithubEnabled {
		if s.Config.GithubClientID == "" || s.Config.GithubSecret == "" {
			s.Logger.Error("Github Provider Config Invalid, Ensure Config Contains [auth.github.client-id] AND [auth.github.secret]")
			s.Logger.Fatal("Aborting")
			panic("failed to enable github provider")
		}
		s.AuthService.AddCustomProvider(
			"github",

			// Github Client Details
			authservice.Client{
				Cid:     s.Config.GithubClientID,
				Csecret: s.Config.GithubSecret,
			},

			// Github Custom Handler Config
			provider.CustomHandlerOpt{
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://github.com/login/oauth/authorize",
					TokenURL: "https://github.com/login/oauth/access_token",
				},
				InfoURL: "https://api.github.com/user",
				MapUserFn: func(data provider.UserData, _ []byte) token.User {
					userInfo := token.User{
						ID:      "github_" + token.HashID(sha1.New(), data.Value("name")),
						Name:    data.Value("name"),
						Picture: data.Value("avatar_url"),
						Email:   data.Value("email"),
					}
					return userInfo
				},
				Scopes: []string{"user:email"},
			},
		)
	}

	if s.Config.GoogleEnabled {
		if s.Config.GoogleClientID == "" || s.Config.GoogleSecret == "" {
			s.Logger.Error("Google Provider Config Invalid, Ensure Config Contains [auth.google.client-id] AND [auth.google.secret]")
			s.Logger.Fatal("Aborting")
			panic("failed to enable google provider")
		}
		s.AuthService.AddProvider("google", s.Config.GoogleClientID, s.Config.GoogleSecret)
	}

	// Generate Auth Middleware
	// wrap all auth middleware for use by echo
	m := s.AuthService.Middleware()

	s.AuthMiddleware.Trace = echo.WrapMiddleware(m.Trace)
	s.AuthMiddleware.IsAuth = echo.WrapMiddleware(m.Auth)
	s.AuthMiddleware.IsAdmin = echo.WrapMiddleware(m.AdminOnly)

	// Apply Middleware
	// middleware runs in the order it is decleared
	s.Router.Use(middleware.Recover())
	s.Router.Use(s.AuthMiddleware.Trace)
	s.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := new(models.User)
			tkn, err := token.GetUserInfo(c.Request())

			fmt.Println("----")
			c.Logger().Infof("Request Path = %s", c.Request().URL.Path)
			c.Logger().Infof("Auth User    = %s", tkn.Name)

			if err == nil {
				s.DB.Where(&models.User{ProviderID: nulls.NewString(tkn.ID)}).First(&user)
			}

			c.Set("Database", s.DB)
			c.Set("User", user)
			c.Set("TemplatesBox", s.TemplateBox)

			return next(c)
		}
	})

	// Register Routes
	// register all the http routes in the application
	s.Router.GET("/", h.HomeHandler)

	assetsBox := http.FileServer(s.AssetsBox.HTTPBox())
	s.Router.GET("/assets/*", echo.WrapHandler(assetsBox))

	registerAPIRoutes(s)
	registerAuthRoutes(s)

	// Redirect Any Bad Traffic
	// return unknown traffic home
	echo.NotFoundHandler = h.HomeHandler
	echo.MethodNotAllowedHandler = h.HomeHandler

	return s
}

// Start loads the database then starts the http web server
// on the configured port
func (s *Server) Start() error {

	// Load The Database
	// if the database fails to load then panic
	db, err := gorm.Open(s.Config.DatabaseType, s.Config.DatabaseConnection)
	if err != nil {
		s.Logger.Error("Failed To Connect To Database")
		s.Logger.Fatal("Aborting")
		panic(err)
	}
	defer db.Close()

	s.DB = db

	s.DB.AutoMigrate(&models.User{})
	s.DB.AutoMigrate(&models.Task{})
	s.DB.AutoMigrate(&models.Step{})

	// Start The Web Server
	s.Logger.Infof("Binding To [:%s]", s.Config.HTTPPort)
	s.Logger.Infof("Ready")
	err = s.Router.Start(":" + s.Config.HTTPPort)
	if err != nil {
		s.Logger.Fatal(err)
	}
	s.Logger.Info("Goodbye")
	return err
}
