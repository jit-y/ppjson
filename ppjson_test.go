package ppjson_test

import (
	"bytes"
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

func TestCompareFormatWithString(t *testing.T) {
	tests := buildTestData()

	for i, test := range tests {
		p := ppjson.NewPrinter(os.Stdout, test.data)
		str := []byte(p.String())
		format, err := ppjson.Format(str)
		if err != nil {
			t.Errorf("tests[%d] parse error: %v", i, err)
		}

		if !bytes.Equal(str, format) {
			t.Errorf("tests[%d] not equal. String=%s, Format=%s", i, str, format)
		}
	}
}

func buildTestData() []data {
	return []data{
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
		{
			data:     int8(123),
			expected: "123",
		},
		{
			data:     int16(123),
			expected: "123",
		},
		{
			data:     int32(123),
			expected: "123",
		},
		{
			data:     int64(123),
			expected: "123",
		},
		{
			data:     uint(123),
			expected: "123",
		},
		{
			data:     uint8(123),
			expected: "123",
		},
		{
			data:     uint16(123),
			expected: "123",
		},
		{
			data:     uint32(123),
			expected: "123",
		},
		{
			data:     uint64(123),
			expected: "123",
		},
		{
			data:     float32(123.456),
			expected: "123.456",
		},
		{
			data:     float64(123.456),
			expected: "123.456",
		},
	}
}
