package models

import (
	"database/sql"
	"errors"
	"example.com/rest-api/db"
)

type Event struct {
	ID           int64  `json:"id"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Date         string `json:"date" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time,omitempty"`
	ImageURL     string `json:"image_url" binding:"required"`
	MaxAttendees int    `json:"max_attendees"`
	Organizer    string `json:"organizer"`
	UserID       int64  `json:"user_id"`
}

func (e *Event) Save() (int64, error) {
	query := `
	INSERT INTO events(name, description, location, date, start_time, end_time, image_url, max_attendees, organizer, user_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.Date,
		e.StartTime,
		e.EndTime,
		e.ImageURL,
		e.MaxAttendees,
		e.Organizer,
		e.UserID,
	)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return id, err
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT id, name, description, location, date, start_time, end_time, image_url, max_attendees, organizer, user_id 
			  FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.Date,
		&event.StartTime,
		&event.EndTime,
		&event.ImageURL,
		&event.MaxAttendees,
		&event.Organizer,
		&event.UserID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, nil
		}
		return Event{}, err
	}

	return event, nil
}

func GetAllEvents(search string, date string) ([]Event, error) {
	query := `
	SELECT id, name, description, location, date, start_time, end_time, image_url, max_attendees, organizer, user_id
	FROM events
	WHERE 1=1
	`

	var args []interface{}

	if search != "" {
		query += " AND (name LIKE ? OR description LIKE ?)"
		searchParam := "%" + search + "%"
		args = append(args, searchParam, searchParam)
	}

	if date != "" {
		query += " AND date = ?"
		args = append(args, date)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Date,
			&event.StartTime,
			&event.EndTime,
			&event.ImageURL,
			&event.MaxAttendees,
			&event.Organizer,
			&event.UserID,
		)
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
	SET name = ?, description = ?, location = ?, date = ?, start_time = ?, end_time = ?, image_url = ?, max_attendees = ?, organizer = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.Date,
		e.StartTime,
		e.EndTime,
		e.ImageURL,
		e.MaxAttendees,
		e.Organizer,
		e.ID,
	)

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
