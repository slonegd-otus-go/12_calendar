Feature: Add event
	As client of api service
	In order to understand that the user add event
	I want to receive event with id in response

	Scenario: Api service is available
		When I send "GET" request to "http://localhost:8080/events" with "date=2019-11-09%208:57:06" in path
		Then The response code should be 200