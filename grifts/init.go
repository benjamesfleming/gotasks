package grifts

import (
	"github.com/benjamesfleming/gotasks/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
