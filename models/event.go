package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserId      string
}

var events = []Event{}

func (e Event) Save() {
	//TODO: add to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
