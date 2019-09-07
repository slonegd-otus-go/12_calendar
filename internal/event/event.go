package event

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Description string
	Date        time.Time
	Duration    time.Duration
}

type ID uint

type Storage struct {
	mtx    sync.RWMutex
	events map[ID]Event
	id     ID
}

func NewStorage() *Storage {
	return &Storage{events: make(map[ID]Event)}
}

func (storage *Storage) Add(event Event) ID {
	storage.mtx.Lock()
	id := storage.id
	storage.events[id] = event
	storage.id++
	storage.mtx.Unlock()
	return id
}

func (storage *Storage) Update(id ID, event Event) {
	storage.mtx.Lock()
	storage.events[id] = event
	storage.mtx.Unlock()
}

func (storage *Storage) Remove(id ID) {
	storage.mtx.Lock()
	delete(storage.events, id)
	storage.mtx.Unlock()
}

func (storage *Storage) Range(f func(id ID, event Event)) {
	storage.mtx.Lock()
	for id, event := range storage.events {
		f(id, event)
	}
	storage.mtx.Unlock()
}

func (storage *Storage) Strings() []string {
	var result []string
	storage.mtx.RLock()
	for _, event := range storage.events {
		result = append(result, fmt.Sprintf("date: %s, description: %s", event.Date.Format("2006-01-02 15:04:05"), event.Description))
	}
	storage.mtx.RUnlock()
	return result
}
