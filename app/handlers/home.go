package handlers

import (
	"bytes"
	"html/template"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

// HomeHandler is a default handler
// GET /
func HomeHandler(e echo.Context) error {
	t, _ := template.New("index").Parse(
		e.Get("TemplatesBox").(*rice.Box).MustString("index.html"),
	)

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, time.Now().Unix()); err != nil {
		e.Logger().Fatalf("Failed To Prase HTML Template, %s", err)
		return e.String(500, "Failed To Parse The HTML Template")
	}

	return e.HTML(200, tpl.String())
}
