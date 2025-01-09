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

	http.HandleFunc("/register", registerHandler) // Serve login form
	http.HandleFunc("/login", loginHandler)       // Handle login form submission
	http.HandleFunc("/", indexHandler)            // Handle the root route (index page)
	http.HandleFunc("/welcome", welcomeHandler)

	http.HandleFunc("/userview", userviewHandler)
	http.HandleFunc("/adminview", adminviewHandler)
	http.HandleFunc("/superadminview", superadminviewHandler)

	http.HandleFunc("/users", displayUsers)
	http.HandleFunc("/admin", displayAdmins)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
