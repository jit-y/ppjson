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

type Printer struct {
	buffer  *bytes.Buffer
	indent  int
	newLine string
	stdout  io.Writer
}

func Marshal(v interface{}) ([]byte, error) {
	p := NewPrinter()

	data, err := json.Marshal(&v)
	if err != nil {
		return nil, err
	}

	val, err := p.format(data)
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func NewPrinter() *Printer {
	return &Printer{
		indent:  2,
		newLine: "\n",
		stdout:  os.Stdout,
	}
}

func (p *Printer) format(data []byte) (string, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return "", err
	}

	switch val := v.(type) {
	case string:
		return p.formatString(val)
	case nil:
		return "null", nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	default:
		// should not reach here
		k := reflect.ValueOf(val).Kind()
		return "", fmt.Errorf("%v type is not supported", k)
	}
}

func (p *Printer) formatString(val string) (string, error) {
	writer := bytes.Buffer{}
	encoder := json.NewEncoder(&writer)

	err := encoder.Encode(val)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(writer.String(), "\n"), nil
}

func (p *Printer) formatInt(val int) (string, error) {
	v := strconv.Itoa(val)

	return v, nil
}
