Feature: Tags

   Scenario: List Tags
      Given a running Docker registry at "127.0.0.1:6363"
      And a Docker image "127.0.0.1:6363/e2e:test" built from "Dockerfile"
      When listing tags of "127.0.0.1:6363/e2e"
      Then the list of tags contains "test"
