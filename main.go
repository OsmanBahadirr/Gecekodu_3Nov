package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	LastName string
}

var katilimcilar []User

func registerHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "register.html")

}

func katilimcilarHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("katilimcilar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, katilimcilar)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	lastname := r.FormValue("lastname")

	tmpl, err := template.ParseFiles("registerSubmit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	U := User{
		Name:     name,
		LastName: lastname,
	}
	katilimcilar = append(katilimcilar, U)
	tmpl.Execute(w, U)

}
func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerSubmit", successHandler)
	http.HandleFunc("/katilimcilar", katilimcilarHandler)
	http.ListenAndServe(":3000", nil)
}
