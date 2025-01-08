package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register the handler for the routes
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/", loginForm)         // Serve login form
	http.HandleFunc("/login", loginHandler) // Handle login form submission
	http.HandleFunc("/welcome", welcomePage)

	http.HandleFunc("/users", displayUsers)
	http.HandleFunc("/admin", displayAdmins)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
