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
	raw     []byte
	value   interface{}
	indent  int
	newLine string
	out     io.Writer
}

func Format(data []byte) ([]byte, error) {
	p := NewPrinter(os.Stdout, data)

	val := p.String()

	return []byte(val), nil
}

func NewPrinter(out io.Writer, raw []byte) *printer {
	return &printer{
		raw:     raw,
		indent:  2,
		newLine: "\n",
		out:     out,
	}
}

func (p *printer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
	val := []byte(p.String())

	return buf.Write(val)
}

func (p *printer) String() string {
	var v interface{}
	err := json.Unmarshal(p.raw, &v)
	if err != nil {
		return fmt.Sprintf("parse error: %v", err)
	}
	p.value = v

	return p.format()
}

func (p *printer) PrettyPrint() {
	fmt.Fprint(p.out, p.String())
}

func (p *printer) PP() {
	p.PrettyPrint()
}

func (p *printer) format() string {
	switch val := p.value.(type) {
	case string:
		return p.formatString(val)
	case nil:
		return "null"
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		k := reflect.ValueOf(val).Kind()
		return fmt.Sprintf("parse failed: type %v is not supported", k)
	}
}

func (p *printer) formatString(val string) string {
	writer := bytes.Buffer{}
	encoder := json.NewEncoder(&writer)

	err := encoder.Encode(val)
	if err != nil {
		return ""
	}

	return strings.TrimRight(writer.String(), "\n")
}
