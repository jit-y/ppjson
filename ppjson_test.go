package ppjson_test

import (
	"testing"

	"github.com/jit-y/ppjson"
)

type data struct {
	data     []byte
	expected string
}

func TestUnmarshal(t *testing.T) {
	tests := []data{
		{
			data:     []byte("\"test\""),
			expected: `test`,
		},
		{
			data:     []byte("null"),
			expected: `null`,
		},
	}

	p := ppjson.NewPrinter()

	for i, test := range tests {
		actual, err := p.Unmarshal(test.data)
		if err != nil {
			t.Errorf("tests[%d] %v", i, err)
		}

		if test.expected != actual {
			t.Errorf("tests[%d] wrong. expected=%s, got=%s", i, test.expected, actual)
		}
	}
}
