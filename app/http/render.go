package http

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

// RenderOptions contains all the renderer config
var RenderOptions = render.Options{
	TemplatesBox: packr.New("app:templates", "../resources/views"),
	AssetsBox:    packr.New("app:assets", "../public"),
}

// R is a reference to the render engine
var R *render.Engine

func init() {
	R = render.New(RenderOptions)
}
