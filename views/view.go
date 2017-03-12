package views

import (
	"html/template"
	"path/filepath"
)

var LayoutDir string = "views/layouts"

func NewView(layout string, files ...string) *View {
	files = append(files, LayoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != err {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
func LayoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

type View struct {
	Template *template.Template
	Layout   string
}
