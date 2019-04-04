package main

import (
	"html/template"
	"io"
	"login-golang/db"
	"login-golang/routers"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

// FUNC RENDEERED
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {

	e := routers.Init()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*/*.html")),
	}
	e.Renderer = renderer
	// e.Renderer = config.NewRenderer("views/*.html", true)
	db.Init()
	e.Start(":9000")
}
