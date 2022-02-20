package controllers

import "go-web-app/views"

// controller for resources that don't involve any dynamic interaction
type Static struct {
	HomeView    *views.View
	ContactView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:    views.New("bootstrap", "views/static/home.gohtml"),
		ContactView: views.New("bootstrap", "views/static/contact.gohtml"),
	}
}
