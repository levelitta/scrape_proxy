package parsers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBase64Parser_Parse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
		err  error
	}{
		{
			name: "filled",
			in:   "dGVzdA==",
			want: "test",
		},
		{
			name: "russian symbols",
			in:   "0YLQtdGB0YI=",
			want: "тест",
		},
	}

	parser := NewBase64Parser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act, err := parser.Parse(tt.in)

			if tt.err == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.err)
			}

			require.Equal(t, tt.want, act)
		})

	}
}
