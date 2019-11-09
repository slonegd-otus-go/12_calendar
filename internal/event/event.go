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

func StartScan(storage Storage, onEvent func(Event)) {
	go Scan(storage, onEvent)
}

func Scan(storage Storage, onEvent func(Event)) {
	ticker := time.NewTicker(1 * time.Second)
	publishedIDs := make(IDset)
	for range ticker.C {
		events := storage.Active(time.Now())
		ids := makeIDset(events)
		publishedIDs = intersection(publishedIDs, ids)
		for id, event := range events {
			if publishedIDs.contains(id) {
				continue
			}
			onEvent(event)
			publishedIDs.add(id)
		}
	}
}

type IDset map[ID]struct{}

func makeIDset(ids map[ID]Event) IDset {
	result := make(IDset)
	for id := range ids {
		result.add(id)
	}
	return result
}

func (set IDset) contains(id ID) bool {
	_, ok := set[id]
	return ok
}

func (set IDset) add(id ID) {
	set[id] = struct{}{}
}

func (set IDset) complement(other IDset) {
	for v := range other {
		delete(set, v)
	}
}

// https://en.wikipedia.org/wiki/Intersection_(set_theory)
func intersection(ids1, ids2 IDset) IDset {
	result := make(IDset)
	for id := range ids1 {
		if ids2.contains(id) {
			result.add(id)
		}
	}
	return result
}
