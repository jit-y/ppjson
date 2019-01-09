package ppjson_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/jit-y/ppjson"
)

type data struct {
	data     interface{}
	expected string
}

func TestString(t *testing.T) {
	tests := []data{
		{
			data:     "test",
			expected: "\"test\"",
		},
		{
			data:     nil,
			expected: "null",
		},
		{
			data:     1234567890,
			expected: "1234567890",
		},
	}

	for i, test := range tests {
		p := ppjson.NewPrinter(os.Stdout, test.data)
		actual := p.String()

		if test.expected != actual {
			fmt.Println(test.expected, actual)
			t.Errorf("tests[%d] wrong. expected=%s, got=%s", i, test.expected, actual)
		}
	}
}

func prepareJSON(v interface{}) []byte {
	val, _ := json.Marshal(v)

	return val
}
