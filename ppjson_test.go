package ppjson_test

import (
	"testing"

	"github.com/jit-y/ppjson"
)

func TestPrettyPrintString(t *testing.T) {
	data := `test`
	expected := `test`
	p := ppjson.NewPrinter()

	actual, err := p.PrettyPrint(data)
	if err != nil {
		t.Error(err)
		return
	}

	if expected != actual {
		t.Errorf("not equal. expected=%s, got=%s", expected, actual)
	}
}

func TestPrettyPrintNil(t *testing.T) {
	expected := "null"
	p := ppjson.NewPrinter()

	actual, err := p.PrettyPrint(nil)
	if err != nil {
		t.Error(err)
		return
	}

	if expected != actual {
		t.Errorf("not equal. expected=%s, got=%s", expected, actual)
	}
}
