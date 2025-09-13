package main

import (
	"log"
	// "net/http"
	sc "door-greeter/scan_service"
)

func main() {
	sc.DatabaseInit()
	members := sc.GetMembers()

	for _, member := range members {
		log.Println(member.FirstName)
	}
}
