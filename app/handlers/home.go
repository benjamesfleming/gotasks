package handlers

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

// HomeHandler is a default handler
// GET /
func HomeHandler(e echo.Context) error {
	indexHTML, _ := e.Get("TemplatesBox").(*rice.Box).String("index.html")
	return e.HTML(200, indexHTML)
}
