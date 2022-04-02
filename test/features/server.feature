@manual @clear
Feature: Server

  Server allows users to manage their application configurations.

  Scenario: Existing cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from HTTP

  Scenario: Existing cfg twice with HTTP
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with HTTP
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from HTTP

  Scenario Outline: Missing cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with HTTP
    Then I should receive a missing cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from HTTP

    Examples:
      | app     | ver    | env     | cmd     |
      | missing | v1.5.0 | staging | server  |
      | test    | v1.5.0 | staging | missing |

  Scenario: Misconfigured cfg with HTTP
    Given I have a misconfigured vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "test" app with "v1.5.0" ver from "test" env and "server" cmd with HTTP
    Then I should receive a missing cfg from "test" app with "v1.5.0" ver and "test" env and "server" cmd from HTTP

  Scenario Outline: Invalid cfg with HTTP
    Given I have a valid vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with HTTP
    Then I should receive an invalid cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from HTTP

    Examples:
      | app  | ver    | env     | cmd    |
      |      | v1.5.0 | staging | server |
      | test |        | staging | server |
      | test | v1.5.0 |         | server |
      | test | v1.5.0 | staging |        |

  Scenario: Existing cfg over HTTP with broken vault
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'vault' to 'close_all'
    And I wait for 2 seconds for the changes to apply
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive an internal error from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from HTTP
    And I should reset the proxy for service 'vault'
    And I wait for 2 seconds for the changes to apply

  Scenario: Existing cfg over HTTP with broken cache
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'redis' to 'close_all'
    And I wait for 2 seconds for the changes to apply
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with HTTP
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from HTTP
    And I should reset the proxy for service 'redis'
    And I wait for 2 seconds for the changes to apply

  Scenario: Existing cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from gRPC

  Scenario: Existing cfg twice with gRPC
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with gRPC
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from gRPC

  Scenario Outline: Missing cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with gRPC
    Then I should receive a missing cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from gRPC

    Examples:
      | app     | ver    | env     | cmd     |
      | missing | v1.5.0 | staging | server  |
      | test    | v1.5.0 | staging | missing |

  Scenario: Misconfigured cfg with gRPC
    Given I have a misconfigured vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "test" app with "v1.5.0" ver from "test" env and "server" cmd with gRPC
    Then I should receive a missing cfg from "test" app with "v1.5.0" ver and "test" env and "server" cmd from gRPC

  Scenario Outline: Invalid cfg with gRPC
    Given I have a valid vcs token
    And I start nonnative
    And I have a fresh cache
    When I request "<app>" app with "<ver>" ver from "<env>" env and "<cmd>" cmd with gRPC
    Then I should receive an invalid cfg from "<app>" app with "<ver>" ver and "<env>" env and "<cmd>" cmd from gRPC

    Examples:
      | app  | ver    | env     | cmd    |
      |      | v1.5.0 | staging | server |
      | test |        | staging | server |
      | test | v1.5.0 |         | server |
      | test | v1.5.0 | staging |        |

  Scenario: Existing cfg over gRPC with broken cache
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'redis' to 'close_all'
    And I wait for 2 seconds for the changes to apply
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive a valid cfg from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from gRPC
    And I should reset the proxy for service 'redis'
    And I wait for 2 seconds for the changes to apply

  Scenario: Existing cfg over gRPC with broken vault
    Given I have a valid vcs token
    And I start nonnative
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'vault' to 'close_all'
    And I wait for 2 seconds for the changes to apply
    When I request "test" app with "v1.5.0" ver from "staging" env and "server" cmd with gRPC
    Then I should receive an internal error from "test" app with "v1.5.0" ver and "staging" env and "server" cmd from gRPC
    And I should reset the proxy for service 'vault'
    And I wait for 2 seconds for the changes to apply
