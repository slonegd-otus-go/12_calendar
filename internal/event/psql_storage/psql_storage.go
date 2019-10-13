package psql_storage

import (
	"context"
	"log"
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
	result, err := storage.db.PrepareNamedContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	var id int
	err = result.Get(&id, map[string]interface{}{
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

func (storage *storage) Update(id event.ID, event event.Event) (ok bool) {
	return true
}

func (storage *storage) Remove(id event.ID) (ok bool) {
	return true
}

func (storage *storage) Range(f func(id event.ID, event event.Event) (ok bool)) {

}

func (storage *storage) Get(id event.ID) (event event.Event, ok bool) {
	return event, ok
}
