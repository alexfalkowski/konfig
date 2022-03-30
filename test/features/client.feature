@manual
Feature: Client

  Client allows the system to download a new configuration.

  Scenario: Download existing config
    Given I have a valid vcs token
    And I start nonnative
    When I download the configuration for "existing" application
    Then I should have a configuration for "existing" application

  Scenario: Download missing config
    Given I have a valid vcs token
    And I start nonnative
    When I download the configuration for "missing" application
    Then I should not have a configuration for "missing" application

   Scenario: Download invalid config
    Given I have a valid vcs token
    And I start nonnative
    When I download the configuration for "invalid" application
    Then I should not have a configuration for "invalid" application
