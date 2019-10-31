package handlers

import (
	"github.com/benjamesfleming/gotasks/app/http"
	"github.com/gobuffalo/buffalo"
)

var r = http.R

// HomeHandler is a default handler
// GET /
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
