package event_test

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

type Action int

const (
	add Action = iota
	remove
	update
	wait
)

type action struct {
	sync.RWMutex
	action Action
	event  event.Event
	i      int // индекс для удаления/изменения
	id     event.ID
}

func TestStorage(t *testing.T) {
	tests := []struct {
		name    string
		actions []action
		want    []string
	}{
		{
			name: "just add",
			actions: []action{
				{action: add, event: Event("2019-09-07 08:20:00/event1")},
				{action: add, event: Event("2019-09-07 08:25:00/event2")},
				{action: add, event: Event("2019-09-07 09:02:00/event3")},
				{action: add, event: Event("2019-09-07 09:15:00/event4")},
				{action: wait},
			},
			want: []string{
				"date: 2019-09-07 08:20:00, description: event1",
				"date: 2019-09-07 08:25:00, description: event2",
				"date: 2019-09-07 09:02:00, description: event3",
				"date: 2019-09-07 09:15:00, description: event4",
			},
		},
		{
			name: "add than remove",
			actions: []action{
				{action: add, event: Event("2019-09-07 08:20:00/event1")},
				{action: add, event: Event("2019-09-07 08:25:00/event2")},
				{action: add, event: Event("2019-09-07 09:02:00/event3")},
				{action: add, event: Event("2019-09-07 09:15:00/event4")},
				{action: wait},
				{action: remove, i: 1},
				{action: add, event: Event("2019-09-07 09:16:00/event5")},
				{action: wait},
			},
			want: []string{
				"date: 2019-09-07 08:20:00, description: event1",
				"date: 2019-09-07 09:02:00, description: event3",
				"date: 2019-09-07 09:15:00, description: event4",
				"date: 2019-09-07 09:16:00, description: event5",
			},
		},
		{
			name: "add than update and remove",
			actions: []action{
				{action: add, event: Event("2019-09-07 08:20:00/event1")},
				{action: add, event: Event("2019-09-07 08:25:00/event2")},
				{action: add, event: Event("2019-09-07 09:02:00/event3")},
				{action: add, event: Event("2019-09-07 09:15:00/event4")},
				{action: wait},
				{action: add, event: Event("2019-09-07 09:16:00/event5")},
				{action: update, i: 1, event: Event("2019-09-07 08:25:00/event2update")},
				{action: remove, i: 2},
				{action: add, event: Event("2019-09-07 09:19:00/event6")},
				{action: wait},
			},
			want: []string{
				"date: 2019-09-07 08:20:00, description: event1",
				"date: 2019-09-07 08:25:00, description: event2update",
				"date: 2019-09-07 09:15:00, description: event4",
				"date: 2019-09-07 09:16:00, description: event5",
				"date: 2019-09-07 09:19:00, description: event6",
			},
		},
	}
	for _, tt := range tests {
		var wg sync.WaitGroup

		storage := event.NewStorage()
		for i, _ := range tt.actions {
			if tt.actions[i].action == wait {
				// time.Sleep(time.Second)
				wg.Wait()
				continue
			}

			wg.Add(1)
			go func(i int) {
				switch tt.actions[i].action {
				case add:
					id := storage.Add(tt.actions[i].event)
					tt.actions[i].Lock()
					tt.actions[i].id = id
					tt.actions[i].Unlock()
				case remove:
					storage.Remove(tt.actions[tt.actions[i].i].id)
				case update:
					storage.Update(tt.actions[tt.actions[i].i].id, tt.actions[i].event)
				}
				wg.Done()
			}(i)
		}

		assert.ElementsMatch(t, tt.want, storage.Strings(), "%s", tt.name)
	}
}

func Time(s string) time.Time {
	result, _ := time.Parse("2006-01-02 15:04:05", s)
	return result
}

func Event(s string) event.Event {
	ss := strings.Split(s, "/")
	return event.Event{
		Date:        Time(ss[0]),
		Description: ss[1],
	}
}
