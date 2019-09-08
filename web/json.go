package web

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/slonegd-otus-go/12_calendar/web/models"
)

type JSONConsumer struct{}
type JSONProducer struct{}

func (producer JSONProducer) Produce(writer io.Writer, v interface{}) error {
	switch v.(type) {
	case *models.Event:
		event, _ := v.(*models.Event)
		bytes, err := event.MarshalJSON()
		if err != nil {
			return err
		}
		_, err = writer.Write(bytes)
		return err

	case []*models.Event:
		events, _ := v.([]*models.Event)
		_, err := writer.Write([]byte("["))
		if err != nil {
			return err
		}
		for i, event := range events {
			err = producer.Produce(writer, event)
			if err != nil {
				return err
			}
			if i != len(events)-1 {
				_, err := writer.Write([]byte(","))
				if err != nil {
					return err
				}
			}
		}
		_, err = writer.Write([]byte("]"))
		return err
	default:
		return errors.New("type not supported by JSONProducer")
	}
}

func (JSONConsumer) Consume(reader io.Reader, v interface{}) error {
	switch v.(type) {
	case *models.Event:
		event, _ := v.(*models.Event)
		bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		err = event.UnmarshalJSON(bytes)
		return err

	default:
		return errors.New("type not supported by JSONProducer")
	}
}
