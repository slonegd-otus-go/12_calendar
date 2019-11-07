package event

import (
	"time"
)

type ID uint64

type Event struct {
	Description string
	Date        time.Time
	Duration    time.Duration
}

type Storage interface {
	Add(Event) ID
	Update(ID, Event) (ok bool)
	Remove(ID) (ok bool)
	Active(time.Time) map[ID]Event
	Get(ID) (event Event, ok bool)
}
