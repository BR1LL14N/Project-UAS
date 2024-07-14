package main

import (
	"ProjekGolangMVC/handler" // Import package handler
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/login", handler.HandleLogin)
	http.HandleFunc("/register", handler.HandleRegister)
	http.HandleFunc("/firsttime", handler.HandleFirstTimeLogin)
	http.HandleFunc("/hasildata", handler.HandleDashboard)
	http.HandleFunc("/logout", handler.HandleLogout)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	// Start the server and set the default route to login page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}