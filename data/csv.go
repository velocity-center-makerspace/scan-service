package data

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func removeQuotes(filename string) (io.Reader, error) {
	var err error

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cleanData := bytes.ReplaceAll(data, []byte(`"`), []byte(""))

	return bytes.NewReader(cleanData), nil
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
