package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	EventId int `json:"eventid"`
	UserId  int `json:"userid"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO attendees (user_id, event_id) VALUES ($1, $2) RETURNING id`
	err := m.DB.QueryRowContext(ctx, query, attendee.UserId, attendee.EventId).Scan(&attendee.ID)
	if err != nil {
		return nil, err
	}
	return attendee, nil
}

func (m *AttendeeModel) GetByEventAndAttendee(eventID, userID int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, user_id, event_id FROM attendees WHERE event_id = $1 AND user_id = $2`
	row := m.DB.QueryRowContext(ctx, query, eventID, userID)

	var attendee Attendee
	err := row.Scan(&attendee.ID, &attendee.UserId, &attendee.EventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &attendee, nil
}

func (m *AttendeeModel) GetAttendeesByEvent(eventID int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT u.id, u.name, u.email FROM users u
	JOIN attendees a ON a.user_id = u.id WHERE a.event_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *AttendeeModel) Delete(userID, eventID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM attendees WHERE user_id = $1 AND event_id = $2`
	_, err := m.DB.ExecContext(ctx, query, userID, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (m *AttendeeModel) GetEventsByAttendee(userID int) ([]Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT e.id, e.owner_id, e.name, e.description, e.date, e.location FROM events e
	JOIN attendees a ON a.event_id = e.id WHERE a.user_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
