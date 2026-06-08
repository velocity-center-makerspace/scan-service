package data

import (
	"github.com/gocarina/gocsv"
	"log"
	"time"
)

var paidMemberCsv string = "tbl_members_paid.csv"

type PaidMember struct {
	MemberID                string    `csv:"MemberID"`
	FirstName               string    `csv:"FirstName"`
	MembershipExpirationStr string    `csv:"MembershipExpirationDate"`
	MembershipExpiration    time.Time `csv:"-"` // calculated field
}

func GetPaidMembers() []*PaidMember {
	var err error

	csvReader, err := removeQuotes(paidMemberCsv)
	if err != nil {
		log.Fatal(err)
	}

	members := []*PaidMember{}
	if err = gocsv.Unmarshal(csvReader, &members); err != nil {
		log.Fatal(err)
	}

	for _, member := range members {
		timeConversion(member.MembershipExpirationStr, member)
	}

	return members
}
