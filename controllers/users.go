package controllers

import (
	"fmt"
	"go-web-app/views"
	"net/http"

	"github.com/gorilla/schema"
)

type Users struct {
	View *views.View
}

type SignupForm struct {
	// adding struct tags to allow the `schema` package to get key names to decode from form data
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		View: views.New("bootstrap", "views/users/new.gohtml"),
	}
}

// New renders a form for user sign-up

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.View.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create takes info from signup form and creates a new User instance

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // must be called explicitly for requests with form data
		panic(err)
	}
	decoder := schema.NewDecoder()
	form := SignupForm{}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
