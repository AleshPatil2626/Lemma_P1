package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		fs.ServeHTTP(w, r)
	})))

	http.HandleFunc("/", loginHandler)            // Login handler for the root route
	http.HandleFunc("/login", loginHandler)       // Handle login form submission
	http.HandleFunc("/register", registerHandler) // Handle registration form submission
	http.HandleFunc("/welcome", welcomeHandler)   // Handle welcome page

	http.HandleFunc("/userview", userviewHandler)             // Handle user view
	http.HandleFunc("/adminview", adminviewHandler)           // Handle admin view
	http.HandleFunc("/superadminview", superadminviewHandler) // Handle super admin view

	http.HandleFunc("/users", displayUsers)  // Handle users list
	http.HandleFunc("/admin", displayAdmins) // Handle admins list

	http.HandleFunc("/product", productHandler) // Handle product page

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
