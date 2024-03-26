package model

import (
	"time"
)

type Event struct {
	Id          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	append(events, e) //TODO save in a database
}
