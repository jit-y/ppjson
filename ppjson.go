package ppjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type printer struct {
	decoder *json.Decoder
	writer  io.Writer
	depth   int
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
	dec := json.NewDecoder(reader)

	return &printer{
		decoder: dec,
		writer:  writer,
		indent:  2,
		depth:   0,
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
	t, err := p.decoder.Token()
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
	case json.Delim:
		return p.formatEnumerable(val)
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

func (p *printer) formatEnumerable(d json.Delim) (string, error) {
	switch d {
	case '[':
		return p.formatSlice(d)
	case '{':
		return p.formatMap(d)
	default:
		return "", errors.New("should not reach here")
	}
}

func (p *printer) formatSlice(d json.Delim) (string, error) {
	var b strings.Builder
	b.WriteString(d.String() + p.newLine)

	err := p.withIndent(func() error {
		for i := 0; p.decoder.More(); i++ {
			tok, err := p.decoder.Token()
			if err != nil {
				return err
			}

			val, err := p.format(tok)
			if err != nil {
				return err
			}

			if i > 0 {
				b.WriteString("," + p.newLine)
			}

			b.WriteString(p.stringWithIndent(val))

		}

		return nil
	})

	if err != nil {
		return "", err
	}

	token, err := p.decoder.Token()
	if err != nil {
		return "", err
	}

	delim, ok := token.(json.Delim)
	if !ok {
		return "", errors.New("invalid format")
	}

	b.WriteString(p.newLine + p.stringWithIndent(delim.String()))

	return b.String(), nil
}

func (p *printer) formatMap(d json.Delim) (string, error) {
	var b strings.Builder
	b.WriteString(d.String() + p.newLine)

	err := p.withIndent(func() error {
		for i := 0; p.decoder.More(); i++ {
			keyTok, err := p.decoder.Token()

			if err != nil {
				return err
			}

			key, ok := keyTok.(string)
			if !ok {
				return errors.New("invalid format")
			}

			p.decoder.More()
			valTok, err := p.decoder.Token()
			if err != nil {
				return err
			}

			val, err := p.format(valTok)
			if err != nil {
				return err
			}

			if i > 0 {
				b.WriteString("," + p.newLine)
			}

			b.WriteString(p.stringWithIndent("\"" + key + "\"" + ": " + val))

		}

		return nil
	})

	if err != nil {
		return "", err
	}

	token, err := p.decoder.Token()
	if err != nil {
		return "", err
	}

	delim, ok := token.(json.Delim)
	if !ok {
		return "", errors.New("invalid format")
	}

	b.WriteString(p.newLine + p.stringWithIndent(delim.String()))

	return b.String(), nil
}

func (p *printer) currentIndent() string {
	return strings.Repeat(" ", p.indent*p.depth)
}

func (p *printer) stringWithIndent(val string) string {
	return p.currentIndent() + val
}

type withIndentFunc func() error

func (p *printer) withIndent(fn withIndentFunc) error {
	p.depth++

	err := fn()

	p.depth--

	return err
}
