package ppjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type printer struct {
	value   interface{}
	indent  int
	newLine string
	out     io.Writer
}

func Format(data []byte) ([]byte, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	p := NewPrinter(os.Stdout, v)

	val := p.String()

	return []byte(val), nil
}

func NewPrinter(out io.Writer, obj interface{}) *printer {
	return &printer{
		value:   obj,
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
	case int:
		return strconv.Itoa(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return "parse failed"
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
