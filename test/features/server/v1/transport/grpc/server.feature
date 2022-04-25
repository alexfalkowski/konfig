@manual
Feature: Server

  Server allows users to manage their application configurations.

  Scenario Outline: Existing config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": "Konfig-server/1.0 http/1.0" }}' value in vault
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
      | git    | test | v1.6.0 | staging | eu      | server |
      | folder | test | v1.6.0 | staging | eu      | server |

  Scenario Outline: Existing config with gRPC multiple times
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": "Konfig-server/1.0 http/1.0" }}' value in vault
    When I request a config with gRPC 2 times:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
      | git    | test | v1.6.0 | staging | eu      | server |
      | folder | test | v1.6.0 | staging | eu      | server |

  Scenario Outline: Existing config with missing vault value with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have no key "secret/data/transport/http/user_agent" in vault
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config with missing vault value from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
      | git    | test | v1.6.0 | staging | eu      | server |
      | folder | test | v1.6.0 | staging | eu      | server |

  Scenario Outline: Existing config with missing vault data with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": {}}' value in vault
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config with missing vault value from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
      | git    | test | v1.6.0 | staging | eu      | server |
      | folder | test | v1.6.0 | staging | eu      | server |

  Scenario Outline: Existing config with invalid vault data with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": 1 }}' value in vault
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config with missing vault value from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
      | git    | test | v1.6.0 | staging | eu      | server |
      | folder | test | v1.6.0 | staging | eu      | server |

  Scenario Outline: Missing config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a missing config from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app     | ver    | env     | cluster | cmd     |
      | git    | missing | v1.6.0 | staging | *       | server  |
      | git    | test    | v1.6.0 | staging | *       | missing |
      | folder | missing | v1.6.0 | staging | *       | server  |
      | folder | test    | v1.6.0 | staging | *       | missing |

  Scenario: Misconfigured config with gRPC
    Given I have a "<source>" invalid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a missing config from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |

  Scenario Outline: Invalid config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a invalid config from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    |      | v1.6.0 | staging | *       | server |
      | git    | test |        | staging | *       | server |
      | git    | test | v1.6.0 |         | *       | server |
      | git    | test | v1.6.0 | staging |         |        |
      | folder |      | v1.6.0 | staging | *       | server |
      | folder | test |        | staging | *       | server |
      | folder | test | v1.6.0 |         | *       | server |
      | folder | test | v1.6.0 | staging |         |        |

  Scenario Outline: Existing config with gRPC and broken vault
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "secret/data/transport/http/user_agent" with '{"data": { "value": "Konfig-server/1.0 http/1.0" }}' value in vault
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive an internal error from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy
    And the process 'server' should consume less than '40mb' of memory

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.6.0 | staging | *       | server |
      | folder | test | v1.6.0 | staging | *       | server |
