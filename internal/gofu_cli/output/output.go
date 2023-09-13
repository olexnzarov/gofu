package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	OUTPUT_TEXT = "text"
	OUTPUT_JSON = "json"
)

type OutputWriter interface {
	Text() string
	Object() interface{}
}

type Output struct {
	writers []indexedWriter
}

type indexedWriter struct {
	key    string
	writer OutputWriter
}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) Add(key string, writer OutputWriter) *Output {
	o.writers = append(o.writers, indexedWriter{key: key, writer: writer})
	return o
}

func (o *Output) json() (string, error) {
	object := map[string]interface{}{}
	for _, iw := range o.writers {
		object[iw.key] = iw.writer.Object()
	}
	bytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (o *Output) text() (string, error) {
	builder := strings.Builder{}
	for _, iw := range o.writers {
		builder.WriteString(iw.writer.Text())
	}
	return builder.String(), nil
}

func (o *Output) build(outputFormat string) (string, error) {
	switch outputFormat {
	case OUTPUT_JSON:
		return o.json()
	case OUTPUT_TEXT:
		return o.text()
	}
	return "", fmt.Errorf("unknown output format - %s", outputFormat)
}

func (o *Output) Print(outputFormat string) error {
	output, err := o.build(outputFormat)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
