package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

// RenderOptions contains all the renderer config
var RenderOptions = render.Options{
	TemplatesBox: packr.New("app:templates", "../templates"),
	AssetsBox:    packr.New("app:assets", "../public"),
}
var r *render.Engine

func init() {
	r = render.New(RenderOptions)
}
