package registry

import (
	"fmt"
)

var (
	listTagsResult []string
)

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
