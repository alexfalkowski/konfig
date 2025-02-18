@manual @http
Feature: Secrets
  Secrets allows get the configured secrets.

  Background:
    Given I have a "git" valid setup
    And I start the system
    And I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |

  Scenario: Secrets with HTTP
    Given I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    When I request secrets with HTTP:
      | vault | vault:/secret/data/transport/http/user_agent |
      | ssm   | ssm:/secret/data/transport/grpc/user_agent   |
      | file  | file:secrets/test                            |
    Then I should receive valid secrets from HTTP:
      | vault | Konfig-server/1.0 http/1.0 |
      | ssm   | Konfig-server/1.0 grpc/1.0 |
      | file  | This is a secret           |

  Scenario: Missing secrets with HTTP
    When I request secrets with HTTP:
      | vault | vault:/secret/data/transport/http/user_agent |
      | ssm   | ssm:/secret/data/transport/grpc/user_agent   |
      | file  | file:../secrets/test                         |
    Then I should receive valid secrets from HTTP:
      | vault | vault:/secret/data/transport/http/user_agent |
      | ssm   | ssm:/secret/data/transport/grpc/user_agent   |
      | file  | file:../secrets/test                         |

  @reset
  Scenario: Secrets with HTTP and broken vault
    Given I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request secrets with HTTP:
      | vault | vault:/secret/data/transport/http/user_agent |
    Then I should receive an internal error from HTTP
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy
