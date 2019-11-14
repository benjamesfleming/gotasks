package handlers

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

// HomeHandler is a default handler
// GET /
func HomeHandler(c echo.Context) error {
	indexHTML, _ := rice.MustFindBox("resources/views").String("index.html")
	return c.HTML(200, indexHTML)
}
