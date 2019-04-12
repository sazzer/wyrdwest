package uritemplate_test

import (
	"testing"

	"github.com/sazzer/wyrdwest/service/internal/api/uritemplate"
	"github.com/stretchr/testify/assert"
)

func TestBuildURIPath(t *testing.T) {
	result := uritemplate.BuildURI("/a/b{/id}", map[string]interface{}{"id": "123"})
	assert.Equal(t, "/a/b/123", result)
}

func TestBuildURIQuery(t *testing.T) {
	result := uritemplate.BuildURI("/a/b{?id}", map[string]interface{}{"id": "123"})
	assert.Equal(t, "/a/b?id=123", result)
}

func TestBuildURIInvalid(t *testing.T) {
	assert.Panics(t, func() {
		uritemplate.BuildURI("/a/b{?id", map[string]interface{}{"id": "123"})
	})
}
