package main

import (
	"fmt"
	"go-web-app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404: page not found</h1>")
}

// the name main is essential for this function and package to run
func main() {
	homeView = views.CreateView("views/home.gohtml", "views/layouts/footer.gohtml")
	contactView = views.CreateView("views/contact.gohtml", "views/layouts/footer.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	h := http.HandlerFunc(pageNotFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
