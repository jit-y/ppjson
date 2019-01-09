package ppjson_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jit-y/ppjson"
)

type data struct {
	data     interface{}
	expected []byte
}

func TestMarshal(t *testing.T) {
	tests := []data{
		{
			data:     "test",
			expected: []byte("\"test\""),
		},
		{
			data:     nil,
			expected: []byte("null"),
		},
		{
			data:     1234567890,
			expected: []byte("1234567890"),
		},
	}

	for i, test := range tests {
		actual, err := ppjson.Marshal(test.data)
		if err != nil {
			t.Errorf("tests[%d] %v", i, err)
		}

		if !bytes.Equal(test.expected, actual) {
			fmt.Println(test.expected, actual)
			t.Errorf("tests[%d] wrong. expected=%s, got=%s", i, test.expected, actual)
		}
	}
}
