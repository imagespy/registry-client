package registry

import (
	"fmt"

	"github.com/docker/distribution/manifest/schema1"
)

var (
	getConfigV1Result schema1.SignedManifest
)

func gettingTheConfigOf(imageName string) error {
	domain, path, tag, _, err := ParseImageName(imageName)
	if err != nil {
		return err
	}

	reg := &Registry{
		Client:   DefaultClient(),
		Domain:   domain,
		Protocol: "http",
	}

	repository := reg.Repository(path)
	image, err := repository.Images.GetByTag(tag)
	if err != nil {
		return err
	}

	getConfigV1Result, err = repository.Configs.GetV1(image.Tag)
	if err != nil {
		return err
	}

	return nil
}

func theConfigForRepositoryAndTagIsReturned(name, tag string) error {
	if getConfigV1Result.Name != name {
		return fmt.Errorf("expected name '%s' != actual name '%s' of config v1", name, getConfigV1Result.Name)
	}

	if getConfigV1Result.Tag != tag {
		return fmt.Errorf("expected tag '%s' != actual tag '%s' of config v1", tag, getConfigV1Result.Tag)
	}

	return nil
}
