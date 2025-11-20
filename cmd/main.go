package main

import (
	"door-greeter/scan_service/data"
	"door-greeter/scan_service/web"
	"html/template"
	"log"
	"net/http"
)

func main() {
	data.DatabaseInit()

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /see-coordinator", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/see-coordinator.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /membership-expired", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/membership-expired.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /member-inactive", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/member-inactive.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /invalid-member-id", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/invalid-member-id.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /success", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/success.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /error", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc(
		"POST /scan-in",
		web.ScanInHandler,
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
