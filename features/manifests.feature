Feature: Manifests

   Scenario: Get Manifest
    Given a running Docker registry at "127.0.0.1:6363"
    And a Docker image "127.0.0.1:6363/e2e:test" built from "Dockerfile"
    And a Docker image "127.0.0.1:6363/e2e:test" pushed
    When getting the manifest of "127.0.0.1:6363/e2e:test" for platform with os "linux" and arch "amd64"
    Then the manifest has the media type "application/vnd.docker.distribution.manifest.v2+json"
