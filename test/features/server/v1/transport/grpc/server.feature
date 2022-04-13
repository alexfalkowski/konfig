@manual @clear
Feature: Server

  Server allows users to manage their application configurations.

  Scenario Outline: Existing config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a valid config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario: Existing config multiple times with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    And I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a valid config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario Outline: Missing config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have a fresh cache
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a missing config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app     | ver    | env     | cmd     |
      | git    | missing | v1.5.0 | staging | server  |
      | git    | test    | v1.5.0 | staging | missing |
      | folder | missing | v1.5.0 | staging | server  |
      | folder | test    | v1.5.0 | staging | missing |

  Scenario: Misconfigured config with gRPC
    Given I have a "<source>" invalid setup
    And I have "<source>" as the config file
    And I start the system
    And I have a fresh cache
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a missing config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario Outline: Invalid config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have a fresh cache
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a invalid config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    |      | v1.5.0 | staging | server |
      | git    | test |        | staging | server |
      | git    | test | v1.5.0 |         | server |
      | git    | test | v1.5.0 | staging |        |
      | folder |      | v1.5.0 | staging | server |
      | folder | test |        | staging | server |
      | folder | test | v1.5.0 |         | server |
      | folder | test | v1.5.0 | staging |        |

  Scenario Outline: Existing config with gRPC and broken vault
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive an internal error from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario Outline: Existing config with gRPC and broken redis
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I have a fresh cache
    And I set the proxy for service 'redis' to 'close_all'
    And I should see "redis" as unhealthy
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a valid config from gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    And I should reset the proxy for service 'redis'
    And I should see "redis" as healthy

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |
