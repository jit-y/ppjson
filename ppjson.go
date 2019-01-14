package ppjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// Printer is a struct for print state.
type Printer struct {
	decoder *json.Decoder
	depth   int
	indent  int
	newLine string
}

// NewPrinter returns a pointer of initialized Printer object.
func NewPrinter(reader io.Reader) *Printer {
	dec := json.NewDecoder(reader)

	return &Printer{
		decoder: dec,
		indent:  2,
		depth:   0,
		newLine: "\n",
	}
}

// Pretty returns pretty formatted string and error
func (p *Printer) Pretty() (string, error) {
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

func (p *Printer) format(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return p.formatString(val)
	case nil:
		return p.formatNil()
	case float64:
		return p.formatFloat64(val)
	case json.Delim:
		return p.formatEnumerable(val)
	default:
		k := reflect.ValueOf(val).Kind()
		return "", fmt.Errorf("parse failed: type %v is not supported", k)
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

func (p *Printer) formatNil() (string, error) {
	return "null", nil
}

func (p *Printer) formatFloat64(val float64) (string, error) {
	return strconv.FormatFloat(val, 'f', -1, 64), nil
}

func (p *Printer) formatEnumerable(d json.Delim) (string, error) {
	switch d {
	case '[':
		return p.formatSlice(d)
	case '{':
		return p.formatMap(d)
	default:
		return "", errors.New("should not reach here")
	}
}

func (p *Printer) formatSlice(d json.Delim) (string, error) {
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

func (p *Printer) formatMap(d json.Delim) (string, error) {
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

func (p *Printer) currentIndent() string {
	return strings.Repeat(" ", p.indent*p.depth)
}

func (p *Printer) stringWithIndent(val string) string {
	return p.currentIndent() + val
}

type withIndentFunc func() error

func (p *Printer) withIndent(fn withIndentFunc) error {
	p.incrementDepth()
	defer p.decrementDepth()

	err := fn()

	return err
}

func (p *Printer) incrementDepth() {
	p.depth++
}

func (p *Printer) decrementDepth() {
	p.depth--
}
