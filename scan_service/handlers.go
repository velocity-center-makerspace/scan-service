package scanservice

import (
	"net/http"
	// "os"
	"log"
	"slices"
	"strconv"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// home page with scan member ID or I don't have a member ID/I'm not a member at start
	member_id, err := strconv.Atoi(r.FormValue("member-id"))
	if err != nil {
		log.Print("Member ID is not a string")
		// TODO: change to HX-Refresh
		return
	}
	// check csv file for membership
	members := GetMembers()
	var member_ids []int
	for _, member := range members {
		member_ids = append(member_ids, member.MemberID)
	}

	if slices.Contains(member_ids, member_id) {
		// check member expiration
		// add visiting member data to db
		// send email to MAKERSPACE_EMAIL and MY_EMAIL
	}
}
