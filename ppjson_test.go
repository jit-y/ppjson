package ppjson_test

import (
	"bytes"
	"encoding/json"
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
		p := ppjson.NewPrinter(buf)
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
			data: []byte("[1,\"111\",null, {\"foo\": null, \"bar\": \"111\", \"baz\": 1, \"array\": [1, \"111\", null]}]"),
			expected: `[
  1,
  "111",
  null,
  {
    "foo": null,
    "bar": "111",
    "baz": 1,
    "array": [
      1,
      "111",
      null
    ]
  }
]`,
		},
		{
			desc: "map",
			data: []byte("{\"foo\":null,\"bar\":\"111\",\"baz\":1,\"array\":[1,\"111\",null],\"map\":{\"a\":\"1\",\"b\":null,\"c\":12345}}"),
			expected: `{
  "foo": null,
  "bar": "111",
  "baz": 1,
  "array": [
    1,
    "111",
    null
  ],
  "map": {
    "a": "1",
    "b": null,
    "c": 12345
  }
}`,
		},
	}
}

func toJson(v interface{}) []byte {
	data, _ := json.Marshal(v)

	return data
}
