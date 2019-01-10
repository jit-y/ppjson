package ppjson_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/jit-y/ppjson"
)

type testData struct {
	data     []byte
	expected string
}

func TestString(t *testing.T) {
	tests := buildTestData()

	for i, test := range tests {
		p := ppjson.NewPrinter(os.Stdout, test.data)
		actual := p.String()

		if test.expected != actual {
			fmt.Println(test.expected, actual)
			t.Errorf("tests[%d] wrong. expected=%s, got=%s", i, test.expected, actual)
		}
	}
}

func buildTestData() []testData {
	return []testData{
		{
			data:     toJson("test"),
			expected: "\"test\"",
		},
		{
			data:     toJson(nil),
			expected: "null",
		},
		{
			data:     toJson(1234567890),
			expected: "1234567890",
		},
	}
}

func toJson(v interface{}) []byte {
	data, _ := json.Marshal(v)

	return data
}
