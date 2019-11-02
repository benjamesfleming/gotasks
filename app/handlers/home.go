package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// HomeHandler is a default handler
// GET /
func HomeHandler(c echo.Context) error {
	return c.Render(200, "index.html", time.Now().Unix())
}
