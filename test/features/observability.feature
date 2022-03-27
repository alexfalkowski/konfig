@startup
Feature: Observability

  Observability is a measure of how well internal states of a system can be inferred by knowledge of its external outputs.

  Scenario: Health with HTTP
    When the system requests the "health" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Liveness with HTTP
    When the system requests the "liveness" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Readiness with HTTP
    When the system requests the "readiness" with HTTP
    Then the system should respond with a healthy status with HTTP

  Scenario: Health with gRPC
    When the system requests the health status with gRPC
    Then the system should respond with a healthy status with gRPC

  Scenario: Metrics
    When the system requests the "metrics" with HTTP
    Then the system should respond with metrics
