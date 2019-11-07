package mapstorage

import (
	"fmt"
	"sync"
	"time"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

type storage struct {
	mtx    sync.RWMutex
	events map[event.ID]event.Event
	id     event.ID
}

func New() *storage {
	return &storage{events: make(map[event.ID]event.Event), id: 1}
}

func (storage *storage) Add(event event.Event) event.ID {
	storage.mtx.Lock()
	id := storage.id
	storage.events[id] = event
	storage.id++
	storage.mtx.Unlock()
	return id
}

func (storage *storage) Update(id event.ID, event event.Event) (ok bool) {
	storage.mtx.Lock()
	defer storage.mtx.Unlock()
	_, ok = storage.events[id]
	if !ok {
		return ok
	}
	storage.events[id] = event
	return ok
}

func (storage *storage) Remove(id event.ID) (ok bool) {
	storage.mtx.Lock()
	_, ok = storage.events[id]
	delete(storage.events, id)
	storage.mtx.Unlock()
	return ok
}

func (storage *storage) Active(date time.Time) map[event.ID]event.Event {
	events := make(map[event.ID]event.Event)
	storage.rangeEvents(func(id event.ID, event event.Event) bool {
		if date.After(event.Date) && event.Date.Add(event.Duration).After(date) {
			events[id] = event
		}
		return true
	})
	return events
}

func (storage *storage) Get(id event.ID) (event event.Event, ok bool) {
	storage.mtx.RLock()
	event, ok = storage.events[id]
	storage.mtx.RUnlock()
	return event, ok
}

func (storage *storage) Strings() []string {
	var result []string
	storage.mtx.RLock()
	for _, event := range storage.events {
		result = append(result, fmt.Sprintf("date: %s, description: %s", event.Date.Format("2006-01-02 15:04:05"), event.Description))
	}
	storage.mtx.RUnlock()
	return result
}

func (storage *storage) rangeEvents(f func(id event.ID, event event.Event) (ok bool)) {
	storage.mtx.Lock()
	for id, event := range storage.events {
		if !f(id, event) {
			break
		}
	}
	storage.mtx.Unlock()
}
