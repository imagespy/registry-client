package registry

import (
	"fmt"

	"github.com/docker/distribution/manifest/schema2"
)

var (
	gettingTheManifestOfForPlatformWithOsAndArchResult schema2.Manifest
)

func gettingTheManifestOfForPlatformWithOsAndArch(imageName, os, arch string) error {
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
	image, err := repository.Images().GetByTag(tag)
	if err != nil {
		return err
	}

	for _, p := range image.Platforms {
		if p.Architecture == arch && p.OS == os {
			m, err := repository.Manifests().Get(p.Digest)
			if err != nil {
				return err
			}

			gettingTheManifestOfForPlatformWithOsAndArchResult = m
			return nil
		}
	}

	return fmt.Errorf("platform with os '%s' and arch '%s' not found", os, arch)
}

func theManifestHasTheMediaType(mediaType string) error {
	if gettingTheManifestOfForPlatformWithOsAndArchResult.MediaType != mediaType {
		return fmt.Errorf("expected media type '%s' != actual media type '%s'", mediaType, gettingTheManifestOfForPlatformWithOsAndArchResult.MediaType)
	}

	return nil
}
