package controller

import (
	U "go-userlogin/until"
	"net/http"
)

// Page type as a structure
type Page struct {
	Title    string
	UserID   interface{}
	Username interface{}
}

func loggedIn(w http.ResponseWriter, url string, r *http.Request) {
	var URL string
	if url == "" {
		URL = "/login"
	} else {
		URL = url
	}
	id, _ := U.AllSessions(r)
	if id == nil {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

