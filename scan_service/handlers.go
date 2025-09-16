package scanservice

import (
	"net/http"
	// "os"
	"log"
	"strconv"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// home page with scan member ID or I don't have a member ID/I'm not a member at start
	member_id, err := strconv.Atoi(r.FormValue("member-id"))
	if err != nil {
		log.Print("Member ID is not a string")
		w.Header().Set("HX-Refresh", "true")
	}
	// check csv file for membership
	members := GetMembers()
	for i := range len(members) {
		if members[i].MemberID == member_id {
			// check member expiration
			// add visiting member data to db
			// send email to MAKERSPACE_EMAIL and MY_EMAIL
		}
	}
}
