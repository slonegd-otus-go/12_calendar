Feature: Update event
	As client of api service
	In order to understand that the user update event
	I want to receive ok status and event must change

	Scenario: Get events from add event scenario
		When I send "GET" request to "http://localhost:8080/events" with data 
        """
        date=2019-11-11%2013:11:06
        """
		Then The response code should be 200
        And I receive events with data
        """
        [
            {
                "date":"2019-11-11 13:11:05",
                "duration":5,
                "description":"сдать домашку",
                "id":1
            },
            {
                "date":"2019-11-11 13:11:05",
                "duration":5,
                "description":"сдать домашку",
                "id":2
            }
        ]
		"""

    Scenario: Update event with id 1
		When I send "POST" request to "http://localhost:8080/events/update/1" with data 
        """
            {
                "date":"2019-11-11 13:11:05",
                "duration":5,
                "description":"получить максимальный бал за домашку"
            }
        """
		Then The response code should be 200
        When I send "GET" request to "http://localhost:8080/events" with data 
        """
        date=2019-11-11%2013:11:06
        """
		Then The response code should be 200
        And I receive events with data
        """
        [
            {
                "date":"2019-11-11 13:11:05",
                "duration":5,
                "description":"получить максимальный бал за домашку",
                "id":1
            },
            {
                "date":"2019-11-11 13:11:05",
                "duration":5,
                "description":"сдать домашку",
                "id":2
            }
        ]
		"""


