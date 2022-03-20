package main

import (
	"fmt"
	"go-web-app/controllers"
	"go-web-app/models"
	"net/http"

	"github.com/gorilla/mux"
)

// DB connection info
const (
	host   = "localhost"
	port   = 5432
	user   = "nithin"
	dbname = "unsploosh"
)

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
	// connect to DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	uService, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer uService.Close()
	// uService.ResetDB()	// uncomment to reset the DB
	uService.AutoMigrate()

	usersController := controllers.NewUsers(uService)
	staticController := controllers.NewStatic()
	galleriesController := controllers.NewGalleries()

	r := mux.NewRouter()
	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.Handle("/faq", staticController.FAQView).Methods("GET")

	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")

	r.Handle("/login", usersController.LoginView).Methods("GET")
	r.HandleFunc("/login", usersController.Login).Methods("POST")

	r.HandleFunc("/galleries/new", galleriesController.New).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)

	http.ListenAndServe(":3000", r)
}
