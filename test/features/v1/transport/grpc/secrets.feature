@manual @secrets
Feature: Secrets
  Secrets allows get the configured secrets.

  Background:
    Given I do not have the following provider information:
      | provider | key                                    |
      | vault    | /secret/data/transport/http/user_agent |
      | ssm      | /secret/data/transport/grpc/user_agent |

  Scenario: Secrets with gRPC
    Given I have a "git" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
      | ssm      | /secret/data/transport/grpc/user_agent | {"data": { "value": "Konfig-server/1.0 grpc/1.0" }} |
    When I request secrets with gRPC:
      | vault | vault:/secret/data/transport/http/user_agent |
      | ssm   | ssm:/secret/data/transport/grpc/user_agent   |
    Then I should receive valid secrets from gRPC:
      | vault | Konfig-server/1.0 http/1.0 |
      | ssm   | Konfig-server/1.0 grpc/1.0 |

  Scenario: Secrets with gRPC and broken vault
    Given I have a "git" valid setup
    And I start the system
    And I have the following provider information:
      | provider | key                                    | value                                               |
      | vault    | /secret/data/transport/http/user_agent | {"data": { "value": "Konfig-server/1.0 http/1.0" }} |
    And I set the proxy for service 'vault' to 'close_all'
    And I should see "vault" as unhealthy
    When I request secrets with gRPC:
      | vault | vault:/secret/data/transport/http/user_agent |
    Then I should receive an internal error from gRPC
    And I should reset the proxy for service 'vault'
    And I should see "vault" as healthy
