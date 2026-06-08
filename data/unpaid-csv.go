package data

import (
	"github.com/gocarina/gocsv"
	"log"
)

var unpaidMemberCsv string = "tbl_members_unpaid.csv"

type UnpaidMember struct {
	MemberID         string `csv:"MemberID"`
	FirstName        string `csv:"FirstName"`
	MembershipActive bool   `csv:"MembershipActive"`
}

func GetUnpaidMembers() []*UnpaidMember {
	var err error

	csvReader, err := removeQuotes(unpaidMemberCsv)
	if err != nil {
		log.Fatal(err)
	}

	members := []*UnpaidMember{}
	if err = gocsv.Unmarshal(csvReader, &members); err != nil {
		log.Fatal(err)
	}

	return members
}
