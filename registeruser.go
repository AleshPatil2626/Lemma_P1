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
	"golang.org/x/crypto/bcrypt"
)

func init() {
	// Load .env file for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get form values
		uname := r.FormValue("uname")
		email := r.FormValue("email")
		mobileno := r.FormValue("mobileno")
		username := r.FormValue("username")
		upassword := r.FormValue("upassword")
		urole := r.FormValue("urole")

		// Simple validations: Check if all fields are filled
		if uname == "" || email == "" || mobileno == "" || username == "" || upassword == "" || urole == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Hash the password before storing it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(upassword), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
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
		err = db.QueryRow("SELECT username FROM registerusers_tbl WHERE username = ?", username).Scan(&existingUsername)
		if err == nil {
			http.Error(w, "Username already taken", http.StatusBadRequest)
			return
		}

		// Insert the user data into the registerusers_tbl table
		_, err = db.Exec("INSERT INTO registerusers_tbl (name, email, username, password, mobile, role) VALUES (?, ?, ?, ?, ?, ?)", uname, email, username, hashedPassword, mobileno, urole)
		if err != nil {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		// Registration success response (can be rendered using a template or a success message)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// Handle GET requests (serve the registration page)
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
