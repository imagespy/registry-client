package registry

import "fmt"

var (
	theImageHasTheDigestResult string
)

func gettingImageByTag(repo, tag string) error {
	reg := &Registry{
		Client:   DefaultClient(),
		Domain:   "127.0.0.1:6363",
		Protocol: "http",
	}

	repository, err := reg.RepositoryFromString(repo)
	if err != nil {
		return err
	}

	image, err := repository.Images.GetByTag(tag)
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
