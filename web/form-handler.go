package web

import (
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	member_id := r.FormValue("member-id")

	isPaidMember := checkPaidMembers(member_id, w, r)
	if isPaidMember {
		return
	}
	checkUnpaidMembers(member_id, w, r)
	http.Redirect(w, r, "/invalid-member-id", http.StatusSeeOther)
}
