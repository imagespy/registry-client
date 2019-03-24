package registry

import (
	"fmt"
)

var (
	theImageHasTheDigestResult string
)

func gettingImage(imageName string) error {
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

	theImageHasTheDigestResult = image.Digest
	return nil
}

func theImageHasTheDigest(digest string) error {
	if theImageHasTheDigestResult != digest {
		return fmt.Errorf("expected digest '%s' not equal actual digest '%s'", digest, theImageHasTheDigestResult)
	}

	return nil
}

func deletingTheImage(imageName string) error {
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
	img, err := repository.Images().GetByTag(tag)
	if err != nil {
		return err
	}

	return repository.Images().DeleteByDigest(img.Digest)
}

func theImageDoesNotExist(imageName string) error {
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
	_, err = repository.Images().GetByTag(tag)
	if err == nil {
		return fmt.Errorf("Expected image %s to not exist but it does", imageName)
	}

	if err != ErrResourceNotFound {
		return fmt.Errorf("Expected error ErrResourceNotFound to be returned but got: %s", err)
	}

	return nil
}
