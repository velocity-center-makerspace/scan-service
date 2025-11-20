package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var dbFile string = "test.db"
var db *sql.DB

func DatabaseInit() {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Panic("Database closure failed:", err)
		}
	}()

	createPaidVisitors := `
	CREATE TABLE IF NOT EXISTS paid_visitors (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		member_id INTEGER NOT NULL,
		first_name TEXT NOT NULL,
		membership_expiration DATE NOT NULL,
		checkin_time DATETIME NOT NULL
	)
	;
	`

	createUnpaidVisitors := `
	CREATE TABLE IF NOT EXISTS unpaid_visitors (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		member_id INTEGER NOT NULL,
		first_name TEXT NOT NULL,
		membership_active INTEGER NOT NULL,
		checkin_time DATETIME NOT NULL
	)
	`

	_, err = db.Exec(createPaidVisitors)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createUnpaidVisitors)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'paid_visitors' created successfully!")
	log.Println("Table 'unpaid_visitors' created successfully!")
}

func InsertPaidMemberCheckin(member *PaidMember, checkinTime time.Time) {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Panic("Database closure failed:", err)
		}
	}()

	query := `
	INSERT INTO paid_visitors 
	(member_id, first_name, membership_expiration, checkin_time) 
	VALUES (?, ?, ?, ?);
	`

	_, err = db.Exec(
		query,
		member.MemberID,
		member.FirstName,
		member.MembershipExpiration,
		checkinTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Member checked in successfully!")
}

func InsertUnpaidMemberCheckin(member *UnpaidMember, checkinTime time.Time) {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Panic("Database closure failed:", err)
		}
	}()

	query := `
	INSERT INTO unpaid_visitors 
	(member_id, first_name, membership_active, checkin_time)
	VALUES (?, ?, ?, ?);
	`

	_, err = db.Exec(
		query,
		member.MemberID,
		member.FirstName,
		member.MembershipActive,
		checkinTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Member checked in successfully!")
}
