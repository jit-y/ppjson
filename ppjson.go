package ppjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

	val := p.Pretty()

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
	val := []byte(p.Pretty())

	return buf.Write(val)
}

func (p *printer) Pretty() string {
	var v interface{}
	var b []byte
	buf := bytes.NewBuffer(b)
	io.Copy(buf, p.reader)

	data, err := ioutil.ReadAll(buf)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Sprintf("parse error: %v", err)
	}
	p.value = v

	return p.format()
}

func (p *printer) PrettyPrint() {
	fmt.Fprint(p.writer, p.Pretty())
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
