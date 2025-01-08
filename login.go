package main

import (
	"database/sql"
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

// Serve the login form
func loginForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("login").ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Handle login logic
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Get database connection string from .env
		connStr := os.Getenv("DB_CONN_STR")
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Error opening database connection: ", err)
		}
		defer db.Close()

		// Query to check if user exists with matching username and password
		var storedPassword string
		err = db.QueryRow("SELECT upassword FROM registeruser WHERE username = ?", username).Scan(&storedPassword)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Check if passwords match
		if storedPassword == password {
			// Redirect to the welcome page on successful login
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		// Serve the login page
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

// Welcome page after successful login
func welcomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("welcome").ParseFiles("templates/welcome.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
