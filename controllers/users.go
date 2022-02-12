package controllers

import (
	"go-web-app/views"
	"net/http"
)

type Users struct {
	View *views.View
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
