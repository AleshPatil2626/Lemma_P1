package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Handler for the homepage (index page)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index page (GET request)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}

// Handler for the welcome page
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the welcome page (GET request)
	tmpl, err := template.ParseFiles("templates/welcome.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}
