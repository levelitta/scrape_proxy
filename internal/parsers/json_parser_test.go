package parsers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var in = `
[
  {
    "key": "initialState",
    "value": {
      "results": {
        "aggregatedOffers": 43
      }
    }
  },
  {
    "key": "invalidKey",
    "value": {
      "results": "invalidValue"
    }
  }
]
`

var invalidIn = `
[
  {
    "key": "invalid_key",
    "value": {
      "results": {
        "aggregatedOffers": 43
      }
    }
  },
  {
    "key": "invalidKey",
    "value": {
      "results": "invalidValue"
    }
  }
]
`

const jsonExpr = "@this.#(key==\"initialState\").value.results"

func TestParseOffers(t *testing.T) {
	t.Run("parse offers", func(t *testing.T) {
		jsonParser := NewJsonParser(jsonExpr)
		out, err := jsonParser.Parse(in)
		assert.NoError(t, err)

		want := `{"aggregatedOffers":43}`

		actual, err := jsonStringCompact(out)
		assert.NoError(t, err)

		assert.Equal(t, want, actual)
	})

	t.Run("parse offers failed", func(t *testing.T) {
		jsonParser := NewJsonParser(jsonExpr)
		out, err := jsonParser.Parse(invalidIn)
		assert.Error(t, err)

		want := ``

		actual, err := jsonStringCompact(out)
		assert.EqualError(t, err, "unexpected end of JSON input")

		assert.Equal(t, want, actual)
	})
}

func jsonStringCompact(in string) (string, error) {
	res := &bytes.Buffer{}
	err := json.Compact(res, []byte(in))

	return res.String(), err
}
