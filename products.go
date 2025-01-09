package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Define the handler for the product page
func productHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and execute the product template (product.html)
	tmpl, err := template.ParseFiles("templates/product.html")
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
