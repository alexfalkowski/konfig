@manual @grpc
Feature: Secrets

  Secrets allows the system to write secrets.

  Scenario: Write existing secrets
    Given I have a "folder" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                   | value                                               |
      | vault    | secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    When I write secrets for "existing" application
    Then I should have secrets for "existing" application
