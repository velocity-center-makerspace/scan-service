package data

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

type Member struct {
	MemberID                string    `csv:"MemberID"`
	FirstName               string    `csv:"FirstName"`
	MembershipExpirationStr string    `csv:"MembershipExpirationDate"`
	MembershipExpiration    time.Time `csv:"-"` // calculated field
}

var csvFileName string = "tbl_members.csv"

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

	members := []*Member{}
	if err = gocsv.UnmarshalFile(csvFile, &members); err != nil {
		log.Fatal(err)
	}

	for _, member := range members {
		timeConversion(member.MembershipExpirationStr, member)
	}

	log.Println(gocsv.MarshalString(members))
	return members
}

func timeConversion(expirationStr string, member *Member) {
	var err error
	dateList := strings.Split(expirationStr, "/")
	month := fmt.Sprintf("%02s", dateList[0])
	day := fmt.Sprintf("%02s", dateList[1])
	year := dateList[2]

	formattedDate := fmt.Sprintf("20%s-%s-%s", year, month, day)
	member.MembershipExpiration, err = time.Parse(time.DateOnly, formattedDate)
	if err != nil {
		log.Fatalf("Error parsing member expiration: %v", err)
	}
}
