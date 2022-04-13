@startup
Feature: Observability

  Observability is a measure of how well internal states of a system can be inferred by knowledge of its external outputs.

  Scenario: Health with gRPC
    When the system requests the health status with gRPC
    Then the system should respond with a healthy status with gRPC
