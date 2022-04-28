@manual
Feature: Client

  Client allows the system to download a new configuration.

  Scenario: Download existing config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": "Konfig-server/1.0 http/1.0" }}' value in vault
    When I download the configuration for "existing" application
    Then I should have a configuration for "existing" application

  Scenario: Download already present config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": "Konfig-server/1.0 http/1.0" }}' value in vault
    And I have already downloaded the configuration for "existing" application
    When I download the configuration for "existing" application
    Then I should not have written a config for "existing" application

  Scenario: Download missing config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    When I download the configuration for "missing" application
    Then I should not have a configuration for "missing" application
    And I should see a log entry of "not found" in the file "reports/missing.client.log"

  Scenario: Download invalid host in config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    When I download the configuration for "invalid_host" application
    Then I should not have a configuration for "invalid_host" application
    And I should see a log entry of "context deadline exceeded" in the file "reports/invalid_host.client.log"

  Scenario: Download invalid content type in config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    When I download the configuration for "invalid_content_type" application
    Then I should not have a configuration for "invalid_content_type" application
    And I should see a log entry of "could not transform" in the file "reports/invalid_content_type.client.log"
