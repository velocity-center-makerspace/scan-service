package web

import (
	"net/http"
	"time"
	// "os"
	"door-greeter/scan_service/data"
	"log"
	//"strconv"
)

func ScanInHandler(w http.ResponseWriter, r *http.Request) {
	// home page with scan member ID or I don't have a member ID/I'm not a member at start
	member_id := r.FormValue("member-id")
	/*if err != nil {
		log.Print("Invalid member ID input")
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}*/

	// check csv file for membership
	members := data.GetMembers()
	for _, member := range members {
		if member.MemberID != member_id {
			log.Printf("Member: %s %s", member.FirstName, member.MemberID)
			continue
		}
		if isActiveMember(member.MembershipExpiration) {
			data.InsertMemberCheckin(member)
			// TODO: send email to MAKERSPACE_EMAIL and MY_EMAIL
			return
		} else {
			log.Println("Membership expired :(")
			data.InsertMemberCheckin(member)
			// TODO: send email to MAKERSPACE_EMAIL and MY_EMAIL
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
