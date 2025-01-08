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

	// Register the handler for the routes
	http.HandleFunc("/register", registerHandler) // Serve login form
	http.HandleFunc("/login", loginHandler)       // Handle login form submission
	http.HandleFunc("/welcome", welcomePage)

	http.HandleFunc("/users", displayUsers)
	http.HandleFunc("/admin", displayAdmins)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
