package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	dir       = "views/layouts/"
	extension = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

// execute the Template with extra data passed along to render the view to the writer
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// returns slice of all filenames of template files in layout directory
func getLayoutFiles() []string {
	files, err := filepath.Glob(dir + "*" + extension)
	if err != nil {
		panic(err)
	}
	return files
}

// parses all layout templates and extra templates passed along and returns a View object
func CreateView(layout string, files ...string) *View {
	files = append(files, getLayoutFiles()...)
	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
