package server

import (
	"door-greeter/scan_service/web"
	"html/template"
	"log"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	mux.Handle("/static/", fileHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load index.html: ", err)
		}
	})

	mux.HandleFunc("GET /see-coordinator", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/see-coordinator.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load see-coordinator.html: ", err)
		}
	})

	mux.HandleFunc("GET /membership-expired", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/membership-expired.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load membership-expired.html: ", err)
		}
	})

	mux.HandleFunc("GET /membership-inactive", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/membership-inactive.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load membership-inactive.html", err)
		}
	})

	mux.HandleFunc("GET /invalid-member-id", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/invalid-member-id.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load membership-inactive.html", err)
		}
	})

	mux.HandleFunc("GET /success", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/success.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load success.html")
		}
	})

	mux.HandleFunc("GET /error", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Unable to load error.html")
		}
	})

	mux.HandleFunc(
		"POST /scan-in",
		web.ScanInHandler,
	)
}
