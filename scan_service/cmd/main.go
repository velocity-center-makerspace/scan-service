package main

import (
	"context"
	sc "door-greeter/scan_service"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func main() {
	sc.DatabaseInit()
	members := sc.GetMembers()

	for _, member := range members {
		log.Println(member.FirstName, member.MembershipExpiration)
	}

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../static/")))
	http.Handle("/static/", fileHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/index.html")
	})

	http.HandleFunc("/see-coordinator", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/see-coordinator.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

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
	client := sc.GetClient(config)

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
}
