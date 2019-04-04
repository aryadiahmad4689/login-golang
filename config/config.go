package config

import (
	"github.com/tkanos/gonfig"
)

// func GetConfig
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/env.json", &configuration)
	return configuration
}

// type M map[string]interface{}

// type Renderer struct {
// 	template *template.Template
// 	debug    bool
// 	location string
// }

// // Inisialisasi Renderer
// func NewRenderer(location string, debug bool) *Renderer {
// 	tpl := new(Renderer)
// 	tpl.location = location
// 	tpl.debug = debug

// 	tpl.ReloadTemplates()

// 	return tpl
// }

// // Melakukan Pharse Glob
// func (t *Renderer) ReloadTemplates() {
// 	t.template = template.Must(template.ParseGlob(t.location))
// }

// // Func Render
// func (t *Renderer) Render(
// 	w io.Writer,
// 	name string,
// 	data interface{},
// 	c echo.Context,
// ) error {
// 	if t.debug {
// 		t.ReloadTemplates()
// 	}

// 	return t.template.ExecuteTemplate(w, name, data)
// }
