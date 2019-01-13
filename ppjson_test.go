package ppjson_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/jit-y/ppjson"
)

type testData struct {
	desc     string
	data     []byte
	expected string
}

func TestString(t *testing.T) {
	tests := buildTestData()

	for _, test := range tests {
		buf := bytes.NewBuffer(test.data)
		p := ppjson.NewPrinter(buf, os.Stdout)
		actual, err := p.Pretty()
		if err != nil {
			t.Errorf("test %s: %v", test.desc, err)
		}

		if test.expected != actual {
			t.Errorf("test %s: wrong. expected=%s, got=%s", test.desc, test.expected, actual)
		}
	}
}

func buildTestData() []testData {
	return []testData{
		{
			desc:     "string",
			data:     toJson("test"),
			expected: "\"test\"",
		},
		{
			desc:     "nil",
			data:     toJson(nil),
			expected: "null",
		},
		{
			desc:     "int",
			data:     toJson(1234567890),
			expected: "1234567890",
		},
		{
			desc: "array",
			data: toJson([]interface{}{1, "111", nil}),
			expected: `[
  1,
  "111",
  null
]`,
		},
		{
			desc: "map",
			data: []byte("{\"foo\":null,\"bar\":\"111\",\"baz\":1,\"array\":[1,\"111\",null]}"),
			expected: `{
  "foo": null,
  "bar": "111",
  "baz": 1,
  "array": [
    1,
    "111",
    null
  ]
}`,
		},
	}
}

func toJson(v interface{}) []byte {
	data, _ := json.Marshal(v)

	return data
}
