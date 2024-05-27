package models

import (
	"database/sql"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserId      string
}

func (e Event) Save(db *sql.DB) error {
	query := `INSERT INTO events(
		name, 
		description, 
		location, 
		time, 
		user_id
	  ) 
	  VALUES 
		(?, ?, ?, ?, ?)`
	cur, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer cur.Close()
	result, err := cur.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.Time,
		e.UserId,
	)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	// e.Id = id
	return err
}

func (e *Event) Update(db *sql.DB) error {
	query := `UPDATE events
		SET name = ?, description = ?, location = ?, time = ?
		WHERE id = ?
	`
	cur, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer cur.Close()
	_, err = cur.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.Time,
		e.Id,
	)
	return err
}

func (e *Event) Delete(db *sql.DB) error {
	query := "DELETE FROM events WHERE id = ?"
	cur, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer cur.Close()
	_, err = cur.Exec(e.Id)
	return err
}

func GetAllEvents(db *sql.DB) ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Time,
			&event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(eventId int64, db *sql.DB) (*Event, error) {

	query := "SELECT * FROM events WHERE  id = ?"

	row := db.QueryRow(query, eventId)

	var event Event
	err := row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.Time,
		&event.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
