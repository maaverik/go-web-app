package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, dest interface{}) error { // interface - to say any type
	if err := r.ParseForm(); err != nil { // must be called explicitly for requests with form data
		return err
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(dest, r.PostForm); err != nil { // dest should be a pointer
		return err
	}
	return nil
}
