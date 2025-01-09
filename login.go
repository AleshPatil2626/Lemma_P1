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

// Struct to pass role data to templates
type PageData struct {
	Role string
}

// Handler for the login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request to show the login page
	if r.Method == http.MethodGet {
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
	} else if r.Method == http.MethodPost {
		// Handle POST request to process the login form submission
		email := r.FormValue("email")
		password := r.FormValue("password") // Capture the password from the form

		// Validate email and password inputs
		if email == "" {
			http.Error(w, "Email cannot be empty", http.StatusBadRequest)
			return
		}
		if password == "" {
			http.Error(w, "Password cannot be empty", http.StatusBadRequest)
			return
		}

		// Connect to the database
		connStr := os.Getenv("DB_CONN_STR")
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			http.Error(w, "Error opening database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Query the database to get the role for the given email
		var role string
		err = db.QueryRow("SELECT role FROM registerusers_tbl WHERE email = ?", email).Scan(&role)
		if err == sql.ErrNoRows {
			http.Error(w, "Email not found", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Error querying the database", http.StatusInternalServerError)
			return
		}

		// Define page data
		pageData := PageData{Role: role}

		// Conditional check for the role "user"
		if role == "user" {
			tmpl, err := template.ParseFiles("templates/userview.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, pageData) // Pass the role data to the template
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
				return
			}
			return
		}

		// Redirect based on other roles
		if role == "admin" {
			tmpl, err := template.ParseFiles("templates/adminview.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, pageData) // Pass the role data to the template
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
				return
			}
			return
		}
		if role == "Super Admin" {
			tmpl, err := template.ParseFiles("templates/superadminview.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, pageData) // Pass the role data to the template
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
				return
			}
			return
		}

		// If role is not found or unknown
		http.Error(w, "Unknown role", http.StatusForbidden)
	}
}

// Handler for the user view page (userview.html)
func userviewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/userview.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}
	pageData := PageData{Role: "user"} // Example role assignment
	err = tmpl.Execute(w, pageData)
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
	pageData := PageData{Role: "admin"} // Example role assignment
	err = tmpl.Execute(w, pageData)
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
	pageData := PageData{Role: "superadmin"} // Example role assignment
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}

// Define the handler for the product page --admin
func productHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the product page
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
//--users
func buyproductHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the product page
	tmpl, err := template.ParseFiles("templates/buyproduct.html")
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
