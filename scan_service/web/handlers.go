package web

import (
	"net/http"
	"time"
	// "os"
	"door-greeter/scan_service/data"
	"log"
	"strconv"
)

func ScanInHandler(w http.ResponseWriter, r *http.Request) {
	// home page with scan member ID or I don't have a member ID/I'm not a member at start
	member_id, err := strconv.Atoi(r.FormValue("member-id"))
	if err != nil {
		log.Println("Invalid member ID input:", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
	// check csv file for membership
	members := data.GetMembers()
	for _, member := range members {
		if member.MemberID == member_id {
			if isActiveMember(member.MembershipExpiration) {
				data.InsertMemberCheckin(member)
				// TODO: send email to MAKERSPACE_EMAIL and MY_EMAIL
			} else {
				log.Println("Membership expired :(")
				http.Redirect(w, r, "/membership-expired", http.StatusPaymentRequired)
			}
		} else {
			log.Println("Invalid member ID.")
			http.Redirect(w, r, "/invalid-member-id", http.StatusUnauthorized)
		}
	}
}

func isActiveMember(expirationDate time.Time) bool {
	return expirationDate.After(time.Now())
}
