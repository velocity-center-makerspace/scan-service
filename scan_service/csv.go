package scanservice

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"time"
)

type Member struct {
	MemberID                int       `csv:"Member ID"`
	FirstName               string    `csv:"First Name"`
	MembershipExpirationStr string    `csv:"Membership Expiration Date"`
	MembershipExpiration    time.Time `csv:"-"` // calculated field
}

var csvFileName string = "../tbl_members.csv"

func GetMembers() []*Member {
	var err error
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = csvFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	gocsv.SetCSVReader(gocsv.LazyCSVReader)

	members := []*Member{}
	if err = gocsv.UnmarshalFile(csvFile, &members); err != nil {
		log.Fatal(err)
	}

	return members
}
