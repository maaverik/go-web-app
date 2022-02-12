package main

import (
	"fmt"
	"go-web-app/controllers"
	"go-web-app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404: page not found</h1>")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// the name main is essential for this function and package to run
func main() {
	homeView = views.New("bootstrap", "views/home.gohtml")
	contactView = views.New("bootstrap", "views/contact.gohtml")
	users := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", users.New)
	h := http.HandlerFunc(pageNotFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
