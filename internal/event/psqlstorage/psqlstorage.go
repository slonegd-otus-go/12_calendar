package psqlstorage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	// _ "github.com/lib/pq"

	"github.com/slonegd-otus-go/12_calendar/internal/event"
)

type storage struct {
	db *sqlx.DB
}

func New() *storage {
	db, err := sqlx.Open("pgx", "host=localhost user=myuser password=mypass dbname=mydb")
	if err != nil {
		log.Fatal(err)
	}
	return &storage{db}
}

func (storage *storage) Add(newEvent event.Event) event.ID {
	query := `insert into events(description, start_time, duration)
		values(:description, :start_time, :duration)
		returning id`
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	state, err := storage.db.PrepareNamedContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	var id int
	err = state.Get(&id, map[string]interface{}{
		"description": newEvent.Description,
		"start_time":  newEvent.Date,
		"duration":    newEvent.Duration.String(),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("add event %v done with id %d", newEvent, id)
	return event.ID(id)
}

func (storage *storage) Update(id event.ID, newEvent event.Event) (ok bool) {
	query := `UPDATE events
		SET description = :description, start_time = :start_time, duration = :duration
		WHERE id = :id`
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	result, err := storage.db.NamedExecContext(ctx, query, map[string]interface{}{
		"description": newEvent.Description,
		"start_time":  newEvent.Date,
		"duration":    newEvent.Duration.String(),
		"id":          id,
	})
	if err != nil {
		log.Printf("update event failed: %s", err)
		return false
	}
	qty, err := result.RowsAffected()
	if err != nil {
		log.Printf("update event failed: %s", err)
		return false
	}
	if qty != 1 {
		log.Printf("update event failed: dont have event with id %d", id)
		return false
	}
	log.Printf("update event %v with id %d", newEvent, id)
	return true
}

func (storage *storage) Remove(id event.ID) (ok bool) {
	query := `DELETE FROM events WHERE id = :id`
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	result, err := storage.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		log.Printf("remove event failed: %s", err)
		return false
	}
	qty, err := result.RowsAffected()
	if err != nil {
		log.Printf("remove event failed: %s", err)
		return false
	}
	if qty != 1 {
		log.Printf("remove event failed: dont have event with id %d", id)
		return false
	}
	log.Printf("remove event with id %d", id)
	return true
}

func (storage *storage) Get(id event.ID) (result event.Event, ok bool) {
	query := `SELECT description, start_time, duration FROM events WHERE id = :id`
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	rows, err := storage.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		log.Printf("%s, with id=%d failed: %s", query, id, err)
		return result, false
	}
	defer rows.Close()

	for rows.Next() {
		var description string
		var start_time time.Time
		var duration string
		err := rows.Scan(&description, &start_time, &duration)
		if err != nil {
			log.Printf("rows scan failed: %s", err)
			return result, false
		}
		resultDuration, err := parseDuration(duration)
		if err != nil {
			log.Printf("parse duration failed: %s", err)
			return result, false
		}
		return event.Event{
			Description: description,
			Date:        start_time,
			Duration:    resultDuration,
		}, true
	}
	return result, false
}

func parseDuration(s string) (time.Duration, error) {
	s = strings.Replace(s, ":", "h", 1)
	s = strings.Replace(s, ":", "m", 1)
	s = fmt.Sprintf("%ss", s)
	return time.ParseDuration(s)
}

func (storage *storage) Active(date time.Time) map[event.ID]event.Event {
	events := make(map[event.ID]event.Event)
	query :=
		`SELECT id, description, start_time, duration 
		 FROM events
		 WHERE :d BETWEEN start_time AND start_time + duration`
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	rows, err := storage.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"d": date,
	})
	if err != nil {
		log.Printf("%s, with date=%s failed: %s", query, date, err)
		return events
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var description string
		var date time.Time
		var duration string
		err := rows.Scan(&id, &description, &date, &duration)
		if err != nil {
			log.Printf("rows scan failed: %s", err)
			continue
		}
		resultDuration, err := parseDuration(duration)
		if err != nil {
			log.Printf("parse duration failed: %s", err)
			continue
		}
		events[event.ID(id)] = event.Event{
			Description: description,
			Date:        date,
			Duration:    resultDuration,
		}
	}

	return events
}
