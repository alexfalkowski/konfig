@manual @grpc
Feature: Server

  Server allows users to manage their application configurations.

  Scenario Outline: Existing config with gRPC
    Given I have a "<source>" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive a valid config from gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | yaml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | git    | test | v1.10.0 | staging | eu        | de      | server | yaml |
      | folder | test | v1.10.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | folder | test | v1.10.0 | staging | eu        | de      | server | yaml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | eu        | de      | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | toml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | toml |
      | git    | test | v1.10.0 | staging | eu        | de      | server | toml |
      | folder | test | v1.10.0 | staging | *         | *       | server | toml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | toml |
      | folder | test | v1.10.0 | staging | eu        | de      | server | toml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | toml |
      | s3     | test | v1.10.0 | staging | eu        | de      | server | toml |

  Scenario Outline: Existing config with non existent provider data with gRPC
    Given I have a "<source>" valid setup
    And I start the system
    And I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive a valid config with missing provider data from gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.10.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | yaml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | toml |
      | folder | test | v1.10.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | toml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | toml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | toml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | toml |

  Scenario Outline: Existing config with missing provider data with gRPC
    Given I have a "<source>" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value        |
      | vault    | /secret/data/transport/http/user_agent | {"data": {}} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": {}} |
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive a valid config with missing provider data from gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.10.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | yaml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | toml |
      | folder | test | v1.10.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | toml |
      | git    | test | v1.10.0 | staging | eu        | *       | server | toml |
      | folder | test | v1.10.0 | staging | eu        | *       | server | toml |
      | s3     | test | v1.10.0 | staging | eu        | *       | server | toml |

  Scenario Outline: Missing config with gRPC
    Given I have a "<source>" valid setup
    And I start the system
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive a missing config from gRPC

    Examples: With YAML kind
      | source | app     | ver     | env     | continent | country | cmd     | kind |
      | git    | missing | v1.10.0 | staging | *         | *       | server  | yaml |
      | git    | test    | v1.10.0 | staging | *         | *       | missing | yaml |
      | folder | missing | v1.10.0 | staging | *         | *       | server  | yaml |
      | folder | test    | v1.10.0 | staging | *         | *       | missing | yaml |
      | s3     | missing | v1.10.0 | staging | *         | *       | server  | yaml |
      | s3     | test    | v1.10.0 | staging | *         | *       | missing | yaml |

    Examples: With TOML kind
      | source | app     | ver     | env     | continent | country | cmd     | kind |
      | git    | missing | v1.10.0 | staging | *         | *       | server  | toml |
      | git    | test    | v1.10.0 | staging | *         | *       | missing | toml |
      | folder | missing | v1.10.0 | staging | *         | *       | server  | toml |
      | folder | test    | v1.10.0 | staging | *         | *       | missing | toml |
      | s3     | missing | v1.10.0 | staging | *         | *       | server  | toml |
      | s3     | test    | v1.10.0 | staging | *         | *       | missing | toml |

  Scenario: Misconfigured config with gRPC
    Given I have a "<source>" invalid setup
    And I start the system
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive an internal error from gRPC

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.10.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | toml |
      | folder | test | v1.10.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | toml |

  Scenario Outline: Invalid config with gRPC
    Given I have a "<source>" valid setup
    And I start the system
    When I request a config with gRPC:
      | source | <source> |
      | app    | <app>    |
      | ver    | <ver>    |
      | env    | <env>    |
      | cmd    | <cmd>    |
      | kind   | <kind>   |
    Then I should receive a invalid config from gRPC

    Examples: With YAML kind
      | source | app  | ver     | env     | cmd    | kind |
      | git    |      | v1.10.0 | staging | server | yaml |
      | git    | test |         | staging | server | yaml |
      | git    | test | v1.10.0 |         | server | yaml |
      | git    | test | v1.10.0 | staging |        | yaml |
      | folder |      | v1.10.0 | staging | server | yaml |
      | folder | test |         | staging | server | yaml |
      | folder | test | v1.10.0 |         | server | yaml |
      | folder | test | v1.10.0 | staging |        | yaml |
      | s3     |      | v1.10.0 | staging | server | yaml |
      | s3     | test |         | staging | server | yaml |
      | s3     | test | v1.10.0 |         | server | yaml |
      | s3     | test | v1.10.0 | staging |        | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | cmd    | kind |
      | git    |      | v1.10.0 | staging | server | toml |
      | git    | test |         | staging | server | toml |
      | git    | test | v1.10.0 |         | server | toml |
      | git    | test | v1.10.0 | staging |        | toml |
      | folder |      | v1.10.0 | staging | server | toml |
      | folder | test |         | staging | server | toml |
      | folder | test | v1.10.0 |         | server | toml |
      | folder | test | v1.10.0 | staging |        | toml |
      | s3     |      | v1.10.0 | staging | server | toml |
      | s3     | test |         | staging | server | toml |
      | s3     | test | v1.10.0 |         | server | toml |
      | s3     | test | v1.10.0 | staging |        | toml |

  Scenario Outline: Existing config with gRPC and broken vault
    Given I have a "<source>" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request a config with gRPC:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive an internal error from gRPC
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.10.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.10.0 | staging | *         | *       | server | toml |
      | folder | test | v1.10.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.10.0 | staging | *         | *       | server | toml |
