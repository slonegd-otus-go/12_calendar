package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DATA-DOG/godog"
)

type addEventTest struct {
	responseStatusCode int
}

func (test *addEventTest) iSendRequestToWithInPath(method, addr, params string) error {

	if method != http.MethodGet {
		return fmt.Errorf("unknown method: %s", method)
	}

	request := addr + "?" + params
	log.Printf(request)
	response, err := http.Get(request)
	if err != nil {
		return fmt.Errorf("GET method failed: %s", err)
	}

	test.responseStatusCode = response.StatusCode
	return nil
}

func (test *addEventTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := addEventTest{}

	s.Step(`^I send "([^"]*)" request to "([^"]*)" with "([^"]*)" in path$`, test.iSendRequestToWithInPath)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
}
