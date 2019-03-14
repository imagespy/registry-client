package registry

import (
	"fmt"
	"log"
)

func ExampleRegistry() {
	// Query an image in the Docker Hub.
	reg := &Registry{
		Authenticator: NewTokenAuthenticator(),
		Client:        DefaultClient(),
		Domain:        "docker.io",
	}
	repo, err := reg.RepositoryFromString("golang")
	if err != nil {
		log.Fatal(err)
	}

	img, err := repo.Images().GetByTag("1.12.0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Image Digest of Tag %s: %s\n", img.Tag, img.Digest)
	for _, p := range img.Platforms {
		fmt.Printf("Platform OS %s Architecture %s", p.OS, p.Architecture)
		m, err := repo.Manifests().Get(p.Digest)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Manifest of type %s", m.MediaType)
	}
}
