package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var dbFile string = "../test.db"
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

	createTable := `
	CREATE TABLE IF NOT EXISTS visitors (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		member_id INTEGER NOT NULL,
		first_name TEXT NOT NULL,
		membership_expiration DATE NOT NULL,
		checkin_time DATETIME NOT NULL
	);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'visitors' created successfully!")
}

func InsertMemberCheckin(member *Member) {
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
	INSERT INTO visitors 
	(member_id, first_name, membership_expiration, checkin_time) 
	VALUES (?, ?, ?, ?);
	`

	_, err = db.Exec(
		query,
		member.MemberID,
		member.FirstName,
		member.MembershipExpiration,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Member checked in successfully!")
}
