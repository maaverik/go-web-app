package controllers

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/views"
	"net/http"
)

type Users struct {
	NewView   *views.View
	LoginView *views.View
	uService  *models.UserService
}

type SignupForm struct {
	// adding struct tags to allow the `schema` package to get key names to decode from form data
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(uService *models.UserService) *Users {
	return &Users{
		NewView:   views.New("bootstrap", "views/users/new.gohtml"),
		LoginView: views.New("bootstrap", "views/users/login.gohtml"),
		uService:  uService,
	}
}

// New renders a form for user sign-up

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
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
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	err := u.uService.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, form.Name)
}

// Login takes info from login form and creates a session for a user

// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := ParseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.uService.Authenticate(form.Email, form.Password)

	if err == nil {
		cookie := http.Cookie{
			Name:  "email",
			Value: user.Email,
		}
		// cookie must be set in header before anything is written to ResponseWriter
		http.SetCookie(w, &cookie)
		fmt.Fprintln(w, user)
		u.CookieTest(w, r)
		return
	}

	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password")
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email: ", cookie.Value)
}
