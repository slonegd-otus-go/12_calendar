Feature: Add event
	As client of api service
	In order to understand that the user add event
	I want to receive event with id in response

	Scenario: Api service is available
		When I send "GET" request to "http://localhost:8080/events" with data 
        """
        date=2019-11-09%208:57:06
        """
		Then The response code should be 200

    Scenario: Add first event
		When I send "POST" request to "http://localhost:8080/events" with data 
        """
        {
            "date":"2019-11-11 13:11:05",
            "duration":5,
            "description":"сдать домашку"
        }
		"""
		Then The response code should be 200
        And I receive event with data
         """
        {
            "date":"2019-11-11 13:11:05",
            "duration":5,
            "description":"сдать домашку",
            "id":1
        }
		"""

    Scenario: Add same event, get new id
		When I send "POST" request to "http://localhost:8080/events" with data 
        """
        {
            "date":"2019-11-11 13:11:05",
            "duration":5,
            "description":"сдать домашку"
        }
		"""
		Then The response code should be 200
        And I receive event with data
         """
        {
            "date":"2019-11-11 13:11:05",
            "duration":5,
            "description":"сдать домашку",
            "id":2
        }
		"""