package handler

import (
	controller "ProjekGolangMVC/Controller"
	"fmt"
	"html/template"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("template/login.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    if r.Method == "POST" {
        // Process the login form submission
        email := r.FormValue("email")
        password := r.FormValue("password")

        // Check login credentials
        user := controller.ControlDataLogin(email, password)
        if user == nil {
            tmpl.Execute(w, map[string]string{"Error": "Username atau password salah"})
            return
        }

        // Set email in cookie
        cookie := &http.Cookie{
            Name:  "email",
            Value: email,
            Path:  "/",
            HttpOnly: true,
        }
        http.SetCookie(w, cookie)
        fmt.Println("Cookie set: ", cookie)

        // Check if the user has already completed their profile
        mahasiswa := controller.ControlGetMahasiswaByEmail(email)
        if mahasiswa == nil {
            // Redirect to first time login page
            http.Redirect(w, r, "/firsttime", http.StatusSeeOther)
        } else {
            // Redirect to dashboard
            http.Redirect(w, r, "/hasildata", http.StatusSeeOther)
        }
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    fmt.Println("Login template served successfully")
}



func HandleFirstTimeLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/firsttime.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	emailCookie, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	email := emailCookie.Value

	if r.Method == "POST" {
		tanggalLahir := r.FormValue("tanggalLahir")
		jurusan := r.FormValue("jurusan")
		transfer := r.FormValue("transfer") == "iya"
		tahunAjar := r.FormValue("tahunAjar")

		// Insert data login mahasiswa
		controller.ControlinsertDataLogin(email, tanggalLahir, jurusan, fmt.Sprintf("%v", transfer), tahunAjar, transfer)

		// Redirect to dashboard
		http.Redirect(w, r, "/hasildata", http.StatusSeeOther)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}





func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	email := cookie.Value

	userData := controller.ControlGetMahasiswaByEmail(email)
	if userData == nil {
		http.Redirect(w, r, "/firsttime", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("template/hasildata.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, userData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HandleHasilData(w http.ResponseWriter, r *http.Request) {
	emailCookie, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	email := emailCookie.Value

	mahasiswa := controller.ControlGetMahasiswaByEmail(email)
	if mahasiswa == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("template/hasildata.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, mahasiswa)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}




func HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Clear the email cookie
	cookie := &http.Cookie{
		Name:   "email",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // This will delete the cookie
	}
	http.SetCookie(w, cookie)

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}