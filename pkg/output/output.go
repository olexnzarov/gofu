package output

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type OutputWriter interface {
	Text() string
	Object() interface{}
}

type Output struct {
	writers      []indexedOutputWriter
	writersMutex *sync.RWMutex
}

type indexedOutputWriter struct {
	index  string
	writer OutputWriter
}

func NewOutput() *Output {
	return &Output{
		writersMutex: &sync.RWMutex{},
	}
}

func (o *Output) Add(key string, writer OutputWriter) {
	o.writersMutex.Lock()
	defer o.writersMutex.Unlock()
	o.writers = append(o.writers, indexedOutputWriter{index: key, writer: writer})
}

func (o *Output) Clear() {
	o.writersMutex.Lock()
	defer o.writersMutex.Unlock()
	o.writers = []indexedOutputWriter{}
}

func (o *Output) JSON(pretty bool) (string, error) {
	o.writersMutex.RLock()
	defer o.writersMutex.RUnlock()

	object := map[string]interface{}{}
	for _, iw := range o.writers {
		object[iw.index] = iw.writer.Object()
	}
	bytes, err := jsonMarshal(object, pretty)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (o *Output) Text() (string, error) {
	o.writersMutex.RLock()
	defer o.writersMutex.RUnlock()

	builder := strings.Builder{}
	for _, iw := range o.writers {
		builder.WriteString(iw.writer.Text())
		builder.WriteString("\n")
	}
	return builder.String(), nil
}

func (o *Output) Build(outputFormat string) (string, error) {
	switch outputFormat {
	case OUTPUT_JSON:
		return o.JSON(false)
	case OUTPUT_PRETTYJSON:
		return o.JSON(true)
	case OUTPUT_TEXT:
		return o.Text()
	}
	return "", fmt.Errorf("unknown output format - %s", outputFormat)
}

func (o *Output) Print(outputFormat string) error {
	str, err := o.Build(outputFormat)
	if err != nil {
		return err
	}
	_, err = fmt.Println(str)
	return err
}

func jsonMarshal(v any, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(v, "", "  ")
	}
	return json.Marshal(v)
}
