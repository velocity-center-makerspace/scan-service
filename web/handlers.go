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
	paidMembers := data.GetPaidMembers()
	for _, paidMember := range paidMembers {
		if paidMember.MemberID != member_id {
			continue
		}
		if isActivePaidMember(paidMember.MembershipExpiration) {
			checkinTime := time.Now()
			data.InsertPaidMemberCheckin(paidMember, checkinTime)
			http.Redirect(w, r, "/success", http.StatusSeeOther)
			return
		} else {
			log.Println("Membership expired :(")
			checkinTime := time.Now()
			data.InsertPaidMemberCheckin(paidMember, checkinTime)
			http.Redirect(w, r, "/membership-expired", http.StatusSeeOther)
			return
		}
	}

	unpaidMembers := data.GetUnpaidMembers()
	for _, unpaidMember := range unpaidMembers {
		if unpaidMember.MemberID != member_id {
			continue
		}
		if unpaidMember.MembershipActive {
			checkinTime := time.Now()
			data.InsertUnpaidMemberCheckin(unpaidMember, checkinTime)
			http.Redirect(w, r, "/success", http.StatusSeeOther)
			return
		} else {
			log.Println("Membership no longer active :(")
			checkinTime := time.Now()
			data.InsertUnpaidMemberCheckin(unpaidMember, checkinTime)
			http.Redirect(w, r, "/membership-inactive", http.StatusSeeOther)
		}
	}

	log.Println("Invalid member ID.")
	http.Redirect(w, r, "/invalid-member-id", http.StatusSeeOther)
}

func isActivePaidMember(expirationDate time.Time) bool {
	return expirationDate.After(time.Now())
}
