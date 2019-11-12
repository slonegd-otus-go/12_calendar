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
)

type addEventTest struct {
	responseStatusCode int
	event              []byte
}

func (test *addEventTest) iSendRequestToWithData(method, addr string, data *gherkin.DocString) error {
	switch method {
	case http.MethodGet:
		request := addr + "?" + data.Content
		response, err := http.Get(request)
		if err != nil {
			return fmt.Errorf("GET method failed: %s", err)
		}
		test.responseStatusCode = response.StatusCode
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

func (test *addEventTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *addEventTest) iReceiveEventWithData(event *gherkin.DocString) error {
	var expected, actual interface{}

	if err := json.Unmarshal([]byte(event.Content), &expected); err != nil {
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

func FeatureContext(s *godog.Suite) {
	test := addEventTest{}

	s.Step(`^I send "([^"]*)" request to "([^"]*)" with data$`, test.iSendRequestToWithData)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^I receive event with data$`, test.iReceiveEventWithData)
}
