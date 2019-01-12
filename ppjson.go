package ppjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type printer struct {
	reader  io.Reader
	writer  io.Writer
	value   interface{}
	indent  int
	newLine string
}

func Format(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	p := NewPrinter(buf, os.Stdout)

	val, err := p.Pretty()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func NewPrinter(reader io.Reader, writer io.Writer) *printer {
	return &printer{
		reader:  reader,
		writer:  writer,
		indent:  2,
		newLine: "\n",
	}
}

func (p *printer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
	val, err := p.Pretty()
	if err != nil {
		return 0, err
	}

	return buf.Write([]byte(val))
}

func (p *printer) Pretty() (string, error) {
	dec := json.NewDecoder(p.reader)
	t, err := dec.Token()
	if err != nil {
		return "", err
	}

	val, err := p.format(t)
	if err != nil {
		return "", err
	}

	return val, nil
}

func (p *printer) format(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return p.formatString(val)
	case nil:
		return "null", nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	default:
		k := reflect.ValueOf(val).Kind()
		return "", fmt.Errorf("parse failed: type %v is not supported", k)
	}
}

func (p *printer) formatString(val string) (string, error) {
	writer := bytes.Buffer{}
	encoder := json.NewEncoder(&writer)

	err := encoder.Encode(val)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(writer.String(), "\n"), nil
}
