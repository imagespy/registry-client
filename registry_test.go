package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestParseImageName_Tag(t *testing.T) {
	domain, path, tag, digest, err := ParseImageName("127.0.0.1:6363/e2e:test")
	require.NoError(t, err)
	assert.Equal(t, "127.0.0.1:6363", domain)
	assert.Equal(t, "e2e", path)
	assert.Equal(t, "test", tag)
	assert.Equal(t, "", digest)
}

func TestParseImageName_Digest(t *testing.T) {
	domain, path, tag, digest, err := ParseImageName("127.0.0.1:6363/e2e@sha256:3d2e482b82608d153a374df3357c0291589a61cc194ec4a9ca2381073a17f58e")
	require.NoError(t, err)
	assert.Equal(t, "127.0.0.1:6363", domain)
	assert.Equal(t, "e2e", path)
	assert.Equal(t, "", tag)
	assert.Equal(t, "sha256:3d2e482b82608d153a374df3357c0291589a61cc194ec4a9ca2381073a17f58e", digest)
}
