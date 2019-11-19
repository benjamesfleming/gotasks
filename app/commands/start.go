package commands

import (
	"fmt"
	"os"
	"os/signal"

	"git.benfleming.nz/benfleming/gotasks/app"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func or(ctx *cli.Context, a, b string) string {
	if ctx.IsSet(a) {
		return ctx.String(a)
	}
	return b
}

var _startFlags = []cli.Flag{
	// Server Config
	altsrc.NewStringFlag(&cli.StringFlag{Name: "server.port", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "server.bind", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "server.host", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "server.avatar-dir", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "server.unique-key", Hidden: true}),

	// Database Config
	altsrc.NewStringFlag(&cli.StringFlag{Name: "database.type", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "database.connection", Hidden: true}),

	// Github Config
	altsrc.NewBoolFlag(&cli.BoolFlag{Name: "auth.github.enabled", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "auth.github.client-id", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "auth.github.secret", Hidden: true}),

	// Google Config
	altsrc.NewBoolFlag(&cli.BoolFlag{Name: "auth.google.enabled", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "auth.google.client-id", Hidden: true}),
	altsrc.NewStringFlag(&cli.StringFlag{Name: "auth.google.secret", Hidden: true}),

	// Other Options
	&cli.StringFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Usage:    "Load configuration from `FILE`",
		Required: true,
	},
}

// StartCommand ...
var StartCommand = &cli.Command{
	Flags:  _startFlags,
	Name:   "start",
	Usage:  "start the server",
	Before: altsrc.InitInputSourceWithContext(_startFlags, altsrc.NewTomlSourceFromFlagFunc("config")),
	Action: func(ctx *cli.Context) error {
		opts := &app.ServerOptions{
			// Server Config
			HTTPPort:  or(ctx, "server.port", "3000"),
			BindAddr:  or(ctx, "server.bind", "0.0.0.0"),
			Host:      or(ctx, "server.host", "https://127.0.0.1:3000"),
			AvatarDir: or(ctx, "server.avatar-dir", "/tmp/gotasks/avatars"),
			UniqueKey: ctx.String("server.unique-key"),

			// Database Config
			DatabaseType:       or(ctx, "database.type", "sqlite3"),
			DatabaseConnection: or(ctx, "database.connection", "/tmp/gotasks/database.db"),

			// Github Config
			GithubEnabled:  ctx.Bool("auth.github.enabled"),
			GithubClientID: ctx.String("auth.github.client-id"),
			GithubSecret:   ctx.String("auth.github.secret"),

			// Google Config
			GoogleEnabled:  ctx.Bool("auth.google.enabled"),
			GoogleClientID: ctx.String("auth.google.client-id"),
			GoogleSecret:   ctx.String("auth.google.secret"),
		}

		s := app.NewServer(opts)
		c := make(chan os.Signal, 1)

		signal.Notify(c, os.Interrupt)
		go func() {
			<-c

			fmt.Println("----")
			s.Logger.Infof("Shutdown Signal Recived, Goodbye")

			s.Router.Close()
			s.DB.Close()
		}()

		return s.Start()
	},
}
