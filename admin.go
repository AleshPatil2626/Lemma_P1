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

// Admin struct to map admin data
type Admin struct {
	ID       int
	Name     string
	Email    string
	Username string
	Mobile   string
	Role     string
}

func init() {
	// Load .env file for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Handler to display all users with role 'admin'
func displayAdmins(w http.ResponseWriter, r *http.Request) {
	// Get database connection string from .env
	connStr := os.Getenv("DB_CONN_STR")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()

	// Query to fetch all users with role 'admin'
	rows, err := db.Query("SELECT id, name, email, username, mobile, role FROM registerusers_tbl WHERE role = 'admin'")
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []Admin
	for rows.Next() {
		var admin Admin
		err := rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Username, &admin.Mobile, &admin.Role)
		if err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		admins = append(admins, admin)
	}

	// Check for errors after looping through rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/admin.html") // Make sure the file path is correct
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}

	// Pass the admins slice to the template
	err = tmpl.Execute(w, admins)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}


