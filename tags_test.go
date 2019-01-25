package registry

import (
	"fmt"
	"os/exec"

	"github.com/DATA-DOG/godog"
)

var (
	listTagsResult []string
)

func aRunningDockerRegistryAt(address string) error {
	cmd := exec.Command("docker-compose", "up", "-d")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func aDockerImageBuiltFrom(image, pathDockerfile string) error {
	cmd := exec.Command("docker", "build", "-t", image, "-f", pathDockerfile, ".")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func listingTagsOf(repo string) error {
	reg := &Registry{
		Client:   DefaultClient(),
		Domain:   "127.0.0.1:6363",
		Protocol: "http",
	}

	repository, err := reg.RepositoryFromString(repo)
	if err != nil {
		return err
	}

	listTagsResult, err = repository.Tags.GetAll()
	if err != nil {
		return err
	}

	return nil
}

func theListOfTagsContains(tag string) error {
	for _, t := range listTagsResult {
		if t == tag {
			return nil
		}
	}

	return fmt.Errorf("Tag '%s' not found", tag)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a running Docker registry at "([^"]*)"$`, aRunningDockerRegistryAt)
	s.Step(`^a Docker image "([^"]*)" built from "([^"]*)"$`, aDockerImageBuiltFrom)
	s.Step(`^listing tags of "([^"]*)"$`, listingTagsOf)
	s.Step(`^the list of tags contains "([^"]*)"$`, theListOfTagsContains)
}
