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
	value   interface{}
	indent  int
	newLine string
	stdout  io.Writer
}

func Format(data []byte) ([]byte, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	p := NewPrinter(v)

	val, err := p.format()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func NewPrinter(obj interface{}) *Printer {
	return &Printer{
		value:   obj,
		indent:  2,
		newLine: "\n",
		stdout:  os.Stdout,
	}
}

func (p *Printer) format() (string, error) {
	switch val := p.value.(type) {
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
