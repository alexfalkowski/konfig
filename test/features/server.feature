Feature: Server

  Server allows users to manage their application configurations.

  Scenario: Existing config with HTTP
    When I request "test" app from "staging" env with HTTP
    Then I should receive a valid config from "test" app with "staging" env from HTTP

  Scenario: Existing config with gRPC
    When I request "test" app from "staging" env with gRPC
    Then I should receive a valid config from "test" app with "staging" env from gRPC
