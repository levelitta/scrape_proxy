package http_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHttpRequest(t *testing.T) {
	t.Run("add user agent header", func(t *testing.T) {
		t.Parallel()

		req, err := NewHttpRequest("url", "GET", map[string]string{"User-Agent": "mozilla"})
		assert.NoError(t, err)

		assert.Equal(t, "mozilla", req.Header.Get("User-Agent"))
	})
	t.Run("default user agent", func(t *testing.T) {
		t.Parallel()

		req, err := NewHttpRequest("url", "GET", map[string]string{})
		assert.NoError(t, err)

		assert.Equal(t, defaultUserAgent, req.Header.Get("User-Agent"))
	})
}
