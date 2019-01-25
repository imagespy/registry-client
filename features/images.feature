Feature: Images

   Scenario: Get Image by tag
    Given a running Docker registry at "127.0.0.1:6363"
    And a Docker image "127.0.0.1:6363/e2e:test" built from "Dockerfile"
    And a Docker image "127.0.0.1:6363/e2e:test" pushed
    When getting image "127.0.0.1:6363/e2e:test"
    Then the image has the digest "sha256:3d2e482b82608d153a374df3357c0291589a61cc194ec4a9ca2381073a17f58e"
