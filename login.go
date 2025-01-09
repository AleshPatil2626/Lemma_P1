package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Login handler function
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if username is empty
		if username == "" {
			http.Error(w, "Username cannot be empty", http.StatusBadRequest)
			return
		}

		// Get database connection string from .env
		connStr := os.Getenv("DB_CONN_STR")
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Error opening database connection: ", err)
		}
		defer db.Close()

		// Query to check if the username exists
		var storedPassword string
		err = db.QueryRow("SELECT upassword FROM registerusers_tbl WHERE username = ?", username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Username does not exist", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			return
		}

		// Check if the password matches
		if storedPassword == password {
			// Redirect to the welcome page on successful login
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
		}
	} else {
		// Handle GET requests (serve the login page)
		tmpl, err := template.ParseFiles("templates/login.html")
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
}

// Welcome page handler
func welcomePage(w http.ResponseWriter, r *http.Request) {
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


