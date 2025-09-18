package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	EventId int `json:"eventid"`
	UserId  int `json:"userid"`
}
