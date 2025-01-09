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

// SuperAdmin struct to map superadmin data
type SuperAdmin struct {
	ID       int
	Name     string
	Email    string
	Username string
	Mobile   int64
	Role     string
}

func init() {
	// Load .env file for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Handler to display all superadmin users
func displaySuperAdmins(w http.ResponseWriter, r *http.Request) {
	// Get database connection string from .env
	connStr := os.Getenv("DB_CONN_STR")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()

	// Query to fetch all users with role 'Super Admin'
	rows, err := db.Query("SELECT id, name, email, username, mobile, role FROM registerusers_tbl WHERE role = 'Super Admin'")
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var superAdmins []SuperAdmin
	for rows.Next() {
		var superAdmin SuperAdmin
		err := rows.Scan(&superAdmin.ID, &superAdmin.Name, &superAdmin.Email, &superAdmin.Username, &superAdmin.Mobile, &superAdmin.Role)
		if err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		superAdmins = append(superAdmins, superAdmin)
	}

	// Check for errors after looping through rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/superadmin.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}

	// Pass the superAdmins slice to the template
	err = tmpl.Execute(w, superAdmins)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}
