package controllers

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/views"
	"net/http"
)

type Users struct {
	View     *views.View
	uService *models.UserService
}

type SignupForm struct {
	// adding struct tags to allow the `schema` package to get key names to decode from form data
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(uService *models.UserService) *Users {
	return &Users{
		View:     views.New("bootstrap", "views/users/new.gohtml"),
		uService: uService,
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
	form := SignupForm{}
	if err := ParseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	err := u.uService.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, form.Name)
	fmt.Fprintln(w, form.Email)
	fmt.Fprintln(w, form.Password)
}
