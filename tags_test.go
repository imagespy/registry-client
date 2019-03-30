package registry

import (
	"fmt"
)

var (
	listTagsResult []string
)

func listingTagsOf(repo string) error {
	domain, path, _, _, err := ParseImageName(repo)
	if err != nil {
		return err
	}

	reg := New(Options{
		Client:   DefaultClient(),
		Domain:   domain,
		Protocol: "http",
	})

	repository := reg.Repository(path)
	listTagsResult, err = repository.Tags().GetAll()
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
