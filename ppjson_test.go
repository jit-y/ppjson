package ppjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jit-y/ppjson"
)

type data struct {
	data     []byte
	expected []byte
}

func TestFormat(t *testing.T) {
	tests := []data{
		{
			data:     prepareJSON("test"),
			expected: []byte("\"test\""),
		},
		{
			data:     prepareJSON(nil),
			expected: []byte("null"),
		},
		{
			data:     prepareJSON(1234567890),
			expected: []byte("1234567890"),
		},
	}

	for i, test := range tests {
		actual, err := ppjson.Format(test.data)
		if err != nil {
			t.Errorf("tests[%d] %v", i, err)
		}

		if !bytes.Equal(test.expected, actual) {
			fmt.Println(test.expected, actual)
			t.Errorf("tests[%d] wrong. expected=%s, got=%s", i, test.expected, actual)
		}
	}
}

func prepareJSON(v interface{}) []byte {
	val, _ := json.Marshal(v)

	return val
}
