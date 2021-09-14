Feature: create list
  Eu como usuário desejo criar uma lista de musicas

  Scenario: should CREATE a list of songs
    When I send "POST" request to "/lists" with json:
      """
      {
        "name": "lista 1",
        "type": "custom",
        "songs": [
          1,
          2,
          3
        ]
      }
      """
    Then the response code should be 200
    And the response should match json:
      """
      {"name":"lista 1","type":"custom","songs":[1,2,3]}
      """

  Scenario: should CREATE a draft list of songs
    When I send "POST" request to "/lists?draft=1" with json:
      """
      {
        "name": "lista 1",
        "type": "custom",
        "songs": [
          1,
          2,
          3
        ]
      }
      """
    Then the response code should be 200
    And the response should match json:
      """
      {"name":"lista 1","type":"custom","songs":[1,2,3]}
      """

  @wip
  Scenario: should NOT CREATE a list of song without name
    When I send "POST" request to "/lists" with json:
      """
      {
        "name": "",
        "type": "custom",
        "songs": []
      }
      """
    Then the response code should be 400
    And the response should match json:
      """
      {"error":"bad arguments"}
      """

  Scenario: should error parsing json
    When I send "POST" request to "/lists" with json:
      """
      {
        "name": "",
        "type":
      }
      """
    Then the response code should be 400
    And the response should match json:
      """
      {"error":"json parse error"}
      """