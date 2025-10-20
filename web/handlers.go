package web

import (
	"net/http"
	"time"

	"door-greeter/scan_service/data"
	"log"
)

func ScanInHandler(w http.ResponseWriter, r *http.Request) {
	// home page with scan member ID or I don't have a member ID/I'm not a member at start
	member_id := r.FormValue("member-id")

	// check csv file for membership
	members := data.GetMembers()
	for _, member := range members {
		if member.MemberID != member_id {
			continue
		}
		if isActiveMember(member.MembershipExpiration) {
			checkinTime := time.Now()
			data.InsertMemberCheckin(member, checkinTime)
			return
		} else {
			log.Println("Membership expired :(")
			checkinTime := time.Now()
			data.InsertMemberCheckin(member, checkinTime)
			http.Redirect(w, r, "/membership-expired", http.StatusSeeOther)
			return
		}
	}
	log.Println("Invalid member ID.")
	http.Redirect(w, r, "/invalid-member-id", http.StatusSeeOther)
}

func isActiveMember(expirationDate time.Time) bool {
	return expirationDate.After(time.Now())
}
