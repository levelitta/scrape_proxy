package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const expr = "window._cianConfig\\['frontend-serp'\\] = \\(window._cianConfig\\['frontend-serp'\\] \\|\\| \\[\\]\\).concat\\((.*\\])\\);"

func TestParse(t *testing.T) {
	t.Run("parse offers list", func(t *testing.T) {
		htmlParser, err := NewHtmlParser(expr)
		assert.NoError(t, err)

		actual, err := htmlParser.Parse("window._cianConfig['frontend-serp'] = (window._cianConfig['frontend-serp'] || []).concat([{\"key\":\"value\"}]);")
		assert.NoError(t, err)

		expected := "[{\"key\":\"value\"}]"
		assert.Equal(t, expected, actual)
	})

	t.Run("parse offers list with invalid string", func(t *testing.T) {
		htmlParser, err := NewHtmlParser(expr)
		assert.NoError(t, err)

		actual, err := htmlParser.Parse("fail string")
		assert.Error(t, err)

		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("compile regexp error", func(t *testing.T) {
		htmlParser, err := NewHtmlParser("(((((:ss")
		assert.Error(t, err)
		assert.Nil(t, htmlParser)
	})
}
