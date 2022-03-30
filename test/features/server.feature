@manual
Feature: Server

  Server allows users to manage their application configurations.

  Scenario: Existing cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a valid cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from HTTP

  Scenario: Existing cfg twice with HTTP
    Given I have a valid vcs token
    And I start nonnative
    And I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with HTTP
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a valid cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from HTTP

  Scenario Outline: Missing cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with HTTP
    Then I should receive a missing cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from HTTP

    Examples:
      | app     | ver    | env     | cmd     |
      | missing | v1.1.0 | staging | server  |
      | test    | v1.1.0 | staging | missing |

  Scenario: Misconfigured cfg with HTTP
    Given I have a misconfigured vcs token
    And I start nonnative
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a missing cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from HTTP

  Scenario Outline: Invalid cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with HTTP
    Then I should receive an invalid cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from HTTP

    Examples:
      | app  | ver    | env     | cmd    |
      |      | v1.1.0 | staging | server |
      | test |        | staging | server |
      | test | v1.1.0 |         | server |
      | test | v1.1.0 | staging |        |

  Scenario: Existing cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a valid cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from gRPC

  Scenario: Existing cfg twice with gRPC
    Given I have a valid vcs token
    And I start nonnative
    And I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with gRPC
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a valid cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from gRPC

  Scenario Outline: Missing cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with gRPC
    Then I should receive a missing cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from gRPC

    Examples:
      | app     | ver    | env     | cmd     |
      | missing | v1.1.0 | staging | server  |
      | test    | v1.1.0 | staging | missing |

  Scenario: Misconfigured cfg with gRPC
    Given I have a misconfigured vcs token
    And I start nonnative
    When I request "test" app with "v1.1.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a missing cfg from "test" app with "v1.1.0" ver and "staging" env and "server" cmd from gRPC

  Scenario Outline: Invalid cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with gRPC
    Then I should receive an invalid cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from gRPC

    Examples:
      | app  | ver    | env     | cmd    |
      |      | v1.1.0 | staging | server |
      | test |        | staging | server |
      | test | v1.1.0 |         | server |
      | test | v1.1.0 | staging |        |
