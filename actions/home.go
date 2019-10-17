package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler
// GET /
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
