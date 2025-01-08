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

// User struct to map user data
type User struct {
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

// Handler to display all users with role 'user'
func displayUsers(w http.ResponseWriter, r *http.Request) {
	// Get database connection string from .env
	connStr := os.Getenv("DB_CONN_STR")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	defer db.Close()

	// Query to fetch all users with role 'user'
	rows, err := db.Query("SELECT id, name, email, username, mobile, role FROM registerusers_tbl WHERE role = 'user'")
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Mobile, &user.Role)
		if err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Check for errors after looping through rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/user.html") // Make sure the file path is correct
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}

	// Pass the users slice to the template
	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %s", err), http.StatusInternalServerError)
		return
	}
}
