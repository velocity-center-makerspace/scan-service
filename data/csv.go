package data

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var paidMemberCsv string = "tbl_members_paid.csv"
var unpaidMemberCsv string = "tbl_members_unpaid.csv"

type PaidMember struct {
	MemberID                string    `csv:"MemberID"`
	FirstName               string    `csv:"FirstName"`
	MembershipExpirationStr string    `csv:"MembershipExpirationDate"`
	MembershipExpiration    time.Time `csv:"-"` // calculated field
}

type UnpaidMember struct {
	MemberID         string `csv:"MemberID"`
	FirstName        string `csv:"FirstName"`
	MembershipActive bool   `csv:"MembershipActive"`
}

func removeQuotes(filename string) (io.Reader, error) {
	var err error

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cleanData := bytes.ReplaceAll(data, []byte(`"`), []byte(""))

	return bytes.NewReader(cleanData), nil
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

	log.Println(gocsv.MarshalString(members))
	return members
}

func timeConversion(expirationStr string, member *PaidMember) {
	var err error
	dateList := strings.Split(expirationStr, "/")
	month := fmt.Sprintf("%02s", dateList[0])
	day := fmt.Sprintf("%02s", dateList[1])
	year := dateList[2]

	formattedDate := fmt.Sprintf("%s-%s-%s", year, month, day)
	member.MembershipExpiration, err = time.Parse(time.DateOnly, formattedDate)
	if err != nil {
		log.Fatalf("Error parsing member expiration: %v", err)
	}
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

	log.Println(gocsv.MarshalString(members))
	return members
}
