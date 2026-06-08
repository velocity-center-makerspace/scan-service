package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type memberID struct {
	ID string `json:"id"`
}

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	var memberData memberID

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&memberData)
	if err != nil {
		log.Println("Unable to unmarshal request body.")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Unable to close request body. Error: %v", err)
		}
	}()

	member_id := memberData.ID
	isPaidMember := checkPaidMembers(member_id, w, r)
	if isPaidMember {
		return
	}
	checkUnpaidMembers(member_id, w, r)
}
