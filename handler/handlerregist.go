package handler

import (
	controller "ProjekGolangMVC/Controller"
	"ProjekGolangMVC/View"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type RegisterData struct {
	Success bool
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/register.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	if r.Method == http.MethodPost {
		// Process the registration form submission
		email := r.FormValue("email")
		password := r.FormValue("password")
		name := r.FormValue("name")
		username := r.FormValue("username")

		// Simple validation
		if email == "" || password == "" || name == "" || username == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Insert data into the database or perform any necessary action
		success := controller.ControlinsertDataRegister(email, password, name, username)
		if !success {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			log.Println("Failed to insert data for user registration")
			return
		}

		// Redirect to a success page or another endpoint
		data := RegisterData{Success: true}
		tmpl.Execute(w, data)
		View.ViewReadallDataRegister()

		return
	}

	err = tmpl.Execute(w, RegisterData{Success: false})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
	fmt.Println("Register template served successfully")
}
