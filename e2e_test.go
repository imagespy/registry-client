package registry

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/DATA-DOG/godog"
)

func aRunningDockerRegistryAt(address string) error {
	cmd := exec.Command("docker-compose", "up", "-d")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker-compose failed with '%s': %s", err, string(out))
	}

	c := DefaultClient()
	for try := 1; try <= 5; try++ {
		resp, err := c.Get("http://" + address + "/v2")
		if err == nil {
			if resp.StatusCode == http.StatusOK {
				return nil
			}
		}

		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("registry container did not start within 10 seconds")
}

func aDockerImageBuiltFrom(image, pathDockerfile string) error {
	cmd := exec.Command("docker", "build", "-t", image, "-f", pathDockerfile, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("'docker build' failed with '%s': %s", err, string(out))
	}

	return nil
}

func aDockerImagePushed(image string) error {
	cmd := exec.Command("docker", "push", image)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("'docker push' failed with '%s': %s", err, string(out))
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a running Docker registry at "([^"]*)"$`, aRunningDockerRegistryAt)
	s.Step(`^a Docker image "([^"]*)" built from "([^"]*)"$`, aDockerImageBuiltFrom)
	s.Step(`^a Docker image "([^"]*)" pushed$`, aDockerImagePushed)
	s.Step(`^listing tags of "([^"]*)"$`, listingTagsOf)
	s.Step(`^the list of tags contains "([^"]*)"$`, theListOfTagsContains)
	s.Step(`^getting image "([^"]*)"$`, gettingImage)
	s.Step(`^the image has the digest "([^"]*)"$`, theImageHasTheDigest)
	s.Step(`^getting the manifest of "([^"]*)" for platform with os "([^"]*)" and arch "([^"]*)"$`, gettingTheManifestOfForPlatformWithOsAndArch)
	s.Step(`^the manifest has the media type "([^"]*)"$`, theManifestHasTheMediaType)
	s.Step(`^deleting the image "([^"]*)"$`, deletingTheImage)
	s.Step(`^the image "([^"]*)" does not exist$`, theImageDoesNotExist)
}
