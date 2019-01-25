package registry

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestParseAuthHeader(t *testing.T) {
	h := `Bearer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:library/python:pull"`
	realm, scope, service, err := parseAuthHeader(h)
	require.NoError(t, err)
	assert.Equal(t, "https://auth.docker.io/token", realm)
	assert.Equal(t, "repository:library/python:pull", scope)
	assert.Equal(t, "registry.docker.io", service)
}
