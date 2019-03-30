Feature: Registry

  Scenario: List repositories
    Given a running Docker registry at "127.0.0.1:6363"
    And a Docker image "127.0.0.1:6363/e2e:test" built from "Dockerfile"
    And a Docker image "127.0.0.1:6363/e2e:test" pushed
    When listing repositories in "127.0.0.1:6363"
    Then the list of repositories contains "e2e"
