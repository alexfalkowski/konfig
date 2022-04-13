@manual
Feature: Client

  Client allows the system to download a new configuration.

  Scenario: Download existing config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    When I download the configuration for "existing" application
    Then I should have a configuration for "existing" application

  Scenario: Download missing config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    And I have a fresh cache
    When I download the configuration for "missing" application
    Then I should not have a configuration for "missing" application

   Scenario: Download invalid config
    Given I have a "folder" valid setup
    And I have "folder" as the config file
    And I start the system
    And I have a fresh cache
    When I download the configuration for "invalid" application
    Then I should not have a configuration for "invalid" application
