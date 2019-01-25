Feature: Configs

   Scenario: Get Config V1
    Given a running Docker registry at "127.0.0.1:6363"
    And a Docker image "127.0.0.1:6363/e2e:test" built from "Dockerfile"
    And a Docker image "127.0.0.1:6363/e2e:test" pushed
    When getting the config of "127.0.0.1:6363/e2e:test"
    Then the config for repository "e2e" and tag "test" is returned
