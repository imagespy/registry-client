package registry

import (
	"net/http"
	"testing"
	"time"

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

func TestTokenAuthenticator_HandleResponse_ErrorIfAuthFails(t *testing.T) {
	ta := &tokenAuthenticator{
		expiresAt: time.Now().Add(30 * time.Second),
		token:     "abc123",
	}

	resp := &http.Response{
		StatusCode: http.StatusUnauthorized,
	}

	_, resend, err := ta.HandleResponse(resp)
	assert.False(t, resend)
	assert.Error(t, err)
	assert.Equal(t, ErrAuthTokenInvalid, err)
}
