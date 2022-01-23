package views

import "html/template"

type View struct {
	Template *template.Template
}

func CreateView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}
