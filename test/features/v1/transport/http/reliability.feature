@manual @http
Feature: Reliability
  This feature groups all the reliability features to assure we handle problems well.

  Scenario Outline: Existing config with HTTP and broken vault
    Given I have a "<source>" valid setup
    And I start the system
    And I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request a config with HTTP:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive an internal error from HTTP
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.11.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.11.0 | staging | *         | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | toml |
      | folder | test | v1.11.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.11.0 | staging | *         | *       | server | toml |

  Scenario Outline: Existing config with HTTP and broken aws
    Given I have a "<source>" valid setup
    And I start the system
    And I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    And I set the proxy for service 'aws' to 'close_all'
    And I should see "aws" as unhealthy
    When I request a config with HTTP:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive an internal error from HTTP
    And I should reset the proxy for service 'aws'
    And I should see "aws" as healthy

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | yaml |
      | folder | test | v1.11.0 | staging | *         | *       | server | yaml |
      | s3     | test | v1.11.0 | staging | *         | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | toml |
      | folder | test | v1.11.0 | staging | *         | *       | server | toml |
      | s3     | test | v1.11.0 | staging | *         | *       | server | toml |

  Scenario Outline: Existing config with HTTP and broken GitHub
    Given I have a "<source>" valid setup
    And the GitHub API is down
    And I start the system
    And I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
    When I request a config with HTTP:
      | source    | <source>    |
      | app       | <app>       |
      | ver       | <ver>       |
      | env       | <env>       |
      | continent | <continent> |
      | country   | <country>   |
      | cmd       | <cmd>       |
      | kind      | <kind>      |
    Then I should receive an internal error from HTTP
    And the GitHub API is back up

    Examples: With YAML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | yaml |

    Examples: With TOML kind
      | source | app  | ver     | env     | continent | country | cmd    | kind |
      | git    | test | v1.11.0 | staging | *         | *       | server | toml |

  Scenario Outline: Invalid service configurations
    When I have a "<source>" invalid setup
    Then starting the system should raise an error
    And I should see a log entry of "no configurator" in the file "reports/server.log"

    Examples:
      | source  |
      | missing |
      | wrong   |
