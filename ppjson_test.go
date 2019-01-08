package ppjson_test

import (
	"testing"

	"github.com/jit-y/ppjson"
)

func TestUnmarshalString(t *testing.T) {
	data := []byte("\"test\"")
	expected := `test`
	p := ppjson.NewPrinter()

	actual, err := p.Unmarshal(data)
	if err != nil {
		t.Error(err)
		return
	}

	if expected != actual {
		t.Errorf("not equal. expected %s, but %s", expected, actual)
	}
}
