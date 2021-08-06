Feature: create list
  Eu como usuário desejo criar uma lista de musicas

  Scenario: should get hello world
    When I send "GET" request to "/"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "Hello World!"
      }
      """

  Scenario: should get hello world
    When I send "GET" request to "/lists"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """

  Scenario: should CREATE a list of songs
    When I send "POST" request to "/lists"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """

  Scenario: should get hello world
    When I send "GET" request to "/lists/1"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """

  Scenario: should UPDATE a list of songs
    When I send "PUT" request to "/lists/1"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """

  Scenario: should DELETE a list of songs
    When I send "DELETE" request to "/lists/1"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "message": "pong"
      }
      """