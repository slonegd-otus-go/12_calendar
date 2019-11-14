Feature: Update event
	As client of scheduler service
	In order to understand that the user recieve event
	I want to receive event at start time

	Scenario: Recieve Event in start time
		When I send "POST" request to "http://localhost:8080/events" with now plus 3 seconds date and description "сдать домашку" 
        Then The response code should be 200
        And I receive message ""
        When I wait 1 seconds
        Then I got message ""
        When I wait 1 seconds
        Then I got message "сдать домашку"
        When I wait 1 seconds
        Then I got message ""

  