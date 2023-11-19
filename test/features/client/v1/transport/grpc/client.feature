@manual @grpc
Feature: Client

  Client allows the system to download a new configuration.

  Scenario: Download existing config
    Given I have a "folder" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                   | value                                               |
      | vault    | secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    When I download the configuration for "existing" application
    Then I should have a configuration for "existing" application

  Scenario: Download missing config
    Given I have a "folder" valid setup
    And I start the system
    When I download the configuration for "missing" application
    Then I should not have a configuration for "missing" application
    And I should see a log entry of "not found" in the file "reports/missing.client.log"

  Scenario: Download invalid host in config
    Given I have a "folder" valid setup
    And I start the system
    When I download the configuration for "invalid_host" application
    Then I should not have a configuration for "invalid_host" application
    And I should see a log entry of "invalid_host: missing port in address" in the file "reports/invalid_host.client.log"

  Scenario: Download invalid content type in config
    Given I have a "folder" valid setup
    And I start the system
    When I download the configuration for "invalid_kind" application
    Then I should not have a configuration for "invalid_kind" application
    And I should see a log entry of "could not transform" in the file "reports/invalid_kind.client.log"
