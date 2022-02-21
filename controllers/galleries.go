package controllers

import (
	"go-web-app/views"
	"net/http"
)

type Galleries struct {
	View *views.View
}

func NewGalleries() *Galleries {
	return &Galleries{
		View: views.New("bootstrap", "views/galleries/new.gohtml"),
	}
}

// GET /galleries/new
func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.View.Render(w, nil); err != nil {
		panic(err)
	}
}
