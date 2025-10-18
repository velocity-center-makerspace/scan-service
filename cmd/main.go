package main

import (
	//"context"
	"door-greeter/scan_service/data"
	"door-greeter/scan_service/web"
	"log"
	"net/http"
	"os"
	//"golang.org/x/oauth2/google"
	//"google.golang.org/api/gmail/v1"
	//"google.golang.org/api/option"
)

func main() {
	data.DatabaseInit()

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))
	http.Handle("/static/", fileHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("GET /see-coordinator", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("templates/see-coordinator.html")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(content)
	})

	http.HandleFunc("GET /membership-expired", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/membership-expired.html")
	})

	http.HandleFunc("GET /invalid-member-id", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/invalid-member-id.html")
	})

	http.HandleFunc("GET /success", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/success.html")
	})

	http.HandleFunc("GET /error", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/error.html")
	})

	http.HandleFunc("POST /scan-in", web.ScanInHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
	   ctx := context.Background()
	   b, err := os.ReadFile("../credentials.json")

	   	if err != nil {
	   		log.Fatalf("Unable to read client secret file: %v", err)
	   	}

	   // If modifying these scopes, delete your previously saved token.json.
	   config, err := google.ConfigFromJSON(b, gmail.GmailModifyScope)

	   	if err != nil {
	   		log.Fatalf("Unable to parse client secret file to config: %v", err)
	   	}

	   client := web.GetClient(config)

	   srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))

	   	if err != nil {
	   		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	   	}

	   user := "me"
	   r, err := srv.Users.Labels.List(user).Do()

	   	if err != nil {
	   		log.Fatalf("Unable to retrieve labels: %v", err)
	   	}

	   	if len(r.Labels) == 0 {
	   		log.Println("No labels found.")
	   		return
	   	}

	   log.Println("Labels:")

	   	for _, l := range r.Labels {
	   		log.Printf("- %s\n", l.Name)
	   	}
	*/
}
