package handlers

import (
	"bytes"
	"html/template"
	"strconv"
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

	tpl := new(bytes.Buffer)
	data := map[string]string{
		"UniqueKey":   "?v=" + strconv.FormatInt(time.Now().Unix(), 10),
		"Environment": e.Get("Environment").(string),
	}

	if data["Environment"] != "development" {
		data["UniqueKey"] = ""
	}

	if err := t.Execute(tpl, data); err != nil {
		e.Logger().Fatalf("Failed To Prase HTML Template, %s", err)
		return e.String(500, "Failed To Parse The HTML Template")
	}

	return e.HTML(200, tpl.String())
}
