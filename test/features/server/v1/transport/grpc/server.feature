@manual
Feature: Server

  Server allows users to manage their application configurations.

  Scenario Outline: Existing config with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
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

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |
      | git    | test | v1.7.0 | staging | eu      | server |
      | folder | test | v1.7.0 | staging | eu      | server |
      | s3     | test | v1.7.0 | staging | eu      | server |

  Scenario Outline: Existing config with gRPC multiple times
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have the following provider information:
      | provider | key                                   | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
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

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |
      | git    | test | v1.7.0 | staging | eu      | server |
      | folder | test | v1.7.0 | staging | eu      | server |
      | s3     | test | v1.7.0 | staging | eu      | server |

  Scenario Outline: Existing config with non existent provider data with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I do not have the following provider information:
      | provider | key                                   |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config with missing provider data from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |
      | git    | test | v1.7.0 | staging | eu      | server |
      | folder | test | v1.7.0 | staging | eu      | server |
      | s3     | test | v1.7.0 | staging | eu      | server |

  Scenario Outline: Existing config with missing provider data with gRPC
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have the following provider information:
      | provider | key                                   | value        |
      | vault    | /secret/data/transport/http/user_agent | {"data": {}} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": {}} |
    When I request a config with gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |
    Then I should receive a valid config with missing provider data from gRPC:
      | source  | <source>  |
      | app     | <app>     |
      | ver     | <ver>     |
      | env     | <env>     |
      | cluster | <cluster> |
      | cmd     | <cmd>     |

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |
      | git    | test | v1.7.0 | staging | eu      | server |
      | folder | test | v1.7.0 | staging | eu      | server |
      | s3     | test | v1.7.0 | staging | eu      | server |

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

    Examples:
      | source | app     | ver    | env     | cluster | cmd     |
      | git    | missing | v1.7.0 | staging | *       | server  |
      | git    | test    | v1.7.0 | staging | *       | missing |
      | folder | missing | v1.7.0 | staging | *       | server  |
      | folder | test    | v1.7.0 | staging | *       | missing |
      | s3     | missing | v1.7.0 | staging | *       | server  |
      | s3     | test    | v1.7.0 | staging | *       | missing |

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

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |

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

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    |      | v1.7.0 | staging | *       | server |
      | git    | test |        | staging | *       | server |
      | git    | test | v1.7.0 |         | *       | server |
      | git    | test | v1.7.0 | staging |         |        |
      | folder |      | v1.7.0 | staging | *       | server |
      | folder | test |        | staging | *       | server |
      | folder | test | v1.7.0 |         | *       | server |
      | folder | test | v1.7.0 | staging |         |        |
      | s3     |      | v1.7.0 | staging | *       | server |
      | s3     | test |        | staging | *       | server |
      | s3     | test | v1.7.0 |         | *       | server |
      | s3     | test | v1.7.0 | staging |         |        |

  Scenario Outline: Existing config with gRPC and broken vault
    Given I have a "<source>" valid setup
    And I have "<source>" as the config file
    And I start the system
    And I have the following provider information:
      | provider | key                                   | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
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

    Examples:
      | source | app  | ver    | env     | cluster | cmd    |
      | git    | test | v1.7.0 | staging | *       | server |
      | folder | test | v1.7.0 | staging | *       | server |
      | s3     | test | v1.7.0 | staging | *       | server |
