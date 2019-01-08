package ppjson

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Printer struct {
	Indent  int
	NewLine string
	Stdin   io.Reader
	Stdout  io.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Indent:  2,
		NewLine: "\n",
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
	}
}

func (p *Printer) Unmarshal(data []byte) (string, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return "", err
	}

	return p.PrettyPrint(v)
}

func (p *Printer) PrettyPrint(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return val, nil
	case nil:
		return "null", nil
	}

	return "", errors.New("should not reach here")
}
