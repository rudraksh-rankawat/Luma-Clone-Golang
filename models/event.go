package models

import (
	"database/sql"
	"errors"
	"example.com/rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() (int64, error) {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmp, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmp.Close()
	result, err := stmp.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return id, err
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, nil
		}
		return Event{}, err
	}

	return event, nil

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	return events, nil

}

func (e Event) Update() error {

	query := `
	UPDATE events
	SET name = ?, location = ?, dateTime = ?, description = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Location, e.DateTime, e.Description, e.ID)

	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err

}

func DeleteAllEvents() error {
	query := `DELETE FROM events`
	_, err := db.DB.Exec(query)
	return err
}
