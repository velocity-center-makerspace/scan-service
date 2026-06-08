package web

import (
	"door-greeter/scan_service/data"
	"log"
	"net/http"
	"time"
)

func checkPaidMembers(member_id string, w http.ResponseWriter, r *http.Request) bool {
	var isPaidMember bool
	paidMembers := data.GetPaidMembers()
	for _, paidMember := range paidMembers {
		if paidMember.MemberID != member_id {
			continue
		}
		if isActivePaidMember(paidMember.MembershipExpiration) {
			log.Println("Active membership found!")
			checkinTime := time.Now()
			data.InsertPaidMemberCheckin(paidMember, checkinTime)
			isPaidMember = true
			http.Redirect(w, r, "/success", http.StatusSeeOther)
			return isPaidMember
		} else {
			log.Println("Membership expired :(")
			checkinTime := time.Now()
			data.InsertPaidMemberCheckin(paidMember, checkinTime)
			isPaidMember = true
			http.Redirect(w, r, "/membership-expired", http.StatusSeeOther)
			return isPaidMember
		}
	}

	log.Println("Invalid paid member ID.")
	isPaidMember = false
	return isPaidMember
}

func checkUnpaidMembers(member_id string, w http.ResponseWriter, r *http.Request) bool {
	var isUnpaidMember bool
	unpaidMembers := data.GetUnpaidMembers()
	for _, unpaidMember := range unpaidMembers {
		if unpaidMember.MemberID != member_id {
			continue
		}
		if unpaidMember.MembershipActive {
			log.Println("Active membership found!")
			checkinTime := time.Now()
			data.InsertUnpaidMemberCheckin(unpaidMember, checkinTime)
			isUnpaidMember = true
			http.Redirect(w, r, "/success", http.StatusSeeOther)
			return isUnpaidMember
		} else {
			log.Println("Membership no longer active :(")
			checkinTime := time.Now()
			data.InsertUnpaidMemberCheckin(unpaidMember, checkinTime)
			isUnpaidMember = true
			http.Redirect(w, r, "/membership-inactive", http.StatusSeeOther)
			return isUnpaidMember
		}
	}

	isUnpaidMember = false
	log.Println("Invalid unpaid member ID.")
	return isUnpaidMember
}

func isActivePaidMember(expirationDate time.Time) bool {
	return expirationDate.After(time.Now())
}
