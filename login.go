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

// Handler for the login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request to show the login page
	if r.Method == http.MethodGet {
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
	} else if r.Method == http.MethodPost {
		// Handle POST request to process the login form submission
		username := r.FormValue("username")

		// Connect to the database
		connStr := os.Getenv("DB_CONN_STR")
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Error opening database connection: ", err)
			http.Error(w, "Error opening database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Query the database to get the role for the given username
		var role string
		err = db.QueryRow("SELECT role FROM registeruser_tbl WHERE username = ?", username).Scan(&role)

		// Check for errors in querying or if the username doesn't exist
		if err == sql.ErrNoRows {
			http.Error(w, "Username not found", http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}

		// Check if username is empty and redirect based on role
		if username == "" {
			switch role {
			case "user":
				http.Redirect(w, r, "/userview", http.StatusSeeOther)
			case "admin":
				http.Redirect(w, r, "/adminview", http.StatusSeeOther)
			case "uperadmin":
				http.Redirect(w, r, "/superadminview", http.StatusSeeOther)
			default:
				http.Error(w, "Unknown role", http.StatusForbidden)
			}
			return
		} else {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}
	}
}

// Handler for the user view page (userview.html)
func userviewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/userview.html")
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

// Handler for the admin view page (adminview.html)
func adminviewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/adminview.html")
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

// Handler for the superadmin view page (superadminview.html)
func superadminviewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/superadminview.html")
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


