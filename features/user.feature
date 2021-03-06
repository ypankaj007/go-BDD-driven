# file: user.feature
Feature: get users
  In order to get user list
  As an API user
  I need to be able to request users

  Scenario: does not allow POST method
    When I send "POST" request to "/users"
    Then the response code should be 405
    And the response should match json:
      """
      {
        "error": "Method not allowed"
      }
      """

  Scenario: should get user list
    When I send "GET" request to "/users"
    Then the response code should be 200
    And the response should match json:
      """
      [
        {
          "id": 1,
          "name": "Jack"
        }
      ]
      """