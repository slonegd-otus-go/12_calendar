package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/slonegd-otus-go/12_calendar/internal/event/amqpsubscriber"
	"github.com/stretchr/testify/assert"
)

// for use testify
type t struct{}

func (t t) Errorf(format string, args ...interface{}) {}
func (t t) FailNow()                                  {}

type eventTest struct {
	responseStatusCode int
	event              []byte
	message            string
	t                  t
}

func NewEventTest() *eventTest {
	eventTest := &eventTest{}
	go amqpsubscriber.Run("amqp://guest:guest@localhost:5672", func(message string) {
		eventTest.message = message
	})
	return eventTest
}

func (test *eventTest) iSendRequestToWithData(method, addr string, data *gherkin.DocString) error {
	switch method {
	case http.MethodGet:
		request := addr + "?" + data.Content
		response, err := http.Get(request)
		if err != nil {
			return fmt.Errorf("GET method failed: %s", err)
		}
		test.responseStatusCode = response.StatusCode
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %s", err)
		}
		test.event = bytes
		return nil

	case http.MethodPost:
		response, err := http.Post(addr, "application/json", strings.NewReader(data.Content))
		if err != nil {
			return fmt.Errorf("POST method failed: %s", err)
		}
		test.responseStatusCode = response.StatusCode
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %s", err)
		}
		test.event = bytes
		return nil

	default:
		return fmt.Errorf("unknown method: %s", method)
	}

}

func (test *eventTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *eventTest) iReceiveEventsWithData(body *gherkin.DocString) error {
	var expected, actual []interface{}

	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return fmt.Errorf("unmarshal actual event failed: %s", err)
	}

	if err := json.Unmarshal(test.event, &actual); err != nil {
		return fmt.Errorf("unmarshal expected event failed: %s", err)
	}

	if ok := assert.ElementsMatch(test.t, expected, actual); !ok {
		return fmt.Errorf("expected JSON does not match actual, %v != %v", expected, actual)
	}
	return nil
}

func (test *eventTest) iReceiveEventWithData(body *gherkin.DocString) error {
	var expected, actual interface{}

	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return fmt.Errorf("unmarshal actual event failed: %s", err)
	}

	if err := json.Unmarshal(test.event, &actual); err != nil {
		return fmt.Errorf("unmarshal expected event failed: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v != %v", expected, actual)
	}
	return nil
}

func (test *eventTest) iSendRequestToWithNowPlusSecondsDateAndDescription(arg1, arg2 string, arg3 int, arg4 string) error {
	return godog.ErrPending
}

func (test *eventTest) iReceiveMessage(arg1 string) error {
	return godog.ErrPending
}

func (test *eventTest) iWaitSeconds(arg1 int) error {
	return godog.ErrPending
}

func (test *eventTest) iGotMessage(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	test := NewEventTest()

	s.Step(`^I send "([^"]*)" request to "([^"]*)" with data$`, test.iSendRequestToWithData)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^I receive event with data$`, test.iReceiveEventWithData)
	s.Step(`^I receive events with data$`, test.iReceiveEventsWithData)

	s.Step(`^I send "([^"]*)" request to "([^"]*)" with now plus (\d+) seconds date and description "([^"]*)"$`, test.iSendRequestToWithNowPlusSecondsDateAndDescription)
	s.Step(`^I receive message "([^"]*)"$`, test.iReceiveMessage)
	s.Step(`^I wait (\d+) seconds$`, test.iWaitSeconds)
	s.Step(`^I got message "([^"]*)"$`, test.iGotMessage)
}
