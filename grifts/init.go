package grifts

import (
	"github.com/benjamesfleming/gotasks/app"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(app.NewApp())
}
