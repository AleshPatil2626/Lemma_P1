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

// Handle user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get form values
		uname := r.FormValue("uname")
		email := r.FormValue("email")
		mobileno := r.FormValue("mobileno")
		username := r.FormValue("username")
		upassword := r.FormValue("upassword")
		urole := r.FormValue("urole")

		// Simple validations
		if uname == "" || email == "" || mobileno == "" || username == "" || upassword == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Get database connection string from .env
		connStr := os.Getenv("DB_CONN_STR")
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Error opening database connection: ", err)
		}
		defer db.Close()

		// Check if username already exists
		var existingUsername string
		err = db.QueryRow("SELECT username FROM registeruser WHERE username = ?", username).Scan(&existingUsername)
		if err == nil {
			http.Error(w, "Username already taken", http.StatusBadRequest)
			return
		}

		// Insert the user data into the registeruser table
		_, err = db.Exec("INSERT INTO registeruser (uname, email, mobileno, username, upassword, urole) VALUES (?, ?, ?, ?, ?, ?)", uname, email, mobileno, username, upassword, urole)
		if err != nil {
			log.Fatal("Error inserting data: ", err)
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		// Render the registration success page
		tmpl, err := template.ParseFiles("templates/registeruser.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
			return
		}

		// Send a success message to the template
		err = tmpl.Execute(w, "Registration successful! Please log in.")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
			return
		}
	} else {
		// Serve the registration page (GET request)
		tmpl, err := template.ParseFiles("templates/register.html")
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
