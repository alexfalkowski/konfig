@manual
Feature: Server

  Server allows users to manage their application configurations.

  Scenario Outline: Existing config with HTTP
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a valid config from HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario: Existing config multiple times with HTTP
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    And I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a valid config from HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario Outline: Missing config with HTTP
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a missing config from HTTP:
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

  Scenario: Misconfigured config with HTTP
    Given I have a "<source>" invalid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a missing config from HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |

    Examples:
      | source | app  | ver    | env     | cmd    |
      | git    | test | v1.5.0 | staging | server |
      | folder | test | v1.5.0 | staging | server |

  Scenario Outline: Invalid config with HTTP
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive a invalid config from HTTP:
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

  Scenario Outline: Existing config with HTTP and broken vault
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have key "transport/http/user_agent" with "Konfig-server/1.0 http/1.0" value in vault
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request a config with HTTP:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
    Then I should receive an internal error from HTTP:
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
